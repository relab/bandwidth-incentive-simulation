package main

import (
	"flag"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	networkdata "go-incentive-simulation/network_data"
	"math/rand"
	"time"
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
	doubleBin := flag.Int("doubleBin", 0, "Number of nodes that use the double bin size")

	flag.Parse()

	if *useconfig {
		config.InitConfig()
		*binSize = config.GetBinSize()
		*bits = config.GetBits()
		*networkSize = config.GetNetworkSize()
		*rSeed = int(config.GetRandomSeed())
	}

	if *rSeed != -1 {
		rand.Seed(int64(*rSeed))
	} else {
		rand.Seed(time.Now().UnixNano())
	}

	if *count < 0 {
		filename := "../network_data/" + networkdata.GetNetworkDataName(*bits, *binSize, *networkSize, *id, -1)
		generateAndDump(*bits, *binSize, *doubleBin, *networkSize, *random, filename)
	}
	for i := 0; i < *count; i++ {
		filename := "../network_data/" + networkdata.GetNetworkDataName(*bits, *binSize, *networkSize, *id, i)
		generateAndDump(*bits, *binSize, *doubleBin, *networkSize, *random, filename)
	}
}

func generateAndDump(bits, binSize, doubleBin, N int, random bool, filename string) {

	network := types.Network{Bits: bits, Bin: binSize}
	network.Generate(N, doubleBin, random)

	err := network.Dump(filename)
	if err != nil {
		panic(fmt.Sprintf("dumping network to file gives error: %v", err))
	}
}
