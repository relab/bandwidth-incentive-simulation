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

func UpdateRerouteDictionary(prevState State, policyInput Policy) State {
	rerouteDict := prevState.RerouteDict
	if Constants.IsRetryWithAnotherPeer() {
		route := policyInput.Route
		originator := route[0]
		if !contains(route, -1) && !contains(route, -2) {
			if _, ok := rerouteDict[originator]; ok {
				val := rerouteDict[originator]
				if val[len(val)-1] == route[len(route)-1] {
					//remove rerouteDict[originator]
					delete(rerouteDict, originator)
				}
			}
		} else {
			if len(route) > 3 {
				if _, ok := rerouteDict[originator]; ok {
					val := rerouteDict[originator]
					if !contains(val, route[1]) {
						val = append([]int{route[1]}, val...)
						rerouteDict[originator] = val
					}
				} else {
					rerouteDict[originator] = []int{route[1], route[len(route)-1]}
				}
			}
		}
		if _, ok := rerouteDict[originator]; ok {
			if len(rerouteDict[originator]) > Constants.GetBinSize() {
				delete(rerouteDict, originator)
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
	//currTinmeStep := prevState.TimeStep + 1
	//route := policyInput.Route
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
