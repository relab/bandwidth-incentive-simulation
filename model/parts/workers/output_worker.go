package workers

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"math"
	"os"
	"sync"
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
	IncomeDict map[int]int
	Writer     *bufio.Writer
}

func (o *NegativeIncome) CalculateNegativeIncome() float64 {
	totalNegativeIncomeCounter := 0
	for _, value := range o.IncomeDict {
		if value < 0 {
			totalNegativeIncomeCounter += 1
		}
	}
	return float64(totalNegativeIncomeCounter) / float64(10000)
}

func OutputWorker(outputChan chan types.Output, wg *sync.WaitGroup) {
	defer wg.Done()
	var output types.Output
	counter := 1
	var meanRewardPerForward MeanRewardPerForward
	var avgNumberOfHops AvgNumberOfHops
	var fractions Fractions
	var rewardFairnessForStoringAction RewardFairnessForStoringAction
	var rewardFairnessForAllActions RewardFairnessForAllActions
	var rewardFairnessForForwardingAction RewardFairnessForForwardingActions
	var negativeIncome NegativeIncome
	negativeIncome.IncomeDict = make(map[int]int)
	filePath := "./results/output.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(file)

	writer := bufio.NewWriter(file) // default writer size is 4096 bytes
	//writer = bufio.NewWriterSize(writer, 1048576) // 1MiB
	defer func(writer *bufio.Writer) {
		err1 := writer.Flush()
		if err1 != nil {
			fmt.Println("Couldn't flush the remaining buffer in the writer for output")
		}
	}(writer)
	if constants.GetMeanRewardPerForward() {
		file := MakeMeanRewardPerForwardFile()
		defer func(file *os.File) {
			err1 := file.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file)
		meanRewardPerForward.Writer = bufio.NewWriter(file)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(meanRewardPerForward.Writer)
	}
	if constants.GetAverageNumberOfHops() {
		file2 := MakeAvgNumberOfHopsFile()
		defer func(file2 *os.File) {
			err1 := file2.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file2)
		avgNumberOfHops.Writer = bufio.NewWriter(file2)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(avgNumberOfHops.Writer)
	}
	if constants.GetAverageFractionOfTotalRewardsK16() {
		file3 := MakeFractionOfRewardsFile()
		defer func(file3 *os.File) {
			err1 := file3.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file3)
		fractions.Writer = bufio.NewWriter(file3)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(fractions.Writer)
	}
	if constants.GetRewardFairnessForStoringAction() {
		file4 := MakeRewardFairnessForStoringActionFile()
		defer func(file4 *os.File) {
			err1 := file4.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file4)
		rewardFairnessForStoringAction.Writer = bufio.NewWriter(file4)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(rewardFairnessForStoringAction.Writer)
	}
	if constants.GetRewardFairnessForAllActions() {
		file5 := MakeRewardFairnessForAllActionsFile()
		defer func(file5 *os.File) {
			err1 := file5.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file5)
		rewardFairnessForAllActions.Writer = bufio.NewWriter(file5)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(rewardFairnessForAllActions.Writer)
	}
	if constants.GetRewardFairnessForForwardingAction() {
		file6 := MakeRewardFairnessForForwardingActionFile()
		defer func(file6 *os.File) {
			err1 := file6.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file6)
		rewardFairnessForForwardingAction.Writer = bufio.NewWriter(file6)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(rewardFairnessForForwardingAction.Writer)
	}
	if constants.GetNegativeIncome() {
		file7 := MakeNegativeIncomeFile()
		defer func(file7 *os.File) {
			err1 := file7.Close()
			if err1 != nil {
				fmt.Println("Couldn't close the file with filepath: ", filePath)
			}
		}(file7)
		negativeIncome.Writer = bufio.NewWriter(file7)
		defer func(writer *bufio.Writer) {
			err1 := writer.Flush()
			if err1 != nil {
				fmt.Println("Couldn't flush the remaining buffer in the writer for output")
			}
		}(negativeIncome.Writer)
	}
	for output = range outputChan {
		counter++
		if counter%100_000 == 0 {
			//fmt.Println("counter is now: ", counter)
			//fmt.Println("time since start: ", time.Since(start))
		}
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
				_, err := meanRewardPerForward.Writer.WriteString(fmt.Sprintf("Mean reward per forward: %f \n", mean))
				if err != nil {
					panic(err)
				}
				//fmt.Println("time since start: ", time.Since(start))
			}
		}
		if constants.GetAverageNumberOfHops() {
			avgNumberOfHops.TotalNumberOfHops += len(output.RoutesWithPrice)
			avgNumberOfHops.NumberOfRoutes++
			if counter%100_000 == 0 {
				hops := avgNumberOfHops.CalculateAverageNumberOfHops()
				avgNumberOfHops.Writer.WriteString(fmt.Sprintf("Average number of hops: %f \n", hops))
				//fmt.Println("time since start: ", time.Since(start))
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
				//writer.WriteString(fmt.Sprintf("hop 1: %f, ", hop1))
				//writer.WriteString(fmt.Sprintf("hop 2: %f, ", hop2))
				//writer.WriteString(fmt.Sprintf("hop 3: %f \n", hop3))
				fractions.Writer.WriteString(fmt.Sprintf("hop 1: %f, hop 2: %f, hop 3: %f \n", hop1, hop2, hop3))
				//fmt.Println("time since start: ", time.Since(start))
			}
		}
		if constants.GetRewardFairnessForStoringAction() {
			route := output.RoutesWithPrice
			if route != nil {
				reward := route[len(route)-1].Price
				rewardFairnessForStoringAction.AllStoringRewards = append(rewardFairnessForStoringAction.AllStoringRewards, reward)
				rewardFairnessForStoringAction.SumAllStoringRewards += reward
			}
			if counter == 100_000 {
				fairness := rewardFairnessForStoringAction.CalculateRewardFairnessForStoringAction()
				writer.WriteString(fmt.Sprintf("Reward fairness for storing action: %f \n", fairness))
				//fmt.Println("time since start: ", time.Since(start))
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
					rewardFairnessForAllActions.Writer.WriteString(fmt.Sprintf("Reward fairness for all actions: %f \n", fairness))
					//fmt.Println("time since start: ", time.Since(start))
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
					rewardFairnessForAllActions.Writer.WriteString(fmt.Sprintf("Reward fairness for forwarding action: %f \n", fairness))
					//fmt.Println("time since start: ", time.Since(start))
				}
			}
		}
		// payment enabled, forgivness enabled, threshold enabled, k = 8
		if constants.GetNegativeIncome() && constants.GetPaymentEnabled() && constants.IsForgivenessDuringRouting() { //Payment enabled can use outout.Payment
			payments := output.PaymentsWithPrice
			if payments != nil {
				for i, payment := range payments {
					if i == len(payments)-1 {
						break
					}
					payer := payment.Payment.FirstNodeId
					payee := payment.Payment.PayNextId
					value := payment.Price
					valPayer, ok := negativeIncome.IncomeDict[payer]
					if !ok {
						negativeIncome.IncomeDict[payer] = 0
					}
					valPayee, ok := negativeIncome.IncomeDict[payee]
					if !ok {
						negativeIncome.IncomeDict[payee] = 0
					}
					negativeIncome.IncomeDict[payer] = valPayer - value
					negativeIncome.IncomeDict[payee] = valPayee + value
				}
			}
			// if counter%500_000==0 or counter==100_000 {
			if counter%100_000 == 0 {
				negativeIncome := negativeIncome.CalculateNegativeIncome()
				writer.WriteString(fmt.Sprintf("Negative income: %f %% \n", negativeIncome*100))
			}
		}

	}
}

func MakeAvgNumberOfHopsFile() *os.File {
	filePath := "./results/avgNumberOfHops.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	avgNumberOfHopsFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return avgNumberOfHopsFile
}

func MakeMeanRewardPerForwardFile() *os.File {
	filePath := "./results/meanRewardPerForward.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	meanRewardPerForwardFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return meanRewardPerForwardFile
}

func MakeFractionOfRewardsFile() *os.File {
	filePath := "./results/fractionOfRewards.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	fractionOfRewards, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return fractionOfRewards
}

func MakeRewardFairnessForStoringActionFile() *os.File {
	filepath := "./results/rewardFairnessForStoringAction.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForStoringActionFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForStoringActionFile
}

func MakeRewardFairnessForAllActionsFile() *os.File {
	filepath := "./results/rewardFairnessForAllActions.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForAllActionsFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForAllActionsFile
}

func MakeRewardFairnessForForwardingActionFile() *os.File {
	filepath := "./results/rewardFairnessForForwardingAction.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	rewardFairnessForForwardingActionFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return rewardFairnessForForwardingActionFile
}

func MakeNegativeIncomeFile() *os.File {
	filepath := "./results/negativeIncome.txt"
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Could not remove the file", filepath)
	}
	negativeIncomeFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return negativeIncomeFile

}
