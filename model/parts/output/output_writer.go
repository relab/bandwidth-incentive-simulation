package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"os"
)

type OutputWriter struct {
	Outputs []types.OutputStruct
	File    *os.File
	Writer  *bufio.Writer
}

func InitOutputWriter() *OutputWriter {
	ow := OutputWriter{}
	ow.Outputs = make([]types.OutputStruct, 0, config.GetEvaluateInterval())
	ow.File = MakeFile("./results/outputs.txt")
	ow.Writer = bufio.NewWriter(ow.File)
	LogExpSting(ow.Writer)
	return &ow
}

func (ow *OutputWriter) Reset() {
	//automatically resets on log
}

func (ow *OutputWriter) Close() {
	err := ow.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for output")
	}
	err = ow.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/outputs.txt")
	}
}

func (ow *OutputWriter) Update(output *types.OutputStruct) {
	ow.Outputs = append(ow.Outputs, *output)
}

func (ow *OutputWriter) Log() {

	for _, o := range ow.Outputs {
		if o.RouteWithPrices != nil {
			ow.Writer.WriteString(fmt.Sprintf("Route: %v \n", o.RouteWithPrices))
		}
		if o.PaymentsWithPrices != nil {
			ow.Writer.WriteString(fmt.Sprintf("Payment Route: %v \n", o.PaymentsWithPrices))

		}
	}

	ow.Outputs = make([]types.OutputStruct, 0, config.GetEvaluateInterval())
}
