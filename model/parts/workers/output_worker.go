package workers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func OutputWorker(outputChan chan types.Output, wg *sync.WaitGroup) {
	defer wg.Done()
	var output types.Output
	//counter := 0
	filePath := "./results/output.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("No need to remove file with path: ", filePath)
	}

	actualFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(actualFile)

	writer := bufio.NewWriter(actualFile) // default writer size is 4096 bytes
	//writer = bufio.NewWriterSize(writer, 1048576) // 1MiB
	defer func(writer *bufio.Writer) {
		err1 := writer.Flush()
		if err1 != nil {
			fmt.Println("Couldn't flush the remaining buffer in the writer for output")
		}
	}(writer)

	for output = range outputChan {
		//counter++
		//fmt.Println("Nr:", counter, "- Routes with price: ", output.NodePairWithPrice)
		//fmt.Println("Nr:", counter, "- Payments with price: ", output.PaymentsWithPrices)
		jsonData, err := json.Marshal(output.RouteWithPrices)
		if err != nil {
			fmt.Println("Couldn't marshal routes with price")
		}
		_, err = writer.Write(jsonData)
		if err != nil {
			return
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			return
		}
	}
}
