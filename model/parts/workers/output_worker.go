package workers

import (
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"os"
)

func OutputWorker(outputChan chan types.Output) {
	//defer wg.Done()
	var output types.Output
	counter := 0
	filePath := "./results/output.txt"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Could not remove the file", filePath)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(file)
	for output = range outputChan {
		counter++
		//fmt.Println("Nr:", counter, "- Routes with price: ", output.RoutesWithPrice)
		//fmt.Println("Nr:", counter, "- Payments with price: ", output.PaymentsWithPrice)
		jsonData, err := json.Marshal(output.RoutesWithPrice)
		if err != nil {
			fmt.Println("Couldn't marshal routes with price")
		}
		file.Write(jsonData)
		file.WriteString("\n")
	}
}
