package results

import (
	"bufio"
	"encoding/json"
	"os"
)

type Transaction struct {
	RequesterNode int
	ProviderNode  int
	Price         int
}

type RouteData struct {
	TimeStep int32
	Route    []int32
}

// func ReadRoutes() []RouteData {
// 	file, err := os.Open("routes.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	routes := make([]RouteData, 0)
// 	for scanner.Scan() {
// 		var data RouteData
// 		err := json.Unmarshal(scanner.Bytes(), &data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		routes = append(routes, data)
// 	}
// 	return routes
// }

// func TestTestTEST() {
// 	routes := ReadRoutes()
// 	for _, route := range routes {
// 		fmt.Println(route)
// 	}

// }

func ReadOutput() [][]Transaction {
	file, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	transactions := make([][]Transaction, 0)
	for scanner.Scan() {
		var transactionList []Transaction
		err := json.Unmarshal(scanner.Bytes(), &transactionList)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transactionList)
	}
	return transactions
}

func AvgRewardPerEachForwardingAction() float64 {
	transactions := ReadOutput()
	var rewards []int
	for _, transactionList := range transactions {
		for i, transaction := range transactionList {
			if i == len(transactionList)-1 {
				break
			}
			reward := transaction.Price - transactionList[i+1].Price
			rewards = append(rewards, reward)
		}
	}
	return float64(sum(rewards)) / float64(len(rewards))
}

func sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func AvgNumberOfHops() float64 {
	transactions := ReadOutput()
	var totalHops int
	for _, transactionList := range transactions {
		totalHops += len(transactionList)
	}
	return float64(totalHops) / float64(len(transactions))
}

type FractionOfRewardsK16 struct {
	hop1 float64
	hop2 float64
	hop3 float64
}

type FractionOfRewardsK8 struct {
	hop1 float64
	hop2 float64
	hop3 float64
	hop4 float64
}

func AvgFractionOfTotalRewards() (FractionOfRewardsK16, FractionOfRewardsK8) {
	transactions := ReadOutput()
	fractoionRewardsK16 := FractionOfRewardsK16Calc(transactions)
	fractoionRewardsK8 := FractionOfRewardsK8Calc(transactions)
	return fractoionRewardsK16, fractoionRewardsK8
}

func FractionOfRewardsK16Calc(transactions [][]Transaction) FractionOfRewardsK16 {
	threeHopRoutes := make([]FractionOfRewardsK16, 0)
	for _, transactionList := range transactions {
		if len(transactionList) == 4 {
			var rewards []int
			threeHopRoute := FractionOfRewardsK16{}
			for i, transaction := range transactionList {
				if i == len(transactionList)-1 {
					break
				}
				reward := transaction.Price - transactionList[i+1].Price
				rewards = append(rewards, reward)
			}
			threeHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			threeHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			threeHopRoute.hop3 = float64(rewards[2]) / float64(sum(rewards))
			threeHopRoutes = append(threeHopRoutes, threeHopRoute)
		}
		if len(transactionList) == 3 {
			var rewards []int
			threeHopRoute := FractionOfRewardsK16{}
			for i, transaction := range transactionList {
				if i == len(transactionList)-1 {
					break
				}
				reward := transaction.Price - transactionList[i+1].Price
				rewards = append(rewards, reward)
			}
			threeHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			threeHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			threeHopRoute.hop3 = 0
			threeHopRoutes = append(threeHopRoutes, threeHopRoute)
		}
	}
	var sumHop1 float64
	var sumHop2 float64
	var sumHop3 float64
	for _, threeHopRoute := range threeHopRoutes {
		sumHop1 += threeHopRoute.hop1
		sumHop2 += threeHopRoute.hop2
		sumHop3 += threeHopRoute.hop3
	}
	return FractionOfRewardsK16{
		hop1: sumHop1 / float64(len(threeHopRoutes)),
		hop2: sumHop2 / float64(len(threeHopRoutes)),
		hop3: sumHop3 / float64(len(threeHopRoutes)),
	}
}

func FractionOfRewardsK8Calc(transactions [][]Transaction) FractionOfRewardsK8 {
	twoHopRoutes := make([]FractionOfRewardsK8, 0)
	for _, transactionList := range transactions {
		if len(transactionList) == 3 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			for i, transaction := range transactionList {
				if i == len(transactionList)-1 {
					break
				}
				reward := transaction.Price - transactionList[i+1].Price
				rewards = append(rewards, reward)
			}
			twoHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			twoHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			twoHopRoutes = append(twoHopRoutes, twoHopRoute)
		}
		if len(transactionList) == 4 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			for i, transaction := range transactionList {
				if i == len(transactionList)-1 {
					break
				}
				reward := transaction.Price - transactionList[i+1].Price
				rewards = append(rewards, reward)
			}
			twoHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			twoHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			twoHopRoute.hop3 = float64(rewards[2]) / float64(sum(rewards))
			twoHopRoutes = append(twoHopRoutes, twoHopRoute)
		}
		if len(transactionList) == 5 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			for i, transaction := range transactionList {
				if i == len(transactionList)-1 {
					break
				}
				reward := transaction.Price - transactionList[i+1].Price
				rewards = append(rewards, reward)
			}
			twoHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			twoHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			twoHopRoute.hop3 = float64(rewards[2]) / float64(sum(rewards))
			twoHopRoute.hop4 = float64(rewards[3]) / float64(sum(rewards))
			twoHopRoutes = append(twoHopRoutes, twoHopRoute)
		}
	}
	var sumHop1 float64
	var sumHop2 float64
	var sumHop3 float64
	var sumHop4 float64
	for _, twoHopRoute := range twoHopRoutes {
		sumHop1 += twoHopRoute.hop1
		sumHop2 += twoHopRoute.hop2
		sumHop3 += twoHopRoute.hop3
		sumHop4 += twoHopRoute.hop4
	}
	return FractionOfRewardsK8{
		hop1: sumHop1 / float64(len(twoHopRoutes)),
		hop2: sumHop2 / float64(len(twoHopRoutes)),
		hop3: sumHop3 / float64(len(twoHopRoutes)),
		hop4: sumHop4 / float64(len(twoHopRoutes)),
	}
}
