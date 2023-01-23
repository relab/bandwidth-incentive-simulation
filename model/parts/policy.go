package policy

import (
	"sort"
)

func findResponisbleNodes(nodesId []int, chunkAdd int) []int {
	v := []int{}
	for i := range nodesId {
		v = append(v, nodesId[i]^chunkAdd)
	}
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })

	return v[:4]
}
