package policy

import (
	"fmt"
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
)

func UpdateSuccessfulFound(prevState State, policyInput Policy) State {
	oldSuccessCounter := prevState.SuccessfulFound
	newSuccessCounter := oldSuccessCounter
	if policyInput.Found {
		newSuccessCounter++
	}
	return prevState
}

func UpdateFailedRequestsThreshold(prevState State, policyInput Policy) State {
	oldFailedCounter := prevState.FailedRequestsThreshold
	newFailedCounter := oldFailedCounter
	found := policyInput.Found
	// thresholdFailed := policyInput.thresholdFailed
	accessFailed := policyInput.AccessFailed
	if !found && !accessFailed {
		newFailedCounter++
	}
	return prevState
}

func UpdateFailedRequestsAccess(prevState State, policyInput Policy) State {
	oldFailedAccessCounter := prevState.FailedRequestsAccess
	accessFailed := policyInput.AccessFailed
	if accessFailed {
		oldFailedAccessCounter++
	}
	return prevState
}

func UpdateOriginatorIndex(prevState State, policyInput Policy) State {
	oldOriginatorIndex := prevState.OriginatorIndex
	newOriginatorIndex := oldOriginatorIndex + 1
	if newOriginatorIndex >= Constants.GetOriginators() {
		newOriginatorIndex = 0
	}
	return prevState
}

// TODO: function convert and dump to file

func UpdateRouteListAndFlush(prevState State, policyInput Policy) State {
	prevState.RouteLists = append(prevState.RouteLists, policyInput.Route)
	currTimestep := prevState.TimeStep + 1
	if currTimestep%6250 == 0 {
		// TODO: call convert_and_dump
		prevState.RouteLists = []Route{}
		return prevState
	}
	return prevState
}

// TODO: Implement this function
func UpdateCacheDictionary(prevState State, policyInput Policy) State {
	return prevState
}

func UpdateRerouteMap(prevState State, policyInput Policy) State {
	rerouteMap := prevState.RerouteMap
	if Constants.IsRetryWithAnotherPeer() {
		route := policyInput.Route
		originator := route[0]
		if !contains(route, -1) && !contains(route, -2) {
			if _, ok := rerouteMap[originator]; ok {
				val := rerouteMap[originator]
				if val[len(val)-1] == route[len(route)-1] {
					//remove rerouteMap[originator]
					delete(rerouteMap, originator)
				}
			}
		} else {
			if len(route) > 3 {
				if _, ok := rerouteMap[originator]; ok {
					val := rerouteMap[originator]
					if !contains(val, route[1]) {
						val = append([]int{route[1]}, val...)
						rerouteMap[originator] = val
					}
				} else {
					rerouteMap[originator] = []int{route[1], route[len(route)-1]}
				}
			}
		}
		if _, ok := rerouteMap[originator]; ok {
			if len(rerouteMap[originator]) > Constants.GetBinSize() {
				delete(rerouteMap, originator)
			}
		}
	}
	return prevState
}

func UpdatePendingDictionary(prevState State, policyInput Policy) State {
	pendingDict := prevState.PendingDict
	if Constants.IsWaitingEnabled() {
		route := policyInput.Route
		originator := route[0]
		if !contains(route, -1) && !contains(route, -2) {
			if _, ok := pendingDict[originator]; ok {
				if pendingDict[originator] == route[len(route)-1] {
					delete(pendingDict, originator)
				}
			}

		} else {
			pendingDict[originator] = route[len(route)-1]
		}
	}
	return prevState
}

func UpdateNetwork(prevState State, policyInput Policy) State {
	network := prevState.Network
	currTinmeStep := prevState.TimeStep + 1
	route := policyInput.Route
	paymentsList := policyInput.PaymentList

	if Constants.GetPaymentEnabled() {
		for _, payment := range paymentsList {
			var p Payment
			if payment != p {
				if payment.FirstNodeId != -1 {
					edgeData1 := network.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
					edgeData2 := network.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
					price := PeerPriceChunk(payment.PayNextId, payment.ChunkId)
					val := edgeData1.A2b - edgeData2.A2b + price
					if Constants.IsPayOnlyForCurrentRequest() {
						val = price
					}
					if val < 0 {
						continue
					} else {
						if !Constants.IsPayOnlyForCurrentRequest() {
							edgeData1.A2b = 0
							edgeData2.A2b = 0
						}
					}
					fmt.Println("Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", val)
				} else {
					edgeData1 := network.GetEdgeData(payment.FirstNodeId, payment.PayNextId)
					edgeData2 := network.GetEdgeData(payment.PayNextId, payment.FirstNodeId)
					price := PeerPriceChunk(payment.PayNextId, payment.ChunkId)
					val := edgeData1.A2b - edgeData2.A2b + price
					if Constants.IsPayOnlyForCurrentRequest() {
						val = price
					}
					if val < 0 {
						continue
					} else {
						if !Constants.IsPayOnlyForCurrentRequest() {
							edgeData1.A2b = 0
							edgeData2.A2b = 0
						}
					}
					fmt.Println("-1", "Payment from ", payment.FirstNodeId, " to ", payment.PayNextId, " for chunk ", payment.ChunkId, " with price ", val) //Means that the first one is the originator
				}
			}
		}
	}
	if !contains(route, -1) && !contains(route, -2) {
		routeWithPrice := []int{}
		if contains(route, -3) {
			chunkId := route[len(route)-2]
			for i := 0; i < len(route)-3; i++ {
				requesterNode := route[i]
				providerNode := route[i+1]
				price := PeerPriceChunk(providerNode, chunkId)
				edgeData1 := network.GetEdgeData(requesterNode, providerNode)
				edgeData1.A2b += price
				if Constants.GetMaxPoCheckEnabled() {
					routeWithPrice = append(routeWithPrice, requesterNode)
					routeWithPrice = append(routeWithPrice, price)
					routeWithPrice = append(routeWithPrice, providerNode)
				}
			}
			if Constants.GetMaxPoCheckEnabled() {
				fmt.Println("Route with price ", routeWithPrice)
			}
		} else {
			chunkId := route[len(route)-1]
			for i := 0; i < len(route)-2; i++ {
				requesterNode := route[i]
				providerNode := route[i+1]
				price := PeerPriceChunk(providerNode, chunkId)
				edgeData1 := network.GetEdgeData(requesterNode, providerNode)
				edgeData1.A2b += price
				if Constants.GetMaxPoCheckEnabled() {
					routeWithPrice = append(routeWithPrice, requesterNode)
					routeWithPrice = append(routeWithPrice, price)
					routeWithPrice = append(routeWithPrice, providerNode)
				}
			}
			if Constants.GetMaxPoCheckEnabled() {
				fmt.Println("Route with price ", routeWithPrice)
			}
		}
	}
	if Constants.GetThresholdEnabled() && Constants.IsForgivenessEnabled() {
		thresholdFailedLists := policyInput.ThresholdFailed
		if len(thresholdFailedLists) > 0 {
			for _, thresholdFailedL := range thresholdFailedLists {
				if len(thresholdFailedL) > 0 {
					for _, couple := range thresholdFailedL {
						requesterNode := couple[0].Id
						providerNode := couple[1].Id
						edgeData1 := network.GetEdgeData(requesterNode, providerNode)
						passedTime := (currTinmeStep - edgeData1.Last) / Constants.GetRequestsPerSecond()
						if passedTime > 0 {
							refreshRate := Constants.GetRefreshRate()
							//if Constants.IsAdjustableThreshold() {
							//	refreshRate = int(math.Ceil(edgeData1.Threshold / 2))
							//}
							removedDeptAmount := passedTime * refreshRate
							edgeData1.A2b -= removedDeptAmount
							if edgeData1.A2b < 0 {
								edgeData1.A2b = 0
							}
							edgeData1.Last = currTinmeStep
						}
					}
				}
			}
		}
	}
	return prevState
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
