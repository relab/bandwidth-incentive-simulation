package workers

import (
	"bufio"
	"encoding/json"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func RouteFlushWorker(routeChan chan types.Route, globalState *types.State, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	counter := 1
	var route types.Route
	os.Remove("routes.json")
	actualFile, err := os.OpenFile("routes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer actualFile.Close()
	for counter < iterations {
		route = <-routeChan
		routeListAndFlush(globalState, route, counter, actualFile)
		counter++
		//fmt.Println(counter)
	}
}

func routeListConvertAndDumpToFile(routes []types.Route, curTimeStep int, actualFile *os.File) error {
	type RouteData struct {
		TimeStep int           `json:"timestep"`
		Routes   []types.Route `json:"routes"`
	}
	data := RouteData{curTimeStep, routes}
	file, _ := json.Marshal(data)
	actualFile, err := os.OpenFile("routes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(actualFile)
	_, err = w.Write(file)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	return nil
}

func routeListAndFlush(state *types.State, route types.Route, curTimeStep int, actualFile *os.File) []types.Route {
	state.RouteLists[curTimeStep%10000] = route
	if (curTimeStep+5000)%10000 == 0 {
		routeListConvertAndDumpToFile(state.RouteLists, curTimeStep, actualFile)
		state.RouteLists = make([]types.Route, 10000)
	}
	return state.RouteLists
}
