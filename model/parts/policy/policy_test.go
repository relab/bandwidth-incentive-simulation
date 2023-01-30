package policy

import (
	"gotest.tools/assert"
	"testing"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{8190, 11683, 11211, 16935, 21020, 21725, 39525, 41162, 41471, 41852, 42779, 43377, 43812, 45280, 47987, 49693, 57841, 59951, 64132}
	chunkAdd := 43000
	values := findResponsibleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}
