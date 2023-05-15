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

	flag.Parse()

	if *rSeed != -1 {
		rand.Seed(int64(*rSeed))
	}
	network := types.Network{Bits: *bits, Bin: *binSize}
	network.Generate(*networkSize)
	filename := fmt.Sprintf("nodes_data_%d_%d.txt", *binSize, *networkSize)
	if *id != "" {
		filename = fmt.Sprintf("nodes_data_%d_%d_%s.txt", *binSize, *networkSize, *id)
	}
	err := network.Dump(filename)
	if err != nil {
		return
	}
}
