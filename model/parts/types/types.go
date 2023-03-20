package types

type Constants struct {
	Runs                             int
	Bits                             int
	NetworkSize                      int
	BinSize                          int
	RangeAddress                     int
	Originators                      int
	RefreshRate                      int
	Threshold                        int
	RandomSeed                       int64
	MaxProximityOrder                int
	Price                            int
	Chunks                           int
	RequestsPerSecond                int
	ThresholdEnabled                 bool
	ForgivenessEnabled               bool
	ForgivenessDuringRouting         bool
	PaymentEnabled                   bool
	MaxPOCheckEnabled                bool
	WaitingEnabled                   bool
	OnlyOriginatorPays               bool
	PayOnlyForCurrentRequest         bool
	PayIfOrigPays                    bool
	ForwarderPayForceOriginatorToPay bool
	RetryWithAnotherPeer             bool
	CacheIsEnabled                   bool
	PreferredChunks                  bool
	AdjustableThreshold              bool
	EdgeLock                         bool
	SameOriginator                   bool
	PrecomputeRespNodes              bool
	WriteRoutesToFile                bool
	WriteStatesToFile                bool
	IterationMeansUniqueChunk        bool
	DebugPrints                      bool
	DebugInterval                    int
	NumRoutingGoroutines             int
	Epoch                            int
}

type Experiments map[int]Constants

type Request struct {
	TimeStep        int
	Epoch           int
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
	Epoch                   int
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
