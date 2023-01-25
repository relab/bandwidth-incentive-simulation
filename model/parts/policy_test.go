package policy

import (
	"testing"

	"gotest.tools/assert"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{64132, 49693, 45280, 42779, 41852, 43812, 47987, 43377, 41471}
	chunkAdd := 11
	values := findResponsibleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}


