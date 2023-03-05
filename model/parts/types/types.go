package types

type Request struct {
	OriginatorId int
	ChunkId      int
	RespNodes    [4]int
}

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]int

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

type Policy struct {
	Found                bool
	Route                Route
	ThresholdFailedLists [][]Threshold
	AccessFailed         bool
	PaymentList          []Payment
}
