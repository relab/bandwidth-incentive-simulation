package utils

type Network struct {
	bits  int
	bin   int
	nodes map[int]*Node
}

type Node struct {
	net Network
	id  int
	adj [][]*Node
	// storageSet
	// cacheSet
	canPay bool
}

func (node *Node) AddNode(other *Node) bool {
	return true
}
