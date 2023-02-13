package update

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/policy"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/state"
	"testing"
)

const path = "../../../data/nodes_data_8_10000.txt"

func MakePolicyOutput(state State) Policy {
	//fmt.Println("start of make initial policy")
	found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(&state, 1)
	policy := Policy{
		Found:                found,
		Route:                route,
		ThresholdFailedLists: thresholdFailed,
		AccessFailed:         accessFailed,
		PaymentList:          paymentsList,
	}
	return policy
}

func TestUpdateSuccessfulFound(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	policy.Found = true
	newState := UpdateSuccessfulFound(state, policy)
	if newState.SuccessfulFound != 1 {
		t.Errorf("UpdateSuccessfulFound() failed, expected 1, got %d", newState.SuccessfulFound)
	}

}

func TestFailedRequestsThreshold(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	policy.Found = false
	policy.AccessFailed = false
	newState := UpdateFailedRequestsThreshold(state, policy)
	if newState.FailedRequestsThreshold != 1 {
		t.Errorf("UpdateFailedRequestsThreshold() failed, expected 1, got %d", newState.FailedRequestsThreshold)
	}
}

func TestFailedRequestsAccess(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	policy.AccessFailed = true
	newState := UpdateFailedRequestsAccess(state, policy)
	if newState.FailedRequestsAccess != 1 {
		t.Errorf("UpdateFailedRequestsAccess() failed, expected 1, got %d", newState.FailedRequestsAccess)
	}
}

func TestUpdateOriginatorIndex(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	newState := UpdateOriginatorIndex(state, policy)
	if newState.OriginatorIndex >= Constants.GetOriginators() {
		if newState.OriginatorIndex != 0 {
			t.Errorf("UpdateOriginatorIndex() failed, expected < %d, got %d", 0, newState.OriginatorIndex)
		}
	}

}

func TestUpdateRouteListAndFlush(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	state.TimeStep = 6249
	newState := UpdateRouteListAndFlush(state, policy)
	if len(newState.RouteLists) != 0 {
		t.Errorf("UpdateRouteLists() failed, expected 0, got %d", len(newState.RouteLists))
	}
}

func TestUpdateRouteList(t *testing.T) {
	state := MakeInitialState(path)
	policy := MakePolicyOutput(state)
	state.RouteLists = []Route{}
	policy.Route = []int{1, 2, 3}
	newState := UpdateRouteListAndFlush(state, policy)
	if newState.RouteLists[0][0] != 1 {
		t.Errorf("UpdateRouteLists() failed, expected 1, got %d", newState.RouteLists[0][0])
	}
	if newState.RouteLists[0][1] != 2 {
		t.Errorf("UpdateRouteLists() failed, expected 2, got %d", newState.RouteLists[0][1])
	}
}
