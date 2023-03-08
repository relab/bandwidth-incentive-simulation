package update

//func SuccessfulFound(state *types.State, policyInput types.RequestResult) int32 {
//	if policyInput.Found {
//		return atomic.AddInt32(&state.SuccessfulFound, 1)
//	}
//	return atomic.LoadInt32(&state.SuccessfulFound)
//}

//func FailedRequestsThreshold(state *types.State, policyInput types.RequestResult) int32 {
//	found := policyInput.Found
//	// thresholdFailedList := policyInput.thresholdFailedList
//	accessFailed := policyInput.AccessFailed
//	if !found && !accessFailed {
//		return atomic.AddInt32(&state.FailedRequestsThreshold, 1)
//	}
//	return atomic.LoadInt32(&state.FailedRequestsThreshold)
//}

//func FailedRequestsAccess(state *types.State, policyInput types.RequestResult) int32 {
//	accessFailed := policyInput.AccessFailed
//	if accessFailed {
//		return atomic.AddInt32(&state.FailedRequestsAccess, 1)
//	}
//	return atomic.LoadInt32(&state.FailedRequestsAccess)
//}

//// OriginatorIndex Used by the requestWorker
//func OriginatorIndex(state *types.State, timeStep int32) int32 {
//
//	curOriginatorIndex := atomic.LoadInt32(&state.OriginatorIndex)
//	if constants.Constants.GetSameOriginator() {
//		if (timeStep)%100 == 0 {
//			if int(curOriginatorIndex+1) >= constants.Constants.GetOriginators() {
//				atomic.StoreInt32(&state.OriginatorIndex, 0)
//				return 0
//			} else {
//				return atomic.AddInt32(&state.OriginatorIndex, 1)
//			}
//		}
//	} else {
//		if int(curOriginatorIndex+1) >= constants.Constants.GetOriginators() {
//			atomic.StoreInt32(&state.OriginatorIndex, 0)
//			return 0
//		} else {
//			if constants.Constants.GetSameOriginator() {
//				if atomic.LoadInt32(&state.TimeStep)%100 == 0 {
//					return atomic.AddInt32(&state.OriginatorIndex, 1)
//				}
//			} else {
//				return atomic.AddInt32(&state.OriginatorIndex, 1)
//			}
//		}
//	}
//	return curOriginatorIndex
//	//state.OriginatorIndex = rand.Intn(Constants.GetOriginators() - 1)
//}

//func convertAndDumpToFileStateList(stateList []types.State, curTimeStep int) error {
//	type StateData struct {
//		TimeStep int                 `json:"timestep"`
//		States   []types.StateSubset `json:"states"`
//	}
//	subList := make([]types.StateSubset, len(stateList))
//	for i, state := range stateList {
//		subList[i] = types.StateSubset{
//			OriginatorIndex:         state.OriginatorIndex,
//			PendingMap:              state.PendingStruct.PendingMap,
//			RerouteMap:              state.RerouteStruct.RerouteMap,
//			SuccessfulFound:         state.SuccessfulFound,
//			FailedRequestsThreshold: state.FailedRequestsThreshold,
//			FailedRequestsAccess:    state.FailedRequestsAccess,
//			TimeStep:                state.TimeStep,
//		}
//	}
//	data := StateData{curTimeStep, subList}
//	file, _ := json.MarshalIndent(data, "", "  ")
//	err := os.WriteFile("states.json", file, 0644)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func RouteListAndFlush(state *types.State, policyInput types.RequestResult, curTimeStep int) []types.Route {
//	state.RouteLists = append(state.RouteLists, policyInput.Route)
//	if curTimeStep%6250 == 0 {
//		convertAndDumpToFile(state.RouteLists, curTimeStep)
//		state.RouteLists = []types.Route{}
//	}
//	return state.RouteLists
//}
//
//func StateListAndFlush(state types.State, stateList []types.State, curTimeStep int) []types.State {
//	stateList = append(stateList, state)
//	if curTimeStep%1000 == 0 {
//		convertAndDumpToFileStateList(stateList, curTimeStep)
//		stateList = []types.State{}
//	}
//	return stateList
//}

//func Timestep(prevState *types.State) int {
//	curTimeStep := int(atomic.AddInt32(&prevState.TimeStep, 1))
//	return curTimeStep
//
//}

//func Graph(state *types.State, policyInput types.RequestResult, curTimeStep int) *types.Graph {
//	//network := state.Graph
//	route := policyInput.Route
//	paymentsList := policyInput.PaymentList
//
//	if constants.Constants.GetPaymentEnabled() {
//		for _, payment := range paymentsList {
//			if payment != (types.Payment{}) {
//				if !payment.IsOriginator {
//					edgeData1 := state.Graph.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
//					edgeData2 := state.Graph.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
//					price := utils.PeerPriceChunk(payment.PayNextId, payment.ChunkId)
//					val := edgeData1.A2B - edgeData2.A2B + price
//					if constants.Constants.IsPayOnlyForCurrentRequest() {
//						val = price
//					}
//					if val < 0 {
//						continue
//					} else {
//						if !constants.Constants.IsPayOnlyForCurrentRequest() {
//							//edgeData1.A2B = 0
//							//edgeData2.A2B = 0
//							newEdgeData1 := edgeData1
//							newEdgeData1.A2B = 0
//							state.Graph.SetEdgeData(payment.FirstNodeId, payment.PayNextId, newEdgeData1)
//
//							newEdgeData2 := edgeData2
//							newEdgeData2.A2B = 0
//							state.Graph.SetEdgeData(payment.PayNextId, payment.FirstNodeId, newEdgeData2)
//						}
//					}
//					// fmt.Println("Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", val)
//				} else {
//					edgeData1 := state.Graph.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
//					edgeData2 := state.Graph.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
//					price := utils.PeerPriceChunk(payment.PayNextId, payment.ChunkId)
//					val := edgeData1.A2B - edgeData2.A2B + price
//					if constants.Constants.IsPayOnlyForCurrentRequest() {
//						val = price
//					}
//					if val < 0 {
//						continue
//					} else {
//						if !constants.Constants.IsPayOnlyForCurrentRequest() {
//							//edgeData1.A2B = 0
//							//edgeData2.A2B = 0
//							newEdgeData1 := edgeData1
//							newEdgeData1.A2B = 0
//							state.Graph.SetEdgeData(payment.FirstNodeId, payment.PayNextId, newEdgeData1)
//
//							newEdgeData2 := edgeData2
//							newEdgeData2.A2B = 0
//							state.Graph.SetEdgeData(payment.PayNextId, payment.FirstNodeId, newEdgeData2)
//						}
//					}
//					//fmt.Println("-1", "Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", val) //Means that the first one is the originator
//				}
//			}
//		}
//	}
//	if !general.Contains(route, -1) && !general.Contains(route, -2) {
//		var routeWithPrice []int
//		if general.Contains(route, -3) {
//			chunkId := route[len(route)-2]
//			for i := 0; i < len(route)-3; i++ {
//				requesterNode := route[i]
//				providerNode := route[i+1]
//				price := utils.PeerPriceChunk(providerNode, chunkId)
//				edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
//				//edgeData1.A2B += price
//				newEdgeData := edgeData
//				newEdgeData.A2B += price
//				state.Graph.SetEdgeData(requesterNode, providerNode, newEdgeData)
//
//				if constants.Constants.GetMaxPOCheckEnabled() {
//					routeWithPrice = append(routeWithPrice, requesterNode)
//					routeWithPrice = append(routeWithPrice, price)
//					routeWithPrice = append(routeWithPrice, providerNode)
//				}
//			}
//			if constants.Constants.GetMaxPOCheckEnabled() {
//				//fmt.Println("Route with price ", routeWithPrice)
//			}
//		} else {
//			chunkId := route[len(route)-1]
//			for i := 0; i < len(route)-2; i++ {
//				requesterNode := route[i]
//				providerNode := route[i+1]
//				price := utils.PeerPriceChunk(providerNode, chunkId)
//				edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
//				//edgeData.A2B += price
//				newEdgeData := edgeData
//				newEdgeData.A2B += price
//				state.Graph.SetEdgeData(requesterNode, providerNode, newEdgeData)
//
//				if constants.Constants.GetMaxPOCheckEnabled() {
//					routeWithPrice = append(routeWithPrice, requesterNode)
//					routeWithPrice = append(routeWithPrice, price)
//					routeWithPrice = append(routeWithPrice, providerNode)
//				}
//			}
//			if constants.Constants.GetMaxPOCheckEnabled() {
//				//fmt.Println("Route with price ", routeWithPrice)
//			}
//		}
//	}
//	if constants.Constants.GetThresholdEnabled() && constants.Constants.IsForgivenessEnabled() {
//		thresholdFailedLists := policyInput.ThresholdFailedLists
//		if len(thresholdFailedLists) > 0 {
//			for _, thresholdFailedL := range thresholdFailedLists {
//				if len(thresholdFailedL) > 0 {
//					for _, couple := range thresholdFailedL {
//						requesterNode := couple[0]
//						providerNode := couple[1]
//						edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
//						passedTime := (curTimeStep - edgeData.Last) / constants.Constants.GetRequestsPerSecond()
//						if passedTime > 0 {
//							refreshRate := constants.Constants.GetRefreshRate()
//							if constants.Constants.IsAdjustableThreshold() {
//								refreshRate = int(math.Ceil(float64(edgeData.Threshold / 2)))
//							}
//							removedDeptAmount := passedTime * refreshRate
//							newEdgeData := edgeData
//							newEdgeData.A2B -= removedDeptAmount
//							if newEdgeData.A2B < 0 {
//								newEdgeData.A2B = 0
//							}
//							newEdgeData.Last = curTimeStep
//							state.Graph.SetEdgeData(requesterNode, providerNode, newEdgeData)
//						}
//					}
//				}
//			}
//		}
//	}
//	// Unlocks all the edges between the nodes in the route
//	if constants.Constants.GetEdgeLock() {
//		if !general.Contains(route, -1) && !general.Contains(route, -2) {
//			if general.Contains(route, -3) {
//				for i := 0; i < len(route)-3; i++ {
//					state.Graph.UnlockEdge(route[i], route[i+1])
//				}
//			} else {
//				for i := 0; i < len(route)-2; i++ {
//					state.Graph.UnlockEdge(route[i], route[i+1])
//				}
//			}
//		} else {
//			for i := 0; i < len(route)-3; i++ {
//				state.Graph.UnlockEdge(route[i], route[i+1])
//			}
//		}
//	}
//
//	//state.Graph = network
//	return state.Graph
//}
