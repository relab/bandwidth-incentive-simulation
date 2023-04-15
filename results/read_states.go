package results

import (
	"fmt"
	"go-incentive-simulation/protoGenerated"
	"google.golang.org/protobuf/proto"
	"os"
)

func readStatesFile() {
	buf, err := os.ReadFile("states.bin")
	if err != nil {
		panic(err)
	}
	stateSubsets := &protoGenerated.StateSubsets{}
	err = proto.Unmarshal(buf, stateSubsets)
	if err != nil {
		panic(err)
	}
	// Access the subset field
	count := 0
	for _, subset := range stateSubsets.Subset {
		count++
		if count > 10 {
			break
		}
		fmt.Printf("OriginatorIndex: %d\n", subset.OriginatorIndex)
		fmt.Printf("PendingMap: %d\n", subset.PendingMap)
		fmt.Printf("RerouteMap: %d\n", subset.RerouteMap)
		fmt.Printf("CacheStruct: %d\n", subset.CacheStruct)
		fmt.Printf("SuccessfulFound: %d\n", subset.SuccessfulFound)
		fmt.Printf("FailedRequestsThreshold: %d\n", subset.FailedRequestsThreshold)
		fmt.Printf("FailedRequestsAccess: %d\n", subset.FailedRequestsAccess)
		fmt.Printf("TimeStep: %d\n", subset.TimeStep)
	}
	// read the binary protobuf message from the file
	buf, err = os.ReadFile("routes.bin")
	if err != nil {
		panic(err)
	}

	// unmarshal the binary protobuf message into a RouteData struct
	routeData := &protoGenerated.RouteData{}
	err = proto.Unmarshal(buf, routeData)
	if err != nil {
		panic(err)
	}

	//// print the RouteData struct
	//fmt.Printf("TimeStep: %d\n", routeData.GetTimeStep())
	//count = 0
	//routedata := routeData.GetRoute()
	//fmt.Println("length", len(routedata))
	//for _, route := range routeData.GetRoutes() {
	//	if count == 10 {
	//		break
	//	}
	//	fmt.Printf("RequestResult: %v\n", route.GetWaypoints())
	//	fmt.Printf("Length: %d\n", route.GetLength())
	//	count++
	//}
}
