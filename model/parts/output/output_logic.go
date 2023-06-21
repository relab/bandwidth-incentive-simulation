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
	Update(output *types.OutputStruct)
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
