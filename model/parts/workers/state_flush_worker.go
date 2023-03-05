package workers

import (
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func StateFlushWorker(stateChan chan types.StateSubset, globalState *types.State, stateList []types.StateSubset, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	counter := 1
	var state types.StateSubset
	os.Remove("states.json")
	actualFile, err := os.OpenFile("states.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer actualFile.Close()
	for counter < iterations {
		state = <-stateChan
		stateListAndFlush(state, stateList, actualFile)
		counter++
		//fmt.Println(counter)
	}
}

func stateListConvertAndDumpToFile(stateList []types.StateSubset, curTimeStep int, actualFile *os.File) error {
	type StateData struct {
		TimeStep int                 `json:"timestep"`
		States   []types.StateSubset `json:"states"`
	}
	//subList := make([]types.StateSubset, len(stateList))
	//for i, state := range stateList {
	//	subList[i] = types.StateSubset{
	//		OriginatorIndex:         state.OriginatorIndex,
	//		PendingMap:              state.PendingStruct.PendingMap,
	//		RerouteMap:              state.RerouteStruct.RerouteMap,
	//		SuccessfulFound:         state.SuccessfulFound,
	//		FailedRequestsThreshold: state.FailedRequestsThreshold,
	//		FailedRequestsAccess:    state.FailedRequestsAccess,
	//		TimeStep:                state.TimeStep,
	//	}
	//}
	data := StateData{curTimeStep, stateList}
	//file, _ := json.Marshal(data)
	file, _ := json.MarshalIndent(data, "", "  ")
	_, err := actualFile.Write(file)
	if err != nil {
		return err
	}
	return nil
}

func stateListAndFlush(state types.StateSubset, stateList []types.StateSubset, actualFile *os.File) []types.StateSubset {
	stateList[state.TimeStep%10000] = state
	if state.TimeStep%10000 == 0 {
		err := stateListConvertAndDumpToFile(stateList, int(state.TimeStep), actualFile)
		if err != nil {
			return nil
		}
		stateList = make([]types.StateSubset, 10000)
	}
	fmt.Println(state.TimeStep)
	return stateList
}
