package types

type Request struct {
	TimeStep        int
	Epoch           int
	OriginatorIndex int
	OriginatorId    NodeId
	ChunkId         ChunkId
	RespNodes       [4]NodeId
}

type Route []NodeId

type RequestResult struct {
	Route           Route
	ChunkId         ChunkId
	Found           bool
	AccessFailed    bool
	ThresholdFailed bool
	FoundByCaching  bool
	PaymentList     []Payment
}

//type RequestResult struct {
//	Route       RequestResult
//	PaymentList []Payment
//}

type Payment struct {
	FirstNodeId  NodeId
	PayNextId    NodeId
	ChunkId      ChunkId
	IsOriginator bool
}

type Threshold [2]NodeId

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
	TimeStep        int32 `json:"t"`
	Route           Route `json:"r"`
	ThresholdFailed bool
	AccessFailed    bool
}

//type StateData struct {
//	TimeStep int         `json:"t"`
//	State    StateSubset `json:"s"`
//}

type State struct {
	Graph                   *Graph
	Originators             []NodeId
	NodesId                 []NodeId
	RouteLists              []RequestResult
	PendingStruct           PendingStruct
	RerouteStruct           RerouteStruct
	CacheStruct             CacheStruct
	OriginatorIndex         int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
	Epoch                   int
}

type RouteWithPrice struct {
	RequesterNode NodeId
	ProviderNode  NodeId
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
