package types

import (
	"errors"
	// "fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"math/rand"
	"sync"
	// "time"
)

type Node struct {
	Network          *Network
	Id               NodeId
	Active           bool
	AdjIds           [][]NodeId
	OriginatorStruct OriginatorStruct
	CacheStruct      CacheStruct
	PendingStruct    PendingStruct
	RerouteStruct    RerouteStruct
	AdjLock          sync.RWMutex
}

// Adds a one-way connection from node to other
func (node *Node) add(other *Node) (bool, error) {
	if node.Network == nil || node.Network != other.Network {
		return false, errors.New("trying to add nodes with different networks")
	}
	if node == other {
		return false, nil
	}
	if !other.Active {
		return false, nil
	}

	node.AdjLock.Lock()
	defer node.AdjLock.Unlock()

	if node.AdjIds == nil {
		node.AdjIds = make([][]NodeId, node.Network.Bits)
	}
	bit := node.Network.Bits - general.BitLength(node.Id.ToInt()^other.Id.ToInt())
	if bit < 0 || bit >= node.Network.Bits {
		return false, errors.New("nodes have distance outside XOR metric")
	}
	if len(node.AdjIds[bit]) < node.Network.Bin && !general.Contains(node.AdjIds[bit], other.Id) {
		node.AdjIds[bit] = append(node.AdjIds[bit], other.Id)
		return true, nil
	}
	return false, nil
}

func (node *Node) UpdateNeighbors() {
	return
	// start := time.Now()
	node.AdjLock.Lock()
	defer node.AdjLock.Unlock()

	candidateNeighbors := make([][]NodeId, node.Network.Bits)
	for l, adjIds := range node.AdjIds {

		shuffledAdjIds := adjIds
		rand.Shuffle(len(shuffledAdjIds), func(i, j int) {
			shuffledAdjIds[i], shuffledAdjIds[j] = shuffledAdjIds[j], shuffledAdjIds[i]
		})
		if len(shuffledAdjIds) > 4 {
			shuffledAdjIds = shuffledAdjIds[:4]
		}

		for _, adjId := range shuffledAdjIds {
			if !general.Contains(candidateNeighbors[l], adjId) {
				candidateNeighbors[l] = append(candidateNeighbors[l], adjId)
			}
			adj := node.Network.NodesMap[adjId]
			adj.AdjLock.RLock()
			for _, adjAdjIds := range adj.AdjIds {

				shuffledAdjAdjIds := adjAdjIds
				rand.Shuffle(len(shuffledAdjAdjIds), func(i, j int) {
					shuffledAdjAdjIds[i], shuffledAdjAdjIds[j] = shuffledAdjAdjIds[j], shuffledAdjAdjIds[i]
				})
				if len(shuffledAdjAdjIds) > 4 {
					shuffledAdjAdjIds = shuffledAdjAdjIds[:4]
				}

				for _, adjAdjId := range shuffledAdjAdjIds {
					bin := config.GetBits() - general.BitLength(node.Id.ToInt()^adjAdjId.ToInt())
					if adjAdjId != node.Id && !general.Contains(candidateNeighbors[bin], adjAdjId) {
						candidateNeighbors[bin] = append(candidateNeighbors[bin], adjAdjId)
					}
				}
			}
			adj.AdjLock.RUnlock()
		}
	}

	for d := 0; d < node.Network.Bits; d++ {
		rand.Shuffle(len(candidateNeighbors[d]), func(i, j int) {
			candidateNeighbors[d][i], candidateNeighbors[d][j] = candidateNeighbors[d][j], candidateNeighbors[d][i]
		})
		if len(candidateNeighbors[d]) > node.Network.Bin {
			node.AdjIds[d] = candidateNeighbors[d][:node.Network.Bin]
		} else {
			node.AdjIds[d] = candidateNeighbors[d]
		}
	}

	// fmt.Println("Time taken:", time.Since(start))
}

func (node *Node) IsNil() bool {
	return node.Id == 0
}

func (node *Node) Deactivate() {
	node.Active = false
}

func (node *Node) Activate() {
	node.Active = true
}
