package config

import (
	"testing"

	"gotest.tools/assert"
)

func TestDefaultConfig(t *testing.T) {
	fileconfig := ReadYamlFile("default_config.yaml")
	defaultconfig := getDefaultConfig()
	assert.DeepEqual(t, fileconfig, defaultconfig)
}
