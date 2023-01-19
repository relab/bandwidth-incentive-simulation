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

func (n *Network) CreateDowloadersList(fileName string) []int {
	fmt.Println("Creating downloaders list...")
	var downloadersList []int

	// net := n.load(fileName)
	// nodes := net.nodes
	// downloadersList

	fmt.Println("Downloaders list create...!")
	return downloadersList
}