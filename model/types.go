package model

import (
	. "go-incentive-simulation/model/parts/utils"
)

type Request struct {
	Originator *Node
	ChunkId    int
}

type CacheListMap map[*Node][]map[int]int

type RerouteMap map[int][]*Node

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]*Node
