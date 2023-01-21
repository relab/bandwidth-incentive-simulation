package utils

import (
	"fmt"
	ct "go-incentive-simulation/model"
	"math/rand"
	"time"
)

func CreateGraphNetwork(filename string) (*Graph, error) {
	fmt.Println("Creating graph network...")
	graph := &Graph{
		edges: make(map[int][]*Edge),
	}
	net := new(Network)
	_, _, nodes := net.load(filename)
	for _, node := range nodes {
		err := graph.AddNode(node)
		if err != nil {
			return nil, err
		}
	}
	for _, node := range graph.Nodes() {
		nodeAdj := node.adj
		for _, adjItems := range nodeAdj {
			for _, item := range adjItems {
				// "a2b" show how much this node asked from other node,
				// "last" is for the last forgiveness time
				attrs := EdgeAttrs{a2b: 0, last: 0}
				edge := Edge{fromNodeId: node.id, toNodeId: item.id, attrs: attrs}
				err := graph.AddEdge(&edge)
				if err != nil {
					return nil, err
				}
				// graph.SetEdgeAttributes()
			}
		}
	}
	fmt.Println("Graph network is created.")
	return graph, nil
}

// Python version stores the GetEdgeData function on the net class, instead of the graph... maybe change later?
func isThresholdFailed(firstNode *Node, secondNode *Node, chunkId int, g *Graph) bool {
	if ct.Constants.GetThresholdEnabled() {
		edgeDataFirst := g.GetEdgeData(firstNode, secondNode)
		p2pFirst := edgeDataFirst.a2b

		edgeDataSecond := g.GetEdgeData(secondNode, firstNode)
		p2pSecond := edgeDataSecond.a2b

		price := p2pFirst - p2pSecond + peerPriceChunk(secondNode, chunkId)
		fmt.Printf("price: %d", price)
		return price > ct.Constants.GetThreshold()
	}
	return false
}

// func getNext(firstNode *Node, chunkId int, graph *Graph, mainOriginator *Node, prevNodePaid *Node, rerouteMap map[int][]*Node) {
// 	var nextNode *Node
// 	nextNode = nil
// 	var payNext *Node
// 	payNext = nil
// 	var thresholdList [][2]*Node
// 	thresholdFailed := false
// 	accessFailed := false
// 	payment := 0
// 	lastDistance := int(firstNode.id ^ chunkId)
// 	fmt.Printf("last distance is : %d, chunk is: %d, first is: %d", lastDistance, chunkId, firstNode.id)
// 	fmt.Printf("which bucket: %d", 16-BitLength(chunkId^firstNode.id))

// 	currDist := int(lastDistance)
// 	payDist := int(lastDistance)
// 	for _, adj := range firstNode.adj {
// 		fmt.Printf("adj: %d", adj)
// 		for _, node := range adj {
// 			dist := int(node.id ^ chunkId)
// 			if BitLength(dist) >= BitLength(lastDistance) {
// 				continue
// 			}

// 			if !isThresholdFailed(firstNode, node, chunkId, graph) {
// 				thresholdFailed = false

// 				// Could probably clean this one up, but keeping it close to original for now
// 				if dist < currDist {
// 					if ct.Constants.IsRetryWithAnotherPeer() {
// 						_, ok := rerouteMap[mainOriginator.id]
// 						if ok {
// 							allExceptLast := len(rerouteMap[mainOriginator.id]) - 1
// 							if containsNode(rerouteMap[mainOriginator.id][:allExceptLast], node) {
// 								continue
// 							} else {
// 								currDist = dist
// 								nextNode = node
// 							}
// 						} else {
// 							currDist = dist
// 							nextNode = node
// 						}
// 					} else {
// 						currDist = dist
// 						nextNode = node
// 					}
// 				}
// 			} else {
// 				thresholdFailed = true
// 				if ct.Constants.GetPaymentEnabled() {
// 					if dist < payDist {
// 						payDist = dist
// 						payNext = node
// 					}
// 				}
// 				listItem := [2]*Node{firstNode, node}
// 				thresholdList = append(thresholdList, listItem)
// 			}
// 		}
// 	}
// 	// FORTSETT HER!
// 	if nextNode != nil {
// 		thresholdFailed = false
// 		accessFailed = false
// 	} else {
// 		if !thresholdFailed {
// 			accessFailed = true
// 		}
// 	}
// }

func getBin(src int, dest int, index int) int {
	distance := src ^ dest
	result := index
	for distance > 0 {
		distance >>= 1
		result -= 1
	}
	return result
}

func whichPowerTwo(rangeAddress int) int {
	return BitLength(rangeAddress) - 1
}

func getProximityChunk(firstNode *Node, chunkId int) int {
	retVal := ct.Constants.GetBits() - BitLength(firstNode.id^chunkId)
	if retVal <= ct.Constants.GetMaxProximityOrder() {
		return retVal
	} else {
		return ct.Constants.GetMaxProximityOrder()
	}
}

func peerPriceChunk(firstNode *Node, chunkId int) int {
	return (ct.Constants.GetMaxProximityOrder() - getProximityChunk(firstNode, chunkId) + 1) * ct.Constants.GetPrice()
}

func choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < k; i++ {
		res = append(res, nodes[rand.Intn(len(nodes))])
	}
	return res
}

func MakeFiles() []int {
	fmt.Println("Making files...")
	var filesList []int

	for i := 0; i <= ct.Constants.GetOriginators(); i++ {
		// TODO: fix this, GetChuncks should be a list?
		// chunksList := choice(ct.Constants.GetChunks(), ct.Constants.GetRangeAddress())
		// filesList = append(chunksList)
		fmt.Println(i)
	}
	// Gets all constants
	consts := ct.Constants

	for i := 0; i <= consts.GetOriginators(); i++ {
		chunksList := rand.Perm(consts.GetChunks())
		filesList = append(chunksList)
	}
	fmt.Println("Files made!")
	return filesList
}

func (net *Network) CreateDowloadersList() []int {
	fmt.Println("Creating downloaders list...")

	nodesValue := make([]int, 0, len(net.nodes))
	for i := range net.nodes {
		nodesValue = append(nodesValue, net.nodes[i].id)
	}
	downloadersList := choice(nodesValue, ct.Constants.GetOriginators())

	fmt.Println("Downloaders list create...!")
	return downloadersList
}

// no need for this function
func (net *Network) PushSync(fileName string, files []string) {
	fmt.Println("Pushing sync...")
	if net == nil {
		fmt.Println("Network is nil!")
		return
	}
	nodes := net.nodes
	for i := range nodes {
		fmt.Println(nodes[i].id)
	}

	fmt.Println("Pushing sync finished...")
}
