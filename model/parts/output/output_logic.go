package output

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type FractionOfRewardsK8 struct {
	hop1 float64
	hop2 float64
	hop3 float64
	hop4 float64
}

type MeanRewardPerForward struct {
	AllRewards []int
	SumRewards int
	Writer     *bufio.Writer
}

func (o *MeanRewardPerForward) CalculateMeanRewardPerForward() float64 {
	return float64(o.SumRewards) / float64(len(o.AllRewards))
}

type AvgNumberOfHops struct {
	TotalNumberOfHops int
	NumberOfRoutes    int
	Writer            *bufio.Writer
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
	Writer          *bufio.Writer
}

type Fractions struct {
	Fractions []FractionOfRewardsK16
	Writer    *bufio.Writer
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
	AllStoringRewards    []int
	SumAllStoringRewards int
	Total                float64
	Counter              int
	Writer               *bufio.Writer
}

func (o *RewardFairnessForStoringAction) CalculateRewardFairnessForStoringAction() float64 {
	total := 0.0
	x := o.AllStoringRewards
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
	Writer        *bufio.Writer
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
	Writer                  *bufio.Writer
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

type NegativeIncome struct {
	IncomeMap map[int]int
	Writer    *bufio.Writer
}

func (o *NegativeIncome) CalculateNegativeIncome() float64 {
	totalNegativeIncomeCounter := 0
	for _, value := range o.IncomeMap {
		if value < 0 {
			totalNegativeIncomeCounter += 1
		}
	}
	return float64(totalNegativeIncomeCounter) / float64(10000)
}

func MakeMeanRewardPerForwardFile() (*os.File, string) {
	filePath := "./results/meanRewardPerForward.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	meanRewardPerForwardFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return meanRewardPerForwardFile, filePath
}

func MakeAvgNumberOfHopsFile() (*os.File, string) {
	filePath := "./results/avgNumberOfHops.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	avgNumberOfHopsFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return avgNumberOfHopsFile, filePath
}

func MakeFractionOfRewardsFile() (*os.File, string) {
	filePath := "./results/fractionOfRewards.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	fractionOfRewards, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return fractionOfRewards, filePath
}

func MakeRewardFairnessForForwardingActionFile() (*os.File, string) {
	filepath := "./results/rewardFairnessForForwardingAction.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForForwardingActionFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForForwardingActionFile, filepath
}

func MakeRewardFairnessForStoringActionFile() (*os.File, string) {
	filepath := "./results/rewardFairnessForStoringAction.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForStoringActionFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForStoringActionFile, filepath
}

func MakeRewardFairnessForAllActionsFile() (*os.File, string) {
	filepath := "./results/rewardFairnessForAllActions.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForAllActionsFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForAllActionsFile, filepath
}

func MakeNegativeIncomeFile() (*os.File, string) {
	filepath := "./results/negativeIncome.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	negativeIncomeFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return negativeIncomeFile, filepath
}
