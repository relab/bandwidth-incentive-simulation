package types

type Request struct {
	Originator *Node
	ChunkId    int
}

type PendingMap map[int]int

type RerouteMap map[int][]int

type CacheListMap map[*Node][]map[int]int

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
	NodesId                 []int
	RouteLists              []Route
	PendingMap              PendingMap
	RerouteMap              RerouteMap
	CacheListMap            CacheListMap
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
