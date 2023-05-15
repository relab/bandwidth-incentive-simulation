package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
	"os"
)

type IncomeInfo struct {
	IncomeMap map[int]int
	File      *os.File
	Writer    *bufio.Writer
}

func InitIncomeInfo() *IncomeInfo {
	iinfo := IncomeInfo{}
	iinfo.IncomeMap = make(map[int]int)
	iinfo.File = MakeFile("./results/income.txt")
	iinfo.Writer = bufio.NewWriter(iinfo.File)
	LogExpSting(iinfo.Writer)
	return &iinfo
}

func (ii *IncomeInfo) Close() {
	err := ii.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for output")
	}
	err = ii.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/income.txt")
	}
}

func (o *IncomeInfo) CalculateIncomeFairness() float64 {
	size := config.GetNetworkSize()
	vals := make([]int, size)
	i := 0
	for _, value := range o.IncomeMap {
		vals[i] = value
		i++
	}
	return utils.Gini(vals)
}

func (o *IncomeInfo) CalculateNegativeIncome() float64 {
	totalNegativeIncomeCounter := 0
	for _, value := range o.IncomeMap {
		if value < 0 {
			totalNegativeIncomeCounter += 1
		}
	}
	return float64(totalNegativeIncomeCounter) / float64(config.GetNetworkSize())
}

func (ii *IncomeInfo) Update(output *types.OutputStruct) {
	payments := output.PaymentsWithPrices
	for _, payment := range payments {
		payer := int(payment.Payment.FirstNodeId)
		payee := int(payment.Payment.PayNextId)

		if !(payment.Payment.IsOriginator) {
			ii.IncomeMap[payer] -= payment.Price
		}
		ii.IncomeMap[payee] += payment.Price
	}
}

func (ii *IncomeInfo) Log() {
	negativeIncomeRes := ii.CalculateNegativeIncome()
	incomeFaireness := ii.CalculateIncomeFairness()
	_, err := ii.Writer.WriteString(fmt.Sprintf("Negative income: %f %% \n", negativeIncomeRes*100))
	if err != nil {
		panic(err)
	}
	_, err = ii.Writer.WriteString(fmt.Sprintf("Income fairness: %f \n", incomeFaireness))
	if err != nil {
		panic(err)
	}
}
