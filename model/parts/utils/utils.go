package utils

import (
	"fmt"
	ct "go-incentive-simulation/model"
	"math/rand"
)

func MakeFiles() []int {
	fmt.Println("Making files...")
	var filesList []int

	// Gets all constants 
	consts := ct.Constants

	for i := 0; i <= consts.GetOriginators(); i++ {
		chunksList := rand.Perm(consts.GetChunks())
		filesList = append(chunksList)
	}
	fmt.Println("Files made!")
	return filesList
}

func (net *Network) CreateDowloadersList(fileName string) []int {
	fmt.Println("Creating downloaders list...")
	var downloadersList []int

	// nodes := net.nodes
	// downloadersList

	fmt.Println("Downloaders list create...!")
	return downloadersList
}

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
	// fmt.Println(nodes)


	fmt.Println("Pushing sync finished...")
}