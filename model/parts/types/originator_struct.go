package types

type OriginatorStruct struct {
	Active        bool
	RequestCount  int
}

func (o *OriginatorStruct) AddRequest() {
	o.RequestCount++
}

func (o *OriginatorStruct) Deactivate() {
	o.Active = false
}

func (o *OriginatorStruct) Activate() {
	o.Active = true
}
