package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"math"
	"os"
)

// func OutputWorker(outputChan chan types.Output) {
// 	//defer wg.Done()
// 	var output types.Output
// 	counter := 0
// 	filePath := "./results/output.txt"
// 	err := os.Remove(filePath)
// 	if err != nil {
// 		fmt.Println("Could not remove the file", filePath)
// 	}
// 	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
// 	defer func(file *os.File) {
// 		err1 := file.Close()
// 		if err1 != nil {
// 			fmt.Println("Couldn't close the file with filepath: ", filePath)
// 		}
// 	}(file)
// 	for output = range outputChan {
// 		counter++
// 		//fmt.Println("Nr:", counter, "- Routes with price: ", output.RoutesWithPrice)
// 		//fmt.Println("Nr:", counter, "- Payments with price: ", output.PaymentsWithPrice)
// 		jsonData, err := json.Marshal(output.RoutesWithPrice)
// 		if err != nil {
// 			fmt.Println("Couldn't marshal routes with price")
// 		}
// 		file.Write(jsonData)
// 		file.WriteString("\n")
// 	}
// }

type FractionOfRewardsK8 struct {
	hop1 float64
	hop2 float64
	hop3 float64
	hop4 float64
}

type MeanRewardPerForward struct {
	AllRewards []int
	SumRewards int
}

func (o *MeanRewardPerForward) CalculateMeanRewardPerForward() float64 {
	return float64(o.SumRewards) / float64(len(o.AllRewards))
}

type AvgNumberOfHops struct {
	TotalNumberOfHops int
	NumberOfRoutes    int
}

func (o *AvgNumberOfHops) CalculateAverageNumberOfHops() float64 {
	return float64(o.TotalNumberOfHops) / float64(o.NumberOfRoutes)
}

type FractionOfRewardsK16 struct {
	Hop1            float64
	Hop2            float64
	Hop3            float64
	RouteRewards    []int
	SumRouteRewards int
}

type Fractions struct {
	Fractions []FractionOfRewardsK16
}

func (o *Fractions) CalculateFractionOfRewards() (float64, float64, float64) {
	var sumHop1 float64
	var sumHop2 float64
	var sumHop3 float64
	for _, reward := range o.Fractions {
		sumHop1 += reward.Hop1
		sumHop2 += reward.Hop2
		sumHop3 += reward.Hop3
	}
	hop1, hop2, hop3 := sumHop1/float64(len(o.Fractions)), sumHop2/float64(len(o.Fractions)), sumHop3/float64(len(o.Fractions))

	return hop1, hop2, hop3
}

type RewardFairnessForStoringAction struct {
	AllStoringRwards     []int
	SumAllStoringRewards int
	Total                float64
	Counter              int
}

func (o *RewardFairnessForStoringAction) CalculateRewardFairnessForStoringAction() float64 {
	total := 0.0
	x := o.AllStoringRwards
	for i, xi := range x[:len(x)-1] {
		for _, xj := range x[i+1:] {
			total += math.Abs(float64(xi - xj))
		}
	}
	return total / (math.Pow(float64(len(x)), 2) * (float64(o.SumAllStoringRewards) / float64(len(x))))
}

type RewardFairnessForAllActions struct {
	AllRewards    []int
	SumAllRewards int
}

func (o *RewardFairnessForAllActions) CalculateRewardFairnessForAllActions() float64 {
	total := 0.0
	x := o.AllRewards
	for i, xi := range x[:len(x)-1] {
		for _, xj := range x[i+1:] {
			total += math.Abs(float64(xi - xj))
		}
	}
	return total / (math.Pow(float64(len(x)), 2) * (float64(o.SumAllRewards) / float64(len(x))))
}

type RewardFairnessForForwardingActions struct {
	AllForwardingRewards    []int
	SumAllForwardingRewards int
}

func (o *RewardFairnessForForwardingActions) CalculateRewardFairnessForForwardingAction() float64 {
	total := 0.0
	x := o.AllForwardingRewards
	for i, xi := range x[:len(x)-1] {
		for _, xj := range x[i+1:] {
			total += math.Abs(float64(xi - xj))
		}
	}
	return total / (math.Pow(float64(len(x)), 2) * (float64(o.SumAllForwardingRewards) / float64(len(x))))
}

func OutputWorker(outputChan chan types.Output) {
	var output types.Output
	counter := 0
	filePath := "./results/output.txt"
	var meanRewardPerForward MeanRewardPerForward
	var avgNumberOfHops AvgNumberOfHops
	var fractions Fractions
	var rewardFairnessForStoringAction RewardFairnessForStoringAction
	var rewardFairnessForAllActions RewardFairnessForAllActions
	var rewardFairnessForForwardingAction RewardFairnessForForwardingActions
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(file)
	for output = range outputChan {
		counter++
		if constants.GetMeanRewardPerForward() {
			for i := range output.RoutesWithPrice {
				if i == len(output.RoutesWithPrice)-1 {
					break
				}
				reward := output.RoutesWithPrice[i].Price - output.RoutesWithPrice[i+1].Price
				meanRewardPerForward.AllRewards = append(meanRewardPerForward.AllRewards, reward)
				meanRewardPerForward.SumRewards += reward
			}
			if counter%100_000 == 0 {
				mean := meanRewardPerForward.CalculateMeanRewardPerForward()
				file.WriteString(fmt.Sprintf("Mean reward per forward: %f \n", mean))
			}
		}
		if constants.GetAverageNumberOfHops() {
			avgNumberOfHops.TotalNumberOfHops += len(output.RoutesWithPrice)
			avgNumberOfHops.NumberOfRoutes++
			if counter%100_000 == 0 {
				hops := avgNumberOfHops.CalculateAverageNumberOfHops()
				file.WriteString(fmt.Sprintf("Average number of hops: %f \n", hops))
			}
		}
		if constants.GetAverageFractionOfTotalRewardsK16() && constants.GetMaxProximityOrder() == 16 {
			var FractionOfRewardsK16 FractionOfRewardsK16
			if len(output.RoutesWithPrice) == 2 {
				FractionOfRewardsK16.RouteRewards = append(FractionOfRewardsK16.RouteRewards, output.RoutesWithPrice[0].Price-output.RoutesWithPrice[1].Price)
				FractionOfRewardsK16.RouteRewards = append(FractionOfRewardsK16.RouteRewards, output.RoutesWithPrice[1].Price)
				FractionOfRewardsK16.SumRouteRewards += output.RoutesWithPrice[0].Price - output.RoutesWithPrice[1].Price
				FractionOfRewardsK16.SumRouteRewards += output.RoutesWithPrice[1].Price
				FractionOfRewardsK16.Hop1 = float64(FractionOfRewardsK16.RouteRewards[0]) / float64(FractionOfRewardsK16.SumRouteRewards)
				FractionOfRewardsK16.Hop2 = float64(FractionOfRewardsK16.RouteRewards[1]) / float64(FractionOfRewardsK16.SumRouteRewards)
				fractions.Fractions = append(fractions.Fractions, FractionOfRewardsK16)
				FractionOfRewardsK16.RouteRewards = nil
				FractionOfRewardsK16.SumRouteRewards = 0
			}
			if len(output.RoutesWithPrice) == 3 {
				FractionOfRewardsK16.RouteRewards = append(FractionOfRewardsK16.RouteRewards, output.RoutesWithPrice[0].Price-output.RoutesWithPrice[1].Price)
				FractionOfRewardsK16.RouteRewards = append(FractionOfRewardsK16.RouteRewards, output.RoutesWithPrice[1].Price-output.RoutesWithPrice[2].Price)
				FractionOfRewardsK16.RouteRewards = append(FractionOfRewardsK16.RouteRewards, output.RoutesWithPrice[2].Price)
				FractionOfRewardsK16.SumRouteRewards += output.RoutesWithPrice[0].Price - output.RoutesWithPrice[1].Price
				FractionOfRewardsK16.SumRouteRewards += output.RoutesWithPrice[1].Price - output.RoutesWithPrice[2].Price
				FractionOfRewardsK16.SumRouteRewards += output.RoutesWithPrice[2].Price
				FractionOfRewardsK16.Hop1 = float64(FractionOfRewardsK16.RouteRewards[0]) / float64(FractionOfRewardsK16.SumRouteRewards)
				FractionOfRewardsK16.Hop2 = float64(FractionOfRewardsK16.RouteRewards[1]) / float64(FractionOfRewardsK16.SumRouteRewards)
				FractionOfRewardsK16.Hop3 = float64(FractionOfRewardsK16.RouteRewards[2]) / float64(FractionOfRewardsK16.SumRouteRewards)
				fractions.Fractions = append(fractions.Fractions, FractionOfRewardsK16)
				FractionOfRewardsK16.RouteRewards = nil
				FractionOfRewardsK16.SumRouteRewards = 0
			}
			if counter%100_000 == 0 {
				hop1, hop2, hop3 := fractions.CalculateFractionOfRewards()
				file.WriteString(fmt.Sprintf("hop 1: %f, ", hop1))
				file.WriteString(fmt.Sprintf("hop 2: %f, ", hop2))
				file.WriteString(fmt.Sprintf("hop 3: %f \n", hop3))
			}
		}
		if constants.GetRewardFairnessForStoringAction() {
			route := output.RoutesWithPrice
			if route != nil {
				reward := route[len(route)-1].Price
				rewardFairnessForStoringAction.AllStoringRwards = append(rewardFairnessForStoringAction.AllStoringRwards, reward)
				rewardFairnessForStoringAction.SumAllStoringRewards += reward
			}
			if counter == 100_000 {
				fairness := rewardFairnessForStoringAction.CalculateRewardFairnessForStoringAction()
				file.WriteString(fmt.Sprintf("Reward fairness for storing action: %f \n", fairness))
			}
		}
		if constants.GetRewardFairnessForAllActions() {
			route := output.RoutesWithPrice
			if route != nil {
				for i := range route {
					if i == len(route)-1 {
						reward := route[i].Price
						rewardFairnessForAllActions.AllRewards = append(rewardFairnessForAllActions.AllRewards, reward)
						rewardFairnessForAllActions.SumAllRewards += reward
						break
					}
					reward := route[i].Price - route[i+1].Price
					rewardFairnessForAllActions.AllRewards = append(rewardFairnessForAllActions.AllRewards, reward)
					rewardFairnessForAllActions.SumAllRewards += reward
				}
				if counter == 100_000 {
					fairness := rewardFairnessForAllActions.CalculateRewardFairnessForAllActions()
					file.WriteString(fmt.Sprintf("Reward fairness for all actions: %f \n", fairness))
				}
			}
		}
		if constants.GetRewardFairnessForForwardingAction() {
			route := output.RoutesWithPrice
			if route != nil {
				for i := range route {
					if i == len(route)-1 {
						break
					}
					reward := route[i].Price - route[i+1].Price
					rewardFairnessForForwardingAction.AllForwardingRewards = append(rewardFairnessForForwardingAction.AllForwardingRewards, reward)
					rewardFairnessForForwardingAction.SumAllForwardingRewards += reward
				}
				if counter == 100_000 {
					fairness := rewardFairnessForForwardingAction.CalculateRewardFairnessForForwardingAction()
					file.WriteString(fmt.Sprintf("Reward fairness for forwarding action: %f \n", fairness))
				}
			}
		}
	}
}
