package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/output"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
)

func Graph(state *types.State, requestResult types.RequestResult, curTimeStep int) output.Route {
	chunkId := requestResult.ChunkId
	route := requestResult.Route
	paymentsList := requestResult.PaymentList
	var nodePairWithPrice types.NodePairWithPrice
	var paymentWithPrice types.PaymentWithPrice
	var output output.Route

	if config.GetPaymentEnabled() && requestResult.Found {
		for _, payment := range paymentsList {
			if !payment.IsNil() {
				edgeData1 := state.Graph.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
				edgeData2 := state.Graph.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
				price := utils.PeerPriceChunk(payment.PayNextId, payment.ChunkId)
				actualPrice := edgeData1.A2B - edgeData2.A2B + price
				if config.IsPayOnlyForCurrentRequest() {
					actualPrice = price
				}
				if actualPrice < 0 {
					continue
				} else {
					if !config.IsPayOnlyForCurrentRequest() {
						state.Graph.SetEdgeA2B(payment.FirstNodeId, payment.PayNextId, 0)
						state.Graph.SetEdgeA2B(payment.PayNextId, payment.FirstNodeId, 0)
					} else {
						// Important fix: Reduce debt here, since it debt will be added again below.
						// Idea is, paying for the current request should not effect the edge balance.
						state.Graph.SetEdgeA2B(payment.FirstNodeId, payment.PayNextId, edgeData1.A2B-price)
					}
				}
				// fmt.Println("Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", actualPrice)
				paymentWithPrice = types.PaymentWithPrice{Payment: payment, Price: actualPrice}
				output.PaymentsWithPrices = append(output.PaymentsWithPrices, paymentWithPrice)
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
			state.Graph.SetEdgeA2B(requesterNode, providerNode, edgeData.A2B+price)

			if config.GetMaxPOCheckEnabled() {
				nodePairWithPrice = types.NodePairWithPrice{RequesterNode: requesterNode, ProviderNode: providerNode, Price: price}
				output.RouteWithPrices = append(output.RouteWithPrices, nodePairWithPrice)
			}
		}
	}

	// Unlocks all the edges between the nodes in the route
	if config.IsEdgeLock() {
		for i := 0; i < len(route)-1; i++ {
			state.Graph.UnlockEdge(route[i], route[i+1])
		}
	}

	return output
}
