package policy

// TODO: keeping this here to compare with other implmentation later
//func oldFindResponsibleNodes(nodesId []int, chunkAdd int) []int {
//	//numNodes := Constants.GetBits()
//	numNodes := 100
//	distances := make([]int, 0, numNodes)
//	var distance int
//	nodesMap := make(map[int]int)
//	returnNodes := make([]int, 4)
//
//	closestNodes := BinarySearchClosest(nodesId, chunkAdd, numNodes)
//
//	for _, nodeId := range closestNodes {
//		distance = nodeId ^ chunkAdd
//		// fmt.Println(distance, nodeId)
//		distances = append(distances, distance)
//		nodesMap[distance] = nodeId
//	}
//
//	sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })
//
//	for i := 0; i < 4; i++ {
//		distance = distances[i]
//		returnNodes[i] = nodesMap[distance]
//	}
//	return returnNodes
//}

//func SendRequest(prevState *State, index int) (bool, Route, [][]Threshold, bool, []Payment) {
//	// Gets one random chunkId from the range of addresses
//	chunkId := rand.Intn(Constants.GetRangeAddress() - 1)
//
//	if Constants.IsPreferredChunksEnabled() {
//		var random float32
//		numPreferredChunks := 1000
//		random = rand.Float32()
//		if float32(random) <= 0.5 {
//			chunkId = rand.Intn(numPreferredChunks)
//		} else {
//			chunkId = rand.Intn(Constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
//		}
//	}
//
//	responsibleNodes := prevState.Graph.FindResponsibleNodes(chunkId)
//	originatorId := prevState.Originators[prevState.OriginatorIndex]
//	//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]
//
//	pendingNodeId := prevState.PendingStruct.GetPending(originatorId)
//	if pendingNodeId != -1 {
//		chunkId = prevState.PendingStruct.GetPending(originatorId)
//		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
//	}
//	//if _, ok := prevState.PendingMap[originatorId]; ok {
//	//	chunkId = prevState.PendingMap[originatorId]
//	//	responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
//	//}
//	reroute := prevState.RerouteStruct.GetRerouteMap(originatorId)
//	if reroute != nil {
//		chunkId = reroute[len(reroute)-1]
//		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
//	}
//	//if _, ok := prevState.RerouteMap[originatorId]; ok {
//	//	chunkId = prevState.RerouteMap[originatorId][len(prevState.RerouteMap[originatorId])-1]
//	//	responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
//	//}
//
//	request := Request{OriginatorId: originatorId, ChunkId: chunkId, RespNodes: responsibleNodes}
//
//	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.Graph, prevState.RerouteStruct, prevState.CacheStruct)
//
//	return found, route, thresholdFailed, accessFailed, paymentsList
//}

//func CacheMap(state *types.State, policyInput types.RequestResult) types.CacheStruct {
//	chunkId := 0
//
//	if constants.Constants.IsCacheEnabled() {
//		route := policyInput.Route
//		if general.Contains(route, -3) {
//			// -3 means found by caching
//			state.CacheStruct.CacheHits++
//			chunkId = route[len(route)-2]
//		} else {
//			chunkId = route[len(route)-1]
//		}
//		if !general.Contains(route, -1) && !general.Contains(route, -2) {
//			if general.Contains(route, -3) {
//				for i := 0; i < len(route)-3; i++ {
//					nodeId := route[i]
//					state.CacheStruct.AddToCache(nodeId, chunkId)
//					node := state.Graph.GetNode(nodeId)
//					node.Mutex.Lock()
//					cacheMap := node.CacheMap
//					if cacheMap != nil {
//						if _, ok := cacheMap[chunkId]; ok {
//							cacheMap[chunkId]++
//						} else {
//							cacheMap[chunkId] = 1
//						}
//					} else {
//						node.CacheMap = map[int]int{node.Id: 1}
//					}
//					node.Mutex.Unlock()
//				}
//			} else {
//				for i := 0; i < len(route)-2; i++ {
//					nodeId := route[i]
//					state.CacheStruct.AddToCache(nodeId, chunkId)
//					node := state.Graph.GetNode(nodeId)
//					node.Mutex.Lock()
//					cacheMap := node.CacheMap
//					if cacheMap != nil {
//						if _, ok := cacheMap[chunkId]; ok {
//							cacheMap[chunkId]++
//						} else {
//							cacheMap[chunkId] = 1
//						}
//					} else {
//						node.CacheMap = map[int]int{node.Id: 1}
//					}
//					node.Mutex.Unlock()
//				}
//			}
//		}
//	}
//	//state.CacheStruct = cacheStruct
//	return state.CacheStruct
//}
//
//func RerouteMap(state *types.State, policyInput types.RequestResult) types.RerouteStruct {
//	//rerouteStruct := state.RerouteStruct
//	if constants.Constants.IsRetryWithAnotherPeer() {
//		route := policyInput.Route
//		originator := route[0]
//		if !general.Contains(route, -1) && !general.Contains(route, -2) {
//			reroute := state.RerouteStruct.GetRerouteMap(originator)
//			if reroute != nil {
//				if reroute[len(reroute)-1] == route[len(route)-1] {
//					//remove rerouteMap[originator]
//					state.RerouteStruct.DeleteReroute(originator)
//				}
//			}
//			//if _, ok := rerouteMap[originator]; ok {
//			//	val := rerouteMap[originator]
//			//	if val[len(val)-1] == route[len(route)-1] {
//			//		//remove rerouteMap[originator]
//			//		delete(rerouteMap, originator)
//			//	}
//			//}
//		} else {
//			if len(route) > 3 {
//				reroute := state.RerouteStruct.GetRerouteMap(originator)
//				state.RerouteStruct.RerouteMutex.Lock()
//				if reroute != nil {
//					if !general.Contains(reroute, route[1]) {
//						reroute = append([]int{route[1]}, reroute...)
//						state.RerouteStruct.RerouteMap[originator] = reroute
//					}
//				} else {
//					state.RerouteStruct.RerouteMap[originator] = []int{route[1], route[len(route)-1]}
//				}
//				state.RerouteStruct.RerouteMutex.Unlock()
//
//				//if _, ok := rerouteMap[originator]; ok {
//				//	val := rerouteMap[originator]
//				//	if !Contains(val, route[1]) {
//				//		val = append([]int{route[1]}, val...)
//				//		rerouteMap[originator] = val
//				//	}
//				//} else {
//				//	rerouteMap[originator] = []int{route[1], route[len(route)-1]}
//				//}
//			}
//		}
//		reroute := state.RerouteStruct.GetRerouteMap(originator)
//		if reroute != nil {
//			if len(reroute) > constants.Constants.GetBinSize() {
//				state.RerouteStruct.DeleteReroute(originator)
//			}
//		}
//		//if _, ok := rerouteMap[originator]; ok {
//		//	if len(rerouteMap[originator]) > Constants.GetBinSize() {
//		//		delete(rerouteMap, originator)
//		//	}
//		//}
//	}
//	//state.RerouteStruct = rerouteStruct
//	return state.RerouteStruct
//}
//
//func PendingMap(state *types.State, policyInput types.RequestResult) types.PendingStruct {
//	//pendingStruct := state.PendingStruct
//	if constants.Constants.IsWaitingEnabled() {
//		route := policyInput.Route
//		originator := route[0]
//		chunkId := route[len(route)-1]
//
//		if constants.Constants.IsRetryWithAnotherPeer() {
//			if !general.Contains(route, -1) && !general.Contains(route, -2) {
//				pendingNodeId := state.PendingStruct.GetPending(originator).NodeId
//				if pendingNodeId != -1 {
//					if pendingNodeId == chunkId {
//						// remove the pending request
//						state.PendingStruct.DeletePending(originator)
//					}
//				}
//				//if _, ok := pendingMap[originator]; ok {
//				//	if pendingMap[originator] == route[len(route)-1] {
//				//		delete(pendingMap, originator)
//				//	}
//				//}
//			} else {
//				pendingNode := state.PendingStruct.GetPending(originator)
//				if pendingNode.NodeId != -1 {
//					if pendingNode.PendingCounter < 100 {
//						state.PendingStruct.IncrementPending(originator)
//					} else {
//						// remove the pending request
//						state.PendingStruct.DeletePending(originator)
//					}
//				} else {
//					// add the pending request
//					state.PendingStruct.AddPending(originator, chunkId)
//				}
//			}
//			//} else {
//			//	pendingMap[originator] = route[len(route)-1]
//			//}
//			//threshold failed
//		} else if general.Contains(route, -1) {
//			state.PendingStruct.AddPending(originator, chunkId)
//		}
//	}
//	//state.PendingStruct = pendingStruct
//	return state.PendingStruct
//}
