package workers

import (
	"bufio"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func StateFlushWorker(stateChan chan []byte, globalState *types.State, stateList []types.StateSubset, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	counter := 1
	//var stateData types.StateData
	var encodedData []byte
	os.Remove("states.json")
	actualFile, err := os.OpenFile("states.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(actualFile)
	writer = bufio.NewWriterSize(writer, 1000000000)
	//start := time.Now()
	defer actualFile.Close()
	for counter < iterations {
		encodedData = <-stateChan

		//fmt.Println(len(stateChan))
		//stateListAndFlush(state, stateList, actualFile)
		//stateListAndFlush(state, counter, writer)
		//encodedData, _ := json.Marshal(stateData)
		//fmt.Println("2: ", time.Since(start))
		writer.Write(encodedData)
		//fmt.Println("3: ", time.Since(start))

		if counter%10000 == 0 {
			writer.Flush()
		}
		counter++
		//fmt.Println(counter)

	}
	writer.Flush()
}

//func stateListConvertAndDumpToFile(state types.StateSubset, curTimeStep int, writer *bufio.Writer) error {
//	//subList := make([]types.StateSubset, len(stateList))
//	//for i, state := range stateList {
//	//	subList[i] = types.StateSubset{
//	//		OriginatorIndex:         state.OriginatorIndex,
//	//		PendingMap:              state.PendingStruct.PendingMap,
//	//		RerouteMap:              state.RerouteStruct.RerouteMap,
//	//		SuccessfulFound:         state.SuccessfulFound,
//	//		FailedRequestsThreshold: state.FailedRequestsThreshold,
//	//		FailedRequestsAccess:    state.FailedRequestsAccess,
//	//		TimeStep:                state.TimeStep,
//	//	}
//	//}
//	data := StateData{curTimeStep, state}
//	file, _ := json.Marshal(data)
//	writer.Write(file)
//
//	//err = writer.Flush()
//	//if err != nil {
//	//	panic(err)
//	//}
//	return nil
//}
//
//func stateListAndFlush(state types.StateSubset, counter int, writer *bufio.Writer) {
//
//	stateListConvertAndDumpToFile(state, int(state.TimeStep), writer)
//
//	if counter%100000 == 0 {
//		go writer.Flush()
//	}
//	//fmt.Println(state.TimeStep)
//}

//func stateListAndFlush(state types.StateSubset, stateList []types.StateSubset, actualFile *os.File) []types.StateSubset {
//	stateList[state.TimeStep%100000] = state
//	if state.TimeStep%100000 == 0 {
//		stateListConvertAndDumpToFile(stateList, int(state.TimeStep), actualFile)
//		stateList = make([]types.StateSubset, 100000)
//	}
//	//fmt.Println(state.TimeStep)
//	return stateList
//}
