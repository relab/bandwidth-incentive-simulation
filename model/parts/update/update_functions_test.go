package update

//
//import (
//	"go-incentive-simulation/model/constants"
//	"go-incentive-simulation/model/parts/policy"
//	"go-incentive-simulation/model/parts/types"
//	"go-incentive-simulation/model/parts/utils"
//	"go-incentive-simulation/model/state"
//	"testing"
//)
//
//const path = "../../../data/nodes_data_8_10000.txt"
//
//func MakePolicyOutput(state types.State) types.Policy {
//	//fmt.Println("start of make initial policy")
//	found, route, thresholdFailed, accessFailed, paymentsList := policy.SendRequest(&state, 1)
//	policy := types.Policy{
//		Found:                found,
//		Route:                route,
//		ThresholdFailedLists: thresholdFailed,
//		AccessFailed:         accessFailed,
//		PaymentList:          paymentsList,
//	}
//	return policy
//}
//
//func TestUpdateSuccessfulFound(t *testing.T) {
//	state := MakeInitialState(path)
//	policy := MakePolicyOutput(state)
//	policy.Found = true
//	UpdateSuccessfulFound(&state, policy)
//	if state.SuccessfulFound != 1 {
//		t.Errorf("UpdateSuccessfulFound() failed, expected 1, got %d", state.SuccessfulFound)
//	}
//
//}
//
//func TestFailedRequestsThreshold(t *testing.T) {
//	state := MakeInitialState(path)
//	policy := MakePolicyOutput(state)
//	policy.Found = false
//	policy.AccessFailed = false
//	UpdateFailedRequestsThreshold(&state, policy)
//	if state.FailedRequestsThreshold != 1 {
//		t.Errorf("UpdateFailedRequestsThreshold() failed, expected 1, got %d", state.FailedRequestsThreshold)
//	}
//}
//
//func TestFailedRequestsAccess(t *testing.T) {
//	state := MakeInitialState(path)
//	policy := MakePolicyOutput(state)
//	policy.AccessFailed = true
//	UpdateFailedRequestsAccess(&state, policy)
//	if state.FailedRequestsAccess != 1 {
//		t.Errorf("UpdateFailedRequestsAccess() failed, expected 1, got %d", state.FailedRequestsAccess)
//	}
//}
//
//func TestUpdateOriginatorIndex(t *testing.T) {
//	state := MakeInitialState(path)
//	UpdateOriginatorIndex(&state, 100)
//	if int(state.OriginatorIndex) >= Constants.GetOriginators() {
//		if state.OriginatorIndex != 0 {
//			t.Errorf("UpdateOriginatorIndex() failed, expected < %d, got %d", 0, state.OriginatorIndex)
//		}
//	}
//
//}
//
//// TODO: RouteFlushing is having issues with the concurrent updates
////func TestUpdateRouteListAndFlush(t *testing.T) {
////	state := MakeInitialState(path)
////	policy := MakePolicyOutput(state)
////	state.TimeStep = 6250
////	UpdateRouteListAndFlush(&state, policy, int(state.TimeStep))
////	if len(state.RouteLists) != 0 {
////		t.Errorf("UpdateRouteLists() failed, expected 0, got %d", len(state.RouteLists))
////	}
////}
////
////func TestUpdateRouteList(t *testing.T) {
////	state := MakeInitialState(path)
////	policy := MakePolicyOutput(state)
////	state.RouteLists = []Route{}
////	policy.Route = Route{}
////	UpdateRouteListAndFlush(&state, policy, int(state.TimeStep))
////	if state.RouteLists[0][0] != 1 {
////		t.Errorf("UpdateRouteLists() failed, expected 1, got %d", state.RouteLists[0][0])
////	}
////	if state.RouteLists[0][1] != 2 {
////		t.Errorf("UpdateRouteLists() failed, expected 2, got %d", state.RouteLists[0][1])
////	}
////}
