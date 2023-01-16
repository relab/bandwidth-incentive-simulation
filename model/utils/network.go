package network

type Network struct {
	bits  int
	bin   int
	nodes map[int]Nodes
}

type Node struct {
	net Network
	id  int
	storageSet
	cacheSet
	canPay bool
}

func (node *Node) add(other *Node) bool {
	net := node.net
	// if (net == nil) || o
}

//  def add(self, other):
//         net = self.net()
//         if net is None or other.net() is not net or self is other:
//             return False
//         bit = net.bits - (self.id ^ other.id).bit_length()
//         if len(self.adj[bit]) < net.bin > len(other.adj[bit]):
//             other.adj[bit].add(self)
//             self.adj[bit].add(other)
//             return True
//         return False
