package main

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"math/rand"
)

func main() {
	binSize := constants.GetBinSize()
	bits := constants.GetBits()
	networkSize := constants.GetNetworkSize()
	rand.Seed(constants.GetRandomSeed())
	network := types.Network{Bits: bits, Bin: binSize}
	network.Generate(networkSize)
	filename := fmt.Sprintf("nodes_data_%d_%d.txt", binSize, networkSize)
	err := network.Dump(filename)
	if err != nil {
		return
	}
}
