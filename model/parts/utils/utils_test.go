package utils

import (
	"testing"
)

func TestCreateGraphNetwork(t *testing.T) {
	// fileName := "input_test.txt"
	fileName := "nodes_data_8_10000.txt"

	CreateGraphNetwork(fileName)

}
