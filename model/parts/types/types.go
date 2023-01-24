package types

type Request struct {
	Originator *Node
	ChunkId    int
}

type CacheListMap map[*Node][]map[int]int

type RerouteMap map[int][]*Node

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]*Node

type State struct {
	Network                 *Graph
	Originators             []int
	OriginatorsIndex        int
	NodesId                 []int
	RouteLists              []Route
	PendingDict             map[int]int
	RerouteMap              map[int][]int
	CacheDict               map[int]int
	OriginatorIndex         int
	SuccessfulFound         int
	FailedRequestsThreshold int
	FailedRequestsAccess    int
	TimeStep                int
}

type Policy struct {
	Found           bool
	Route           Route
	ThresholdFailed []int
	OriginatorIndex int
	AccessFailed    bool
	PaymentList     []Payment
}
