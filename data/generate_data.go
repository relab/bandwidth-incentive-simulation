package main

import (
	"flag"
	"fmt"
	"go-incentive-simulation/model/parts/types"
	"math/rand"
)

func main() {

	binSize := flag.Int("binSize", 16, "Number of nodes in each address table bin, k in Kademlia")

	bits := flag.Int("bits", 16, "address length in bits")
	networkSize := flag.Int("N", 10000, "network size")
	rSeed := flag.Int("rSeed", -1, "random Seed")
	id := flag.String("id", "", "an id")
	count := flag.Int("count", 0, "generate count many networks with ids 0,1,...")

	flag.Parse()

	if *count == 0 {
		generateAndDump(*bits, *binSize, *networkSize, *rSeed, *id)
	}
	for i := *count; i > 0; i-- {
		generateAndDump(*bits, *binSize, *networkSize, *rSeed, fmt.Sprint(*count-i))
	}
}

func generateAndDump(bits, binSize, N, rSeed int, id string) {
	if rSeed != -1 {
		rand.Seed(int64(rSeed))
	}
	network := types.Network{Bits: bits, Bin: binSize}
	network.Generate(N)
	filename := fmt.Sprintf("nodes_data_%d_%d.txt", binSize, N)
	if id != "" {
		filename = fmt.Sprintf("nodes_data_%d_%d_%s.txt", binSize, N, id)
	}
	err := network.Dump(filename)
	if err != nil {
		return
	}
}
