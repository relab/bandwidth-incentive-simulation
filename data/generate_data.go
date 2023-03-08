package main

import (
	"fmt"
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	"math/rand"
)

func main() {
	binSize := Constants.GetBinSize()
	bits := Constants.GetBits()
	networkSize := Constants.GetNetworkSize()
	rand.Seed(Constants.GetRandomSeed())
	network := Network{Bits: bits, Bin: binSize}
	network.Generate(networkSize)
	filename := fmt.Sprintf("nodes_data_%d_%d.txt", binSize, networkSize)
	network.Dump(filename)
}