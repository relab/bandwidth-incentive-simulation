package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/utils"
	"os"
)

type AttackInfo struct {
	IncomeMap  map[int]int
	Requesters map[int]int
	File       *os.File
	Writer     *bufio.Writer
}

func InitAttackInfo() *AttackInfo {
	ai := AttackInfo{}
	ai.IncomeMap = make(map[int]int)
	ai.Requesters = make(map[int]int) //This map is currently used to find out who is an originator. This should instead be looked up somewhere else.

	ai.File = MakeFile("./results/attack.txt")
	ai.Writer = bufio.NewWriter(ai.File)
	LogExpSting(ai.Writer)
	return &ai
}

func (ai *AttackInfo) Reset() {
	ai.IncomeMap = make(map[int]int)
	ai.Requesters = make(map[int]int) //This map is currently used to find out who is an originator. This should instead be looked up somewhere else.
}

func (ai *AttackInfo) Close() {
	err := ai.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for output")
	}
	err = ai.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/attack.txt")
	}
}

func isAttacker(nodeid int) bool {
	return nodeid%10 == 1
}

func (o *AttackInfo) CalculateAvgIncome() (attackAvg, normalAvg float64) {
	size := config.GetNetworkSize()
	normalincome := make([]int, 0, size)
	attackerincome := make([]int, 0, size/10)
	for id, value := range o.IncomeMap {
		if isAttacker(id) {
			attackerincome = append(attackerincome, value)
		} else {
			normalincome = append(normalincome, value)
		}
	}
	return utils.Mean(attackerincome), utils.Mean(normalincome)
}

func (ai *AttackInfo) Update(output *Route) {
	if output.failed() {
		return
	}
	payments := output.PaymentsWithPrices
	for hop, payment := range payments {
		payer := int(payment.Payment.FirstNodeId)
		payee := int(payment.Payment.PayNextId)

		if !(payment.Payment.IsOriginator) {
			ai.IncomeMap[payer] -= payment.Price
		} else {
			ai.Requesters[payee]++
			if hop != 0 {
				panic("First payment in list is not from originator.")
			}
		}
		ai.IncomeMap[payee] += payment.Price
	}
}

func (ai *AttackInfo) Log() {

	attackAvg, normalAvg := ai.CalculateAvgIncome()
	_, err := ai.Writer.WriteString(fmt.Sprintf("Average income normal nodes: %f \n", normalAvg))
	if err != nil {
		panic(err)
	}

	_, err = ai.Writer.WriteString(fmt.Sprintf("Average income attacker nodes: %f \n", attackAvg))
	if err != nil {
		panic(err)
	}
}
