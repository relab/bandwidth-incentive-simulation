package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"os"
)

type LogResetUpdater interface {
	Log()
	Reset()
	Update(output *OutputStruct)
}

type OutputStruct struct {
	RouteWithPrices    []types.NodePairWithPrice
	PaymentsWithPrices []types.PaymentWithPrice
	Found              bool
	AccessFailed       bool
	ThresholdFailed    bool
	FoundByCaching     bool
}

func (o *OutputStruct) failed() bool {
	return o.ThresholdFailed || o.AccessFailed
}

func MakeFile(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func LogExpSting(writer *bufio.Writer) {
	_, err := writer.WriteString(fmt.Sprintf("\n %s \n\n", config.GetExperimentString()))
	if err != nil {
		panic(err)
	}
}
