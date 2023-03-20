package utils

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
	"sort"
)

func PrecomputeRespNodes(nodesId []types.NodeId) [][4]types.NodeId {
	numPossibleChunks := constants.GetRangeAddress()
	result := make([][4]types.NodeId, numPossibleChunks)
	numNodesSearch := constants.GetBits()

	for chunkId := 0; chunkId < numPossibleChunks; chunkId++ {
		chunkIdInt32 := int32(chunkId)
		closestNodes := types.BinarySearchClosest(nodesId, int32(chunkId), numNodesSearch)
		distances := make([]int32, len(closestNodes))

		for j, nodeId := range closestNodes {
			distances[j] = nodeId.ToInt32() ^ chunkIdInt32
		}

		sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })

		for k := 0; k < 4; k++ {
			result[chunkId][k] = types.NodeId(distances[k] ^ chunkIdInt32) // this results in the nodeId again
		}
	}
	return result
}

func SortedKeys(nodeMap map[types.NodeId]*types.Node) []types.NodeId {
	keys := make([]types.NodeId, len(nodeMap))
	i := 0
	for k := range nodeMap {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func CreateGraphNetwork(net *types.Network) (*types.Graph, error) {
	//fmt.Println("Creating graph network...")
	sortedNodeIds := SortedKeys(net.NodesMap)
	numNodes := len(net.NodesMap)
	Edges := make(map[types.NodeId]map[types.NodeId]*types.Edge)
	respNodes := make([][4]types.NodeId, constants.GetRangeAddress())
	if constants.IsPrecomputeRespNodes() {
		respNodes = PrecomputeRespNodes(sortedNodeIds)
	}

	graph := &types.Graph{
		Network:   net,
		Nodes:     make([]*types.Node, 0, numNodes),
		Edges:     Edges,
		NodeIds:   sortedNodeIds,
		RespNodes: respNodes,
	}

	for _, nodeId := range sortedNodeIds {
		graph.Edges[nodeId] = make(map[types.NodeId]*types.Edge)

		node := net.NodesMap[nodeId]
		err1 := graph.AddNode(node)
		if err1 != nil {
			return nil, err1
		}

		nodeAdj := node.AdjIds
		for _, adjItems := range nodeAdj {
			for _, item := range adjItems {
				threshold := general.BitLength(nodeId.ToInt32() ^ item.ToInt32())
				epoch := constants.GetEpoch()
				attrs := types.EdgeAttrs{A2B: 0, Last: 0, EpochLastForgiven: epoch, Threshold: threshold}
				err := graph.AddEdge(node.Id, item, attrs)
				if err != nil {
					return nil, err
				}
				// graph.SetEdgeAttributes()
			}
		}
	}

	//fmt.Println("Graph network is created.")
	return graph, nil
}

func isThresholdFailed(firstNodeId types.NodeId, secondNodeId types.NodeId, chunkId types.ChunkId, graph *types.Graph, request types.Request) bool {
	if constants.GetThresholdEnabled() {
		edgeDataFirst := graph.GetEdgeData(firstNodeId, secondNodeId)
		p2pFirst := edgeDataFirst.A2B
		edgeDataSecond := graph.GetEdgeData(secondNodeId, firstNodeId)
		p2pSecond := edgeDataSecond.A2B

		threshold := constants.GetThreshold()
		if constants.IsAdjustableThreshold() {
			threshold = edgeDataFirst.Threshold
		}

		peerPriceChunk := PeerPriceChunk(secondNodeId, chunkId)
		price := p2pFirst - p2pSecond + peerPriceChunk
		//fmt.Printf("price: %d = p2pFirst: %d - p2pSecond: %d + PeerPriceChunk: %d \n", price, p2pFirst, p2pSecond, peerPriceChunk)

		if price > threshold {
			if constants.IsForgivenessEnabled() {
				newP2pFirst, forgiven := CheckForgiveness(edgeDataFirst, firstNodeId, secondNodeId, graph, request)
				//_, _ = CheckForgiveness(edgeDataSecond, secondNodeId, firstNodeId, graph, request)
				if forgiven {
					price = newP2pFirst - p2pSecond + peerPriceChunk
				}
			}
		}
		return price > threshold
	}
	return false
}

func getNext(request types.Request, firstNodeId types.NodeId, prevNodePaid bool, graph *types.Graph, rerouteStruct types.RerouteStruct) (types.NodeId, bool, bool, bool, types.Payment) {
	var nextNodeId types.NodeId = 0
	var payNextId types.NodeId = 0
	var payment types.Payment
	var thresholdFailed bool
	var accessFailed bool
	mainOriginatorId := request.OriginatorId
	chunkId := request.ChunkId
	lastDistance := firstNodeId.ToInt32() ^ chunkId.ToInt32()
	//fmt.Printf("\n last distance is : %d, chunk is: %d, first is: %d", lastDistance, chunkId, firstNodeId)
	//fmt.Printf("\n which bucket: %d \n", 16-BitLength(chunkId^firstNodeId))

	currDist := lastDistance
	payDist := lastDistance

	//var lockedEdges []int

	bin := constants.GetBits() - general.BitLength(firstNodeId.ToInt32()^chunkId.ToInt32())
	firstNodeAdjIds := graph.GetNodeAdj(firstNodeId)

	for _, nodeId := range firstNodeAdjIds[bin] {
		dist := nodeId.ToInt32() ^ chunkId.ToInt32()
		if general.BitLength(dist) >= general.BitLength(lastDistance) {
			continue
		}
		if dist >= currDist {
			continue
		}
		// This means the node is now actively trying to communicate with the other node
		if constants.GetEdgeLock() {
			graph.LockEdge(firstNodeId, nodeId)
		}
		if !isThresholdFailed(firstNodeId, nodeId, chunkId, graph, request) {
			thresholdFailed = false
			if constants.IsRetryWithAnotherPeer() {
				if reroute := rerouteStruct.GetRerouteMap(mainOriginatorId).Reroute; reroute != nil {
					if general.Contains(reroute, nodeId) {
						if constants.GetEdgeLock() {
							graph.UnlockEdge(firstNodeId, nodeId)
						}
						continue // skips node that's been part of a failed route before
					}
				}
			}

			if constants.GetEdgeLock() {
				if nextNodeId != 0 {
					graph.UnlockEdge(firstNodeId, nextNodeId)
				}
				if payNextId != 0 {
					graph.UnlockEdge(firstNodeId, payNextId)
					payNextId = 0 // IMPORTANT!
				}
			}
			currDist = dist
			nextNodeId = nodeId

		} else {
			thresholdFailed = true
			if constants.GetPaymentEnabled() {
				if dist < payDist && nextNodeId == 0 {
					if constants.GetEdgeLock() && payNextId != 0 {
						graph.UnlockEdge(firstNodeId, payNextId)
					}
					payDist = dist
					payNextId = nodeId
				} else {
					if constants.GetEdgeLock() {
						graph.UnlockEdge(firstNodeId, nodeId)
					}
				}
			} else {
				if constants.GetEdgeLock() {
					graph.UnlockEdge(firstNodeId, nodeId)
				}
			}
		}
	}

	if nextNodeId != 0 {
		thresholdFailed = false
		accessFailed = false
	} else if !thresholdFailed {
		accessFailed = true
		//nextNodeId = -2 // Access Failed
	} else {
		//nextNodeId = -1 // Threshold Failed
	}

	if constants.GetPaymentEnabled() && payNextId != 0 {
		accessFailed = false

		if firstNodeId == mainOriginatorId {
			payment.IsOriginator = true
		}

		payment.FirstNodeId = firstNodeId
		payment.PayNextId = payNextId
		payment.ChunkId = chunkId
		nextNodeId = payNextId

		if constants.IsOnlyOriginatorPays() {
			if firstNodeId != mainOriginatorId {
				nextNodeId = -1
				payNextId = 0
			}
		} else if constants.IsPayIfOrigPays() {
			if !prevNodePaid && firstNodeId != mainOriginatorId {
				thresholdFailed = true
				nextNodeId = -1
				payNextId = 0
			}
		} else {
			thresholdFailed = false
		}
	}

	//// unlocks all nodes except the nextNodeId lock
	//if constants.GetEdgeLock() {
	//	for _, nodeId := range lockedEdges {
	//		if nodeId != nextNodeId {
	//			graph.UnlockEdge(firstNodeId, nodeId)
	//		}
	//	}
	//}

	// TODO: Not used anymore since forgiveness is done in the isThresholdFailed check
	//if constants.GetPaymentEnabled() {
	//out:
	//	for i, item := range thresholdList {
	//		for _, nodeId := range item {
	//			if nodeId == payNextId {
	//				if constants.IsPayIfOrigPays() {
	//					if firstNodeId == mainOriginatorId {
	//						thresholdList = append(thresholdList[:i], thresholdList[i+1:]...)
	//					}
	//				} else {
	//					thresholdList = append(thresholdList[:i], thresholdList[i+1:]...)
	//				}
	//				break out
	//			}
	//		}
	//	}
	//}

	if payment != (types.Payment{}) {
		prevNodePaid = true
	} else {
		prevNodePaid = false
	}

	if nextNodeId != 0 {
		// fmt.Println("Next node is: ", nextNodeId)
	}

	return nextNodeId, thresholdFailed, accessFailed, prevNodePaid, payment
}

// ConsumeTask cacheDict is map of nodes containing an array of maps with key as a chunkAddr and a popularity counter
func ConsumeTask(request types.Request, graph *types.Graph, rerouteStruct types.RerouteStruct, cacheStruct types.CacheStruct) types.RequestResult {
	chunkId := request.ChunkId
	respNodes := request.RespNodes
	mainOriginatorId := request.OriginatorId
	curNextNodeId := request.OriginatorId
	requestResult := types.RequestResult{
		Route:       []types.NodeId{mainOriginatorId},
		ChunkId:     chunkId,
		PaymentList: []types.Payment{},
	}
	found := false
	accessFailed := false
	thresholdFailed := false
	var nextNodeId types.NodeId
	var payment types.Payment
	var prevNodePaid bool

	if constants.IsPayIfOrigPays() {
		prevNodePaid = true
	}
	if general.ArrContains(respNodes, mainOriginatorId) {
		// originator has the chunk --> chunk is found
	} else {
	out:
		for !general.ArrContains(respNodes, curNextNodeId) {

			//fmt.Printf("\n orig: %d, chunk_id: %d", mainOriginatorId, chunkId)
			//nextNodeId, thresholdList, _, accessFailed, payment, prevNodePaid = getNext(originatorId, chunkId, graph, mainOriginatorId, prevNodePaid, rerouteMap)

			nextNodeId, thresholdFailed, accessFailed, prevNodePaid, payment = getNext(request, curNextNodeId, prevNodePaid, graph, rerouteStruct)

			//if nextNodeId == -2 {
			//	// Access Failed
			//	fmt.Println("Access Failed")
			//}

			if payment != (types.Payment{}) {
				requestResult.PaymentList = append(requestResult.PaymentList, payment)
			}

			if nextNodeId != 0 {
				requestResult.Route = append(requestResult.Route, nextNodeId)
			}
			// if not isinstance(next_node, int), originale versjonen
			if !thresholdFailed && !accessFailed {
				if general.ArrContains(respNodes, nextNodeId) {
					//fmt.Println("is not in cache")
					found = true
					break out
				}
				if constants.IsCacheEnabled() {
					//if ok := cacheStruct.Contains(nextNodeId, chunkId); ok {
					//	found = true
					//	foundByCaching = true
					//	break out
					//}
					node := graph.GetNode(nextNodeId)
					node.Mutex.Lock()
					if _, ok := node.CacheMap[chunkId]; ok {
						//fmt.Println("is in cache")
						requestResult.FoundByCaching = true
						found = true
						node.Mutex.Unlock()
						break out
					}
					node.Mutex.Unlock()
				}
				// NOTE !
				curNextNodeId = nextNodeId
			} else {
				break out
			}
		}
	}

	if constants.IsForwarderPayForceOriginatorToPay() {
		//if nextNodeId != -2 {
		if !accessFailed {
			if len(requestResult.PaymentList) > 0 {
				firstPayment := requestResult.PaymentList[0]
				if !firstPayment.IsOriginator {
					for i := range requestResult.Route {
						p := types.Payment{FirstNodeId: requestResult.Route[i], PayNextId: requestResult.Route[i+1], ChunkId: requestResult.ChunkId}

						for _, tmp := range requestResult.PaymentList {
							if p.PayNextId == tmp.PayNextId && p.FirstNodeId == tmp.FirstNodeId && p.ChunkId == tmp.ChunkId {
								break
							}
						}
						// payment is now not in paymentList
						if i == 0 {
							p.IsOriginator = true
						}
						if i != len(requestResult.Route)-1 {
							if i != len(requestResult.Route)-2 {
								requestResult.PaymentList = append(requestResult.PaymentList[:i+1], requestResult.PaymentList[i:]...)
							}
							requestResult.PaymentList[i] = p
						} else {
							continue
						}

					}
				} else { // firstPayment = originator
					for i := range requestResult.Route[1:] {
						p := types.Payment{FirstNodeId: requestResult.Route[i], PayNextId: requestResult.Route[i+1], ChunkId: requestResult.ChunkId}
						for _, tmp := range requestResult.PaymentList {
							if p.PayNextId == tmp.PayNextId && p.FirstNodeId == tmp.FirstNodeId && p.ChunkId == tmp.ChunkId {
								break
							}
						}
						// payment is now not in paymentList
						if i != len(requestResult.Route)-1 {
							if i != len(requestResult.Route)-2 {
								requestResult.PaymentList = append(requestResult.PaymentList[:i+1], requestResult.PaymentList[i:]...)
							}
							requestResult.PaymentList[i] = p
						} else {
							continue
						}
					}
				}
			}
		} else {
			requestResult.PaymentList = []types.Payment{}
		}
	}
	requestResult.Found = found
	requestResult.AccessFailed = accessFailed
	requestResult.ThresholdFailed = thresholdFailed

	return requestResult
}

func getProximityChunk(firstNodeId types.NodeId, chunkId types.ChunkId) int {
	retVal := constants.GetBits() - general.BitLength(firstNodeId.ToInt32()^chunkId.ToInt32())
	if retVal <= constants.GetMaxProximityOrder() {
		return retVal
	} else {
		return constants.GetMaxProximityOrder()
	}
}

func PeerPriceChunk(firstNodeId types.NodeId, chunkId types.ChunkId) int {
	val := (constants.GetMaxProximityOrder() - getProximityChunk(firstNodeId, chunkId) + 1) * constants.GetPrice()
	return val
}

func CreateDownloadersList(g *types.Graph) []types.NodeId {
	//fmt.Println("Creating downloaders list...")

	downloadersList := types.Choice(g.NodeIds, constants.GetOriginators())

	//fmt.Println("Downloaders list create...!")
	return downloadersList
}

func CreateNodesList(g *types.Graph) []types.NodeId {
	//fmt.Println("Creating nodes list...")
	nodesValue := g.NodeIds
	//fmt.Println("NodesMap list create...!")
	return nodesValue
}

// TODO: Not used in original
//func getBin(src int, dest int, index int) int {
//	distance := src ^ dest
//	result := index
//	for distance > 0 {
//		distance >>= 1
//		result -= 1
//	}
//	return result
//}

// TODO: Not used in original
//func whichPowerTwo(rangeAddress int) int {
//	return BitLength(rangeAddress) - 1
//}

// TODO: Not used in original
//func MakeFiles() []int {
//	fmt.Println("Making files...")
//	var filesList []int
//
//	for i := 0; i <= ct.constants.GetOriginators(); i++ {
//		// chunksList := choice(ct.constants.GetChunks(), ct.constants.GetRangeAddress())
//		// filesList = append(chunksList)
//		fmt.Println(i)
//	}
//	// Gets all constants
//	consts := ct.constants
//
//	for i := 0; i <= consts.GetOriginators(); i++ {
//		chunksList := rand.Perm(consts.GetChunks())
//		filesList = append(chunksList)
//	}
//	fmt.Println("Files made!")
//	return filesList
//}

// TODO: Not used in original
//func (net *Network) PushSync(fileName string, files []string) {
//	fmt.Println("Pushing sync...")
//	if net == nil {
//		fmt.Println("Network is nil!")
//		return
//	}
//	nodes := net.nodes
//	for i := range nodes {
//		fmt.Println(nodes[i].id)
//	}
//
//	fmt.Println("Pushing sync finished...")
//}
