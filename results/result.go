package results

import (
	"bufio"
	"encoding/json"
	"math"
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

func ReadOutput(filename string) [][]Transaction {
	file, err := os.Open(filename)
	if err != nil {
		return nil
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

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func AvgRewardPerEachForwardingAction() float64 {
	transactions := ReadOutput("output.txt")
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

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func AvgNumberOfHops() float64 {
	transactions := ReadOutput("output.txt")
	var totalHops int
	for _, transactionList := range transactions {
		totalHops += len(transactionList)
	}
	return float64(totalHops) / float64(len(transactions))
}

type FractionOfRewardsK16 struct {
	Hop1 float64
	Hop2 float64
	Hop3 float64
}

type FractionOfRewardsK8 struct {
	hop1 float64
	hop2 float64
	hop3 float64
	hop4 float64
}

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func AvgFractionOfTotalRewards(filename string) (FractionOfRewardsK16, FractionOfRewardsK8) {
	transactions := ReadOutput(filename)
	fractoionRewardsK16 := FractionOfRewardsK16Calc(transactions)
	fractoionRewardsK8 := FractionOfRewardsK8Calc(transactions)
	return fractoionRewardsK16, fractoionRewardsK8
}

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func FractionOfRewardsK16Calc(transactions [][]Transaction) FractionOfRewardsK16 {
	threeHopRoutes := make([]FractionOfRewardsK16, 0)
	for _, transactionList := range transactions {
		if len(transactionList) == 2 {
			var rewards []int
			threeHopRoute := FractionOfRewardsK16{}
			rewards = append(rewards, transactionList[0].Price-transactionList[1].Price)
			rewards = append(rewards, transactionList[1].Price)
			// for i, transaction := range transactionList {
			// 	if i == len(transactionList)-1 {
			// 		rewards = append(rewards, transaction.Price)
			// 		break
			// 	}
			// 	reward := transaction.Price - transactionList[i+1].Price
			// 	rewards = append(rewards, reward)
			// }
			threeHopRoute.Hop1 = float64(rewards[0]) / float64(sum(rewards))
			threeHopRoute.Hop2 = float64(rewards[1]) / float64(sum(rewards))
			threeHopRoutes = append(threeHopRoutes, threeHopRoute)
		}
		if len(transactionList) == 3 {
			var rewards []int
			threeHopRoute := FractionOfRewardsK16{}
			rewards = append(rewards, transactionList[0].Price-transactionList[1].Price)
			rewards = append(rewards, transactionList[1].Price-transactionList[2].Price)
			rewards = append(rewards, transactionList[2].Price)
			// for i, transaction := range transactionList {
			// 	if i == len(transactionList)-1 {
			// 		rewards = append(rewards, transaction.Price)
			// 		break
			// 	}
			// 	reward := transaction.Price - transactionList[i+1].Price
			// 	rewards = append(rewards, reward)
			// }
			threeHopRoute.Hop1 = float64(rewards[0]) / float64(sum(rewards))
			threeHopRoute.Hop2 = float64(rewards[1]) / float64(sum(rewards))
			threeHopRoute.Hop3 = float64(rewards[2]) / float64(sum(rewards))
			threeHopRoutes = append(threeHopRoutes, threeHopRoute)
		}
	}
	var sumHop1 float64
	var sumHop2 float64
	var sumHop3 float64
	for _, threeHopRoute := range threeHopRoutes {
		sumHop1 += threeHopRoute.Hop1
		sumHop2 += threeHopRoute.Hop2
		sumHop3 += threeHopRoute.Hop3
	}
	return FractionOfRewardsK16{
		Hop1: sumHop1 / float64(len(threeHopRoutes)),
		Hop2: sumHop2 / float64(len(threeHopRoutes)),
		Hop3: sumHop3 / float64(len(threeHopRoutes)),
	}
}

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func FractionOfRewardsK8Calc(transactions [][]Transaction) FractionOfRewardsK8 {
	twoHopRoutes := make([]FractionOfRewardsK8, 0)
	for _, transactionList := range transactions {
		if len(transactionList) == 2 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			rewards = append(rewards, transactionList[0].Price-transactionList[1].Price)
			rewards = append(rewards, transactionList[1].Price)
			// for i, transaction := range transactionList {
			// 	if i == len(transactionList)-1 {
			// 		break
			// 	}
			// 	reward := transaction.Price - transactionList[i+1].Price
			// 	rewards = append(rewards, reward)
			// }
			twoHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			twoHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			twoHopRoutes = append(twoHopRoutes, twoHopRoute)
		}
		if len(transactionList) == 3 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			rewards = append(rewards, transactionList[0].Price-transactionList[1].Price)
			rewards = append(rewards, transactionList[1].Price-transactionList[2].Price)
			rewards = append(rewards, transactionList[2].Price)
			// for i, transaction := range transactionList {
			// 	if i == len(transactionList)-1 {
			// 		break
			// 	}
			// 	reward := transaction.Price - transactionList[i+1].Price
			// 	rewards = append(rewards, reward)
			// }
			twoHopRoute.hop1 = float64(rewards[0]) / float64(sum(rewards))
			twoHopRoute.hop2 = float64(rewards[1]) / float64(sum(rewards))
			twoHopRoute.hop3 = float64(rewards[2]) / float64(sum(rewards))
			twoHopRoutes = append(twoHopRoutes, twoHopRoute)
		}
		if len(transactionList) == 4 {
			var rewards []int
			twoHopRoute := FractionOfRewardsK8{}
			rewards = append(rewards, transactionList[0].Price-transactionList[1].Price)
			rewards = append(rewards, transactionList[1].Price-transactionList[2].Price)
			rewards = append(rewards, transactionList[2].Price-transactionList[3].Price)
			rewards = append(rewards, transactionList[3].Price)
			// for i, transaction := range transactionList {
			// 	if i == len(transactionList)-1 {
			// 		break
			// 	}
			// 	reward := transaction.Price - transactionList[i+1].Price
			// 	rewards = append(rewards, reward)
			// }
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

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func RewardFairnessForForwardingActions(filename string) float64 {
	transactions := ReadOutput(filename)
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
	return gini(rewards)
}

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func RewardFairnessForStoringActions(filename string) float64 {
	transactions := ReadOutput(filename)
	var rewards []int
	for _, transactionList := range transactions {
		if transactionList == nil {
			continue
		}
		reward := transactionList[len(transactionList)-1].Price
		rewards = append(rewards, reward)
	}
	return gini(rewards)
}

// thresholdEnabled: true,    forgivenessEnabled: true,   forgivenessDuringRouting: true,   paymentEnabled: true,   maxPOCheckEnabled: true,
func RewardFarinessForAllActions(filename string) float64 {
	transactions := ReadOutput(filename)
	var rewards []int
	for _, transactionList := range transactions {
		if transactionList == nil {
			continue
		}
		for i, transaction := range transactionList {
			if i == len(transactionList)-1 {
				rewards = append(rewards, transaction.Price)
				break
			}
			reward := transaction.Price - transactionList[i+1].Price
			rewards = append(rewards, reward)
		}
	}
	return gini(rewards)
}

func gini(x []int) float64 {
	total := 0.0
	for i, xi := range x[:len(x)-1] {
		for _, xj := range x[i+1:] {
			total += math.Abs(float64(xi) - float64(xj))
		}
	}
	return total / (math.Pow(float64(len(x)), 2) * mean(x))
}

func mean(x []int) float64 {
	total := 0.0
	for _, xi := range x {
		total += float64(xi)
	}
	return total / float64(len(x))
}
