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
			}
		}
	}

	// Update edges debt based on price
	if requestResult.Found {
		for i := 0; i < len(route)-1; i++ {
			requesterNode := route[i]
			providerNode := route[i+1]
			price := utils.PeerPriceChunk(providerNode, chunkId)
			edgeData := state.Graph.GetEdgeData(requesterNode, providerNode)
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

	// Unlocks all the edges between the nodes in the route
	if constants.GetEdgeLock() {
		for i := 0; i < len(route)-1; i++ {
			state.Graph.UnlockEdge(route[i], route[i+1])
		}
	}

	return output
}
