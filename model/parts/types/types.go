package types

type Request struct {
	TimeStep        int
	OriginatorIndex int
	OriginatorId    int
	ChunkId         int
	RespNodes       [4]int
}

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]int

type StateSubset struct {
	OriginatorIndex         int32
	PendingMap              int32
	RerouteMap              int32
	CacheStruct             int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
}

type RouteData struct {
	TimeStep int32 `json:"t"`
	Route    Route `json:"r"`
}

//type StateData struct {
//	TimeStep int         `json:"t"`
//	State    StateSubset `json:"s"`
//}

type State struct {
	Graph                   *Graph
	Originators             []int
	NodesId                 []int
	RouteLists              []Route
	PendingStruct           PendingStruct
	RerouteStruct           RerouteStruct
	CacheStruct             CacheStruct
	OriginatorIndex         int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
}

type RequestResult struct {
	Found                bool
	Route                Route
	ThresholdFailedLists [][]Threshold
	AccessFailed         bool
	PaymentList          []Payment
}

type RouteWithPrice struct {
	RequesterNode int
	ProviderNode  int
	Price         int
}

type PaymentWithPrice struct {
	Payment Payment
	Price   int
}

type Output struct {
	RoutesWithPrice   []RouteWithPrice
	PaymentsWithPrice []PaymentWithPrice
}

type Outputs struct {
	Outputs []Output
}
