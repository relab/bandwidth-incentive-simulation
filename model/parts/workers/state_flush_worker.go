package workers

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/protoGenerated"
	"google.golang.org/protobuf/proto"
	"os"
	"sync"
)

func StateFlushWorker(stateChan chan types.StateSubset, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	var message *protoGenerated.StateSubsets
	var stateSubset types.StateSubset
	var bytes []byte
	filePath := "./results/states.bin"

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

	writer := bufio.NewWriter(actualFile)
	//writer = bufio.NewWriterSize(writer, 1048576) // 1MiB
	defer func(writer *bufio.Writer) {
		err1 := writer.Flush()
		if err1 != nil {
			fmt.Println("Couldn't flush the remaining buffer in the writer for states")
		}
	}(writer)

	for counter := 0; counter < iterations; counter++ {
		stateSubset = <-stateChan

		message = &protoGenerated.StateSubsets{
			Subset: make([]*protoGenerated.StateSubset, 0),
		}

		message.Subset = append(message.Subset, &protoGenerated.StateSubset{
			OriginatorIndex:         stateSubset.OriginatorIndex,
			PendingMap:              stateSubset.PendingMap,
			RerouteMap:              stateSubset.RerouteMap,
			SuccessfulFound:         stateSubset.SuccessfulFound,
			FailedRequestsThreshold: stateSubset.FailedRequestsThreshold,
			FailedRequestsAccess:    stateSubset.FailedRequestsAccess,
			TimeStep:                stateSubset.TimeStep,
		})

		bytes, err = proto.Marshal(message)
		if err != nil {
			panic(err)
		}

		_, err = writer.Write(bytes)
		if err != nil {
			panic(err)
		}
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
