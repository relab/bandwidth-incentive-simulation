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

	for i := 0; i < numPossibleChunks; i++ {
		chunkId := int32(i)
		closestNodes := types.BinarySearchClosest(nodesId, chunkId, numNodesSearch)
		distances := make([]int32, len(closestNodes))

		for j, nodeId := range closestNodes {
			distances[j] = nodeId.ToInt32() ^ chunkId
		}

		sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })

		for k := 0; k < 4; k++ {
			result[chunkId][k] = types.NodeId(distances[k] ^ chunkId) // this results in the nodeId again
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
			for _, otherNodeId := range adjItems {
				threshold := general.BitLength(nodeId.ToInt32() ^ otherNodeId.ToInt32())
				epoch := constants.GetEpoch()
				attrs := types.EdgeAttrs{A2B: 0, LastEpoch: epoch, Threshold: threshold}
				err := graph.AddEdge(node.Id, otherNodeId, attrs)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	//fmt.Println("Graph network is created.")
	return graph, nil
}

func isThresholdFailed(firstNodeId types.NodeId, secondNodeId types.NodeId, graph *types.Graph, request types.Request) bool {
	if constants.GetThresholdEnabled() {
		edgeDataFirst := graph.GetEdgeData(firstNodeId, secondNodeId)
		p2pFirst := edgeDataFirst.A2B
		edgeDataSecond := graph.GetEdgeData(secondNodeId, firstNodeId)
		p2pSecond := edgeDataSecond.A2B

		threshold := constants.GetThreshold()
		if constants.IsAdjustableThreshold() {
			threshold = edgeDataFirst.Threshold
		}

		peerPriceChunk := PeerPriceChunk(secondNodeId, request.ChunkId)
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
		if !isThresholdFailed(firstNodeId, nodeId, graph, request) {
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
				if !nextNodeId.IsNil() {
					graph.UnlockEdge(firstNodeId, nextNodeId)
				}
				if !payNextId.IsNil() {
					graph.UnlockEdge(firstNodeId, payNextId)
					payNextId = 0 // IMPORTANT!
				}
			}
			currDist = dist
			nextNodeId = nodeId

		} else {
			thresholdFailed = true
			if constants.GetPaymentEnabled() {
				if dist < payDist && nextNodeId.IsNil() {
					if constants.GetEdgeLock() && !payNextId.IsNil() {
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

	if !nextNodeId.IsNil() {
		thresholdFailed = false
		accessFailed = false
	} else if !thresholdFailed {
		accessFailed = true
	}

	if constants.GetPaymentEnabled() && !payNextId.IsNil() {
		accessFailed = false

		if firstNodeId == mainOriginatorId {
			payment.IsOriginator = true
		}

		if constants.IsOnlyOriginatorPays() {
			// Only set payment if the firstNode is the MainOriginator
			if payment.IsOriginator {
				payment.FirstNodeId = firstNodeId
				payment.PayNextId = payNextId
				payment.ChunkId = chunkId
				nextNodeId = payNextId
			}

		} else if constants.IsPayIfOrigPays() {
			// Pay if the originator pays or if the previous node has paid
			if payment.IsOriginator || prevNodePaid {
				payment.FirstNodeId = firstNodeId
				payment.PayNextId = payNextId
				payment.ChunkId = chunkId
				nextNodeId = payNextId
				thresholdFailed = false

			} else {
				thresholdFailed = true
			}

		} else {
			// Always pays
			payment.FirstNodeId = firstNodeId
			payment.PayNextId = payNextId
			payment.ChunkId = chunkId
			nextNodeId = payNextId
			thresholdFailed = false
		}
	}

	if !payment.IsNil() {
		prevNodePaid = true
	} else {
		prevNodePaid = false
	}

	if !nextNodeId.IsNil() {
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
		found = true
	} else {
	out:
		for !general.ArrContains(respNodes, curNextNodeId) {
			//fmt.Printf("\n orig: %d, chunk_id: %d", mainOriginatorId, chunkId)

			nextNodeId, thresholdFailed, accessFailed, prevNodePaid, payment = getNext(request, curNextNodeId, prevNodePaid, graph, rerouteStruct)

			if !payment.IsNil() {
				requestResult.PaymentList = append(requestResult.PaymentList, payment)
			}
			if !nextNodeId.IsNil() {
				requestResult.Route = append(requestResult.Route, nextNodeId)
			}
			if !thresholdFailed && !accessFailed {
				if general.ArrContains(respNodes, nextNodeId) {
					//fmt.Println("is not in cache")
					found = true
					break out
				}
				if constants.IsCacheEnabled() {
					// TODO: if we use the global cacheStruct instead of on the nodes
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

	if constants.IsForwardersPayForceOriginatorToPay() {
		if !accessFailed && len(requestResult.PaymentList) > 0 {

			for i := 0; i < len(requestResult.Route)-1; i++ {
				newPayment := types.Payment{
					FirstNodeId: requestResult.Route[i],
					PayNextId:   requestResult.Route[i+1],
					ChunkId:     requestResult.ChunkId}

				paymentAlreadyExists := false
				for _, tmp := range requestResult.PaymentList {
					if newPayment.FirstNodeId == tmp.FirstNodeId && newPayment.PayNextId == tmp.PayNextId {
						paymentAlreadyExists = true
						break
					}
				}
				if !paymentAlreadyExists {
					if i == 0 {
						newPayment.IsOriginator = true
					}
					if i > len(requestResult.PaymentList)-1 {
						requestResult.PaymentList = append(requestResult.PaymentList, newPayment)
					} else {
						requestResult.PaymentList = append(requestResult.PaymentList[:i+1], requestResult.PaymentList[i:]...)
						requestResult.PaymentList[i] = newPayment
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
