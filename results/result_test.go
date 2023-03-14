package results

import (
	"go-incentive-simulation/model/constants"
	"testing"
)

func TestAvgRewardPerEachForwardingAction(t *testing.T) {
	result := AvgRewardPerEachForwardingAction()
	t.Log("Average reward per each forwarding action: ", result, "with k: ", constants.GetBinSize())
}

func TestAvgNumberOfHops(t *testing.T) {
	result := AvgNumberOfHops()
	t.Log("Average number of hops from originator to store: ", result, "with k: ", constants.GetBinSize())
}

func TestAvgFractionOfTotalRewards(t *testing.T) {
	fractoionRewardsK16, fractionRewardsK8 := AvgFractionOfTotalRewards()
	if constants.GetBinSize() == 16 {
		t.Log("Average percent of total rewards for 1 hop: ", fractoionRewardsK16.hop1*100, "with k: ", constants.GetBinSize())
		t.Log("Average percent of total rewards for 2 hop: ", fractoionRewardsK16.hop2*100, "with k: ", constants.GetBinSize())
		t.Log("Average percent of total rewards for 3 hop: ", fractoionRewardsK16.hop3*100, "with k: ", constants.GetBinSize())
	} else if constants.GetBinSize() == 8 {
		t.Log("Average percent of total rewards for 1 hop: ", fractionRewardsK8.hop1*100, "with k: ", constants.GetBinSize())
		t.Log("Average percent of total rewards for 2 hop: ", fractionRewardsK8.hop2*100, "with k: ", constants.GetBinSize())
		t.Log("Average percent of total rewards for 3 hop: ", fractionRewardsK8.hop3*100, "with k: ", constants.GetBinSize())
		t.Log("Average percnt of total rewards for 4 hop: ", fractionRewardsK8.hop4*100, "with k: ", constants.GetBinSize())
	}
}

func TestTest(t *testing.T) {
	transactions := ReadOutput()
	var num1Length int
	var num2Length int
	var num3Length int
	var num4Length int
	var num5Lenght int
	var num6Length int
	for _, transactionList := range transactions {
		if len(transactionList) == 1 {
			num1Length++
		}
		if len(transactionList) == 2 {
			num2Length++
		}
		if len(transactionList) == 3 {
			num3Length++
		}
		if len(transactionList) == 4 {
			num4Length++
		}
		if len(transactionList) == 5 {
			num5Lenght++
		}
		if len(transactionList) == 6 {
			num6Length++
		}
	}
	t.Log("Number of 1 hop routes: ", num1Length)
	t.Log("Number of 2 hop routes: ", num2Length)
	t.Log("Number of 3 hop routes: ", num3Length)
	t.Log("Number of 4 hop routes: ", num4Length)
	t.Log("Number of 5 hop routes: ", num5Lenght)
	t.Log("Number of 6 hop routes: ", num6Length)
	t.Log("Total number of routes: ", num1Length+num2Length+num3Length+num4Length+num5Lenght+num6Length)
}
