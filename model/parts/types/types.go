package types

type Request struct {
	TimeStep        int
	Epoch           int
	OriginatorIndex int32
	OriginatorId    NodeId
	ChunkId         ChunkId
	RespNodes       [4]NodeId
}

type RequestResult struct {
	Route           []NodeId
	PaymentList     []Payment
	ChunkId         ChunkId
	Found           bool
	AccessFailed    bool
	ThresholdFailed bool
	FoundByCaching  bool
}

type Payment struct {
	FirstNodeId  NodeId
	PayNextId    NodeId
	ChunkId      ChunkId
	IsOriginator bool
}

func (p Payment) IsNil() bool {
	if p.PayNextId == 0 && p.FirstNodeId == 0 && p.ChunkId == 0 {
		return true
	} else {
		return false
	}
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
	Epoch                   int32
}

type RouteData struct {
	Epoch           int32    `json:"e"`
	Route           []NodeId `json:"r"`
	ChunkId         ChunkId  `json:"c"`
	Found           bool     `json:"f"`
	ThresholdFailed bool     `json:"t"`
	AccessFailed    bool     `json:"a"`
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
	UniqueWaitingCounter    int32
	UniqueRetryCounter      int32
	CacheHits               int32
	OriginatorIndex         int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
	Epoch                   int
}

func (s *State) GetOriginatorId(originatorIndex int32) NodeId {
	return s.Originators[originatorIndex]
}

type NodePairWithPrice struct {
	RequesterNode NodeId
	ProviderNode  NodeId
	Price         int
}

type PaymentWithPrice struct {
	Payment Payment
	Price   int
}

type Output struct {
	RouteWithPrices    []NodePairWithPrice
	PaymentsWithPrices []PaymentWithPrice
}

type Outputs struct {
	Outputs []Output
}
