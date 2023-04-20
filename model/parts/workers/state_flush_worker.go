package workers

import (
	"encoding/csv"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"os"
	"strconv"
	"sync"
)

func StateFlushWorker(stateChan chan types.StateSubset, wg *sync.WaitGroup) {
	defer wg.Done()
	//var message *protoGenerated.StateSubsets
	//var bytes []byte
	//filePath := "./results/states.bin"
	filePath := "./results/states.csv"

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("No need to remove file with path: ", filePath)
	}

	actualFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func(actualFile *os.File) {
		err1 := actualFile.Close()
		if err1 != nil {
			fmt.Println("Couldn't close the file with filepath: ", filePath)
		}
	}(actualFile)

	//writer := bufio.NewWriter(actualFile) // default writer size is 4096 bytes
	////writer = bufio.NewWriterSize(writer, 1048576) // 1MiB
	//defer func(writer *bufio.Writer) {
	//	err1 := writer.Flush()
	//	if err1 != nil {
	//		fmt.Println("Couldn't flush the remaining buffer in the writer for states")
	//	}
	//}(writer)

	writer := csv.NewWriter(actualFile) // default writer size is 4096 bytes
	defer writer.Flush()

	err = writer.Write([]string{
		"WaitingCounter",
		"RetryCounter",
		"CacheCounter",
		"SuccessfulFound",
		"FailedRequestsThreshold",
		"FailedRequestsAccess",
		"TimeStep",
		"Epoch"})

	if err != nil {
		panic(err)
	}

	var stateSubset types.StateSubset
	var flushInterval = config.GetRequestsPerSecond()
	//var flushInterval = 1000
	counter := 0

	for stateSubset = range stateChan {
		counter++

		if counter%flushInterval == 0 {

			// Write the CSV row
			err = writer.Write([]string{
				strconv.Itoa(stateSubset.WaitingCounter),
				strconv.Itoa(stateSubset.RetryCounter),
				strconv.Itoa(stateSubset.CacheHits),
				strconv.Itoa(stateSubset.SuccessfulFound),
				strconv.Itoa(stateSubset.FailedRequestsThreshold),
				strconv.Itoa(stateSubset.FailedRequestsAccess),
				strconv.Itoa(stateSubset.TimeStep),
				strconv.Itoa(stateSubset.Epoch)})

			if err != nil {
				panic(err)
			}

			writer.Flush()
		}

		//outStateSubset.WaitingCounter += inStateSubset.WaitingCounter
		//outStateSubset.RetryCounter += inStateSubset.RetryCounter
		//outStateSubset.CacheHits += inStateSubset.CacheHits
		//outStateSubset.SuccessfulFound += inStateSubset.SuccessfulFound
		//outStateSubset.FailedRequestsThreshold += inStateSubset.FailedRequestsThreshold
		//outStateSubset.FailedRequestsAccess += inStateSubset.FailedRequestsAccess

		//message = &protoGenerated.StateSubsets{
		//	Subset: make([]*protoGenerated.StateSubset, 0),
		//}
		//
		//// TODO: Update the proto generation based on new stateSubset
		////message.Subset = append(message.Subset, &protoGenerated.StateSubset{
		////	OriginatorIndex:         stateSubset.OriginatorIndex,
		////	PendingMap:              stateSubset.PendingMap,
		////	RerouteMap:              stateSubset.RerouteMap,
		////	SuccessfulFound:         stateSubset.SuccessfulFound,
		////	FailedRequestsThreshold: stateSubset.FailedRequestsThreshold,
		////	FailedRequestsAccess:    stateSubset.FailedRequestsAccess,
		////	TimeStep:                stateSubset.TimeStep,
		////})
		//fmt.Println(stateSubset)
		//
		//bytes, err = proto.Marshal(message)
		//if err != nil {
		//	panic(err)
		//}
		//
		//_, err = writer.Write(bytes)
		//if err != nil {
		//	panic(err)
		//}
	}
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
