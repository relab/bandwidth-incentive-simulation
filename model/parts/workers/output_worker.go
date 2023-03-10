package workers

import (
	"fmt"
	"go-incentive-simulation/model/parts/types"
)

func OutputWorker(outputChan chan types.Output) {
	//defer wg.Done()
	var output types.Output
	counter := 0

	for output = range outputChan {
		counter++
		fmt.Println("Nr:", counter, "- Routes with price: ", output.RoutesWithPrice)
		fmt.Println("Nr:", counter, "- Payments with price: ", output.PaymentsWithPrice)
	}
}
