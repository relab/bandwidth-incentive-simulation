package main

import (
	"flag"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"math/rand"
)

func main() {

	binSize := flag.Int("binSize", 16, "Number of nodes in each address table bin, k in Kademlia")

	bits := flag.Int("bits", 16, "address length in bits")
	networkSize := flag.Int("N", 10000, "network size")
	rSeed := flag.Int("rSeed", -1, "random Seed")
	id := flag.String("id", "", "an id")
	count := flag.Int("count", -1, "generate count many networks with ids i0,i1,...")
	random := flag.Bool("random", true, "spread nodes randomly")
	useconfig := flag.Bool("config", false, "use config.yaml to initialize bits, binSize, NetworkSize and randomness")

	flag.Parse()

	if *useconfig {
		config.InitConfigs()
		*binSize = config.GetBinSize()
		*bits = config.GetBits()
		*networkSize = config.GetNetworkSize()
		*rSeed = int(config.GetRandomSeed())
	}

	if *count == 0 {
		filename := GetNetworkDataName(*bits, *binSize, *networkSize, *id, -1)
		generateAndDump(*bits, *binSize, *networkSize, *rSeed, *random, filename)
	}
	for i := 0; i < *count; i++ {
		filename := GetNetworkDataName(*bits, *binSize, *networkSize, *id, i)
		generateAndDump(*bits, *binSize, *networkSize, *rSeed, *random, filename)
	}
}

func generateAndDump(bits, binSize, N, rSeed int, random bool, filename string) {
	if rSeed != -1 {
		rand.Seed(int64(rSeed))
	}
	network := types.Network{Bits: bits, Bin: binSize}
	network.Generate(N, random)

	err := network.Dump(filename)
	if err != nil {
		panic(fmt.Sprintf("dumping network to file gives error: %v", err))
	}
}

func GetNetworkDataName(bits, binSize, N int, id string, iteration int) string {
	if iteration >= 0 {
		iterstr := fmt.Sprintf("i%d", iteration)
		if len(id) > 0 {
			id = id + "-" + iterstr
		} else {
			id = iterstr
		}
	}

	return fmt.Sprintf("nodes_data_b%d_k%d_%d_%s.txt", bits, binSize, N, id)
}
