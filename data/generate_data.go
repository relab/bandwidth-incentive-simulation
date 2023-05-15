package main

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"math/rand"
)

func main() {
	config.InitConfigs()
	binSize := config.GetBinSize()
	bits := config.GetBits()
	networkSize := config.GetNetworkSize()
	rand.Seed(config.GetRandomSeed())
	network := types.Network{Bits: bits, Bin: binSize}
	network.Generate(networkSize)
	filename := fmt.Sprintf("./nodes_data_%d_%d.txt", binSize, networkSize)
	err := network.Dump(filename)
	if err != nil {
		return
	}
}
