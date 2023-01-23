package policy

import (
	"testing"
	"gotest.tools/assert"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{2, 11, 8, 6, 5, 4, 3}
	chunkAdd := 11
	values := findResponisbleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}