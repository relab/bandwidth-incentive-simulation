package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
)

func Graph(state *types.State, requestResult types.RequestResult, curTimeStep int) types.Output {
	chunkId := requestResult.ChunkId
	route := requestResult.Route
	paymentsList := requestResult.PaymentList
	var routeWithPrice types.RouteWithPrice
	var paymentWithPrice types.PaymentWithPrice
	var output types.Output

	if constants.GetPaymentEnabled() {
		for _, payment := range paymentsList {
			if payment != (types.Payment{}) {
				if !payment.IsOriginator {
					edgeData1 := state.Graph.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
					edgeData2 := state.Graph.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
					price := utils.PeerPriceChunk(payment.PayNextId, payment.ChunkId)
					actualPrice := edgeData1.A2B - edgeData2.A2B + price
					if constants.IsPayOnlyForCurrentRequest() {
						actualPrice = price
					}
					if actualPrice < 0 {
						continue
					} else {
						if !constants.IsPayOnlyForCurrentRequest() {
							newEdgeData1 := edgeData1
							newEdgeData1.A2B = 0
							state.Graph.SetEdgeData(payment.FirstNodeId, payment.PayNextId, newEdgeData1)

							newEdgeData2 := edgeData2
							newEdgeData2.A2B = 0
							state.Graph.SetEdgeData(payment.PayNextId, payment.FirstNodeId, newEdgeData2)
						}
					}
					// fmt.Println("Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", actualPrice)
					paymentWithPrice = types.PaymentWithPrice{Payment: payment, Price: actualPrice}
					output.PaymentsWithPrice = append(output.PaymentsWithPrice, paymentWithPrice)

				} else {
					edgeData1 := state.Graph.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
					edgeData2 := state.Graph.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
					price := utils.PeerPriceChunk(payment.PayNextId, payment.ChunkId)
					actualPrice := edgeData1.A2B - edgeData2.A2B + price
					if constants.IsPayOnlyForCurrentRequest() {
						actualPrice = price
					}
					if actualPrice < 0 {
						continue
					} else {
						if !constants.IsPayOnlyForCurrentRequest() {
							newEdgeData1 := edgeData1
							newEdgeData1.A2B = 0
							state.Graph.SetEdgeData(payment.FirstNodeId, payment.PayNextId, newEdgeData1)

							newEdgeData2 := edgeData2
							newEdgeData2.A2B = 0
							state.Graph.SetEdgeData(payment.PayNextId, payment.FirstNodeId, newEdgeData2)
						}
					}
					//fmt.Println("-1", "Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", actualPrice) //Means that the first one is the originator
					paymentWithPrice = types.PaymentWithPrice{Payment: payment, Price: actualPrice}
					output.PaymentsWithPrice = append(output.PaymentsWithPrice, paymentWithPrice)
				}
			}
		}
	}

	if requestResult.Found {
		for i := 0; i < len(route)-1; i++ {
			requesterNode := route[i]
			providerNode := route[i+1]
			price := utils.PeerPriceChunk(providerNode, chunkId)
			edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
			//edgeData1.A2B += price
			newEdgeData := edgeData
			newEdgeData.A2B += price
			state.Graph.SetEdgeData(requesterNode, providerNode, newEdgeData)

			if constants.GetMaxPOCheckEnabled() {
				//routeWithPrice = append(routeWithPrice, requesterNode)
				//routeWithPrice = append(routeWithPrice, price)
				//routeWithPrice = append(routeWithPrice, providerNode)
				routeWithPrice = types.RouteWithPrice{RequesterNode: requesterNode, ProviderNode: providerNode, Price: price}
				output.RoutesWithPrice = append(output.RoutesWithPrice, routeWithPrice)
			}
		}
		if constants.GetMaxPOCheckEnabled() {
			//fmt.Println("RequestResult with price ", routeWithPrice)
		}

	}
	// TODO: Not needed, forgiveness is done during routing
	//if constants.GetThresholdEnabled() && constants.IsForgivenessEnabled() {
	//	thresholdFailedLists := policyInput.ThresholdFailedLists
	//	if len(thresholdFailedLists) > 0 {
	//		for _, thresholdFailedL := range thresholdFailedLists {
	//			if len(thresholdFailedL) > 0 {
	//				for _, couple := range thresholdFailedL {
	//					requesterNode := couple[0]
	//					providerNode := couple[1]
	//					edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
	//					passedTime := (curTimeStep - edgeData.Last) / constants.GetRequestsPerSecond()
	//					if passedTime > 0 {
	//						refreshRate := constants.GetRefreshRate()
	//						if constants.IsAdjustableThreshold() {
	//							refreshRate = int(math.Ceil(float64(edgeData.Threshold / 2)))
	//						}
	//						removedDeptAmount := passedTime * refreshRate
	//						newEdgeData := edgeData
	//						newEdgeData.A2B -= removedDeptAmount
	//						if newEdgeData.A2B < 0 {
	//							newEdgeData.A2B = 0
	//						}
	//						newEdgeData.Last = curTimeStep
	//						state.Graph.SetEdgeData(requesterNode, providerNode, newEdgeData)
	//					}
	//				}
	//			}
	//		}
	//	}
	//}

	// Unlocks all the edges between the nodes in the route
	if constants.GetEdgeLock() {
		for i := 0; i < len(route)-1; i++ {
			state.Graph.UnlockEdge(route[i], route[i+1])
		}
	}

	return output
}
