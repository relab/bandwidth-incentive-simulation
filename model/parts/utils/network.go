package utils

type Network struct {
	bits  int
	bin   int
	nodes map[int]Node
}

type Node struct {
	net Network
	id  int
	// storageSet
	// cacheSet
	canPay bool
}

// func

// func (node *Node) add(other *Node) bool {
// 	net := node.net
// 	if (net == nil) || other.net() != net || node == other {
// 		return false
// 	}
// 	bits := net.bits - (node.id ^ other.id)
// }
