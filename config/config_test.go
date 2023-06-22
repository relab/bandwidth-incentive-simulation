package config

import (
	"runtime"
	"testing"

	"gotest.tools/assert"
)

func TestReadDefaultConfig(t *testing.T) {
	fileconfig, err := ReadYamlFile("default_config.yaml")
	if err != nil {
		fileconfig, _ = ReadYamlFile("config/default_config.yaml")
	}
	defaultconfig := getDefaultConfig()
	assert.DeepEqual(t, fileconfig, defaultconfig)
}

func TestValidateBaseOptions(t *testing.T) {
	theconfig = getDefaultConfig()
	ValidateBaseOptions(theconfig.BaseOptions)
	assert.Equal(t, theconfig.BaseOptions.NumGoroutines, runtime.NumCPU())
	assert.Equal(t, theconfig.BaseOptions.OutputOptions.EvaluateInterval, theconfig.BaseOptions.Iterations)
}
