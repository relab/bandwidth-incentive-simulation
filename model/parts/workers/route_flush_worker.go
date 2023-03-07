package workers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"os"
	"sync"
)

func RouteFlushWorker(routeChan chan types.RouteData, globalState *types.State, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	var routeData types.RouteData
	var bytes []byte
	filePath := "./results/routes.json"

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
	defer func(writer *bufio.Writer) {
		err1 := writer.Flush()
		if err1 != nil {
			fmt.Println("Couldn't flush the remaining buffer in the writer for states")
		}
	}(writer)

	for counter := 1; counter < iterations; counter++ {
		routeData = <-routeChan

		bytes, err = json.Marshal(routeData)
		if err != nil {
			panic(err)
		}

		_, err1 := writer.Write(bytes)
		if err1 != nil {
			panic(err1)
		}
		
	}
}

//func routeListConvertAndDumpToFile(routes []types.Route, curTimeStep int, actualFile *os.File) error {

// message := &protoGenerated.RouteData{
// 	TimeStep: int32(curTimeStep),
// 	Routes:   make([]*protoGenerated.Route, len(routes)),
// }
// for i, route := range routes {
// 	var routeList []int32
// 	for _, node := range route {
// 		routeList = append(routeList, int32(node))
// 	}
// 	message.Routes[i] = &protoGenerated.Route{
// 		Waypoints: routeList,
// 	}
// }
// data1, err := proto.Marshal(message)
//data := RouteData{curTimeStep, routes}
//file, _ := json.Marshal(data)
////actualFile, err := os.OpenFile("routes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
////if err != nil {
////	return err
////}
//w := bufio.NewWriter(actualFile)
//_, err := w.Write(file)
//if err != nil {
//	panic(err)
//}
//err = w.Flush()
//if err != nil {
//	panic(err)
//}
//return nil
//}

//func routeListAndFlush(state *types.State, route types.Route, curTimeStep int, actualFile *os.File) []types.Route {
//	state.RouteLists[curTimeStep%10000] = route
//	if (curTimeStep+5000)%10000 == 0 {
//		routeListConvertAndDumpToFile(state.RouteLists, curTimeStep, actualFile)
//		state.RouteLists = make([]types.Route, 10000)
//	}
//	return state.RouteLists
//}
