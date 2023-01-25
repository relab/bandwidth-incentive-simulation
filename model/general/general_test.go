package general

import (
	"gotest.tools/assert"
	"testing"
)

func TestChoice(t *testing.T) {
	// List of nodes
	nodes := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Originators
	k := 2
	c := Choice(nodes, k)
	assert.Equal(t, len(c), k)
}
