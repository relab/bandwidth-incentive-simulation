package workers

import "go-incentive-simulation/model/constants"

func waitForRoutingWorkers(pauseChan chan bool, continueChan chan bool) {
	for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
		pauseChan <- true
	}
	for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
		<-continueChan
	}
	return
}
