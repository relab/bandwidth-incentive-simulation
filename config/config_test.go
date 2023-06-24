package config

import (
	"runtime"
	"testing"

	"gotest.tools/assert"
)

func TestReadDefaultConfig(t *testing.T) {
	fileconfig, _ := ReadYamlFile("testdata/default_config.yaml")
	defaultconfig := getDefaultConfig()
	assert.DeepEqual(t, fileconfig, defaultconfig)
}

func TestValidateBaseOptions(t *testing.T) {
	theconfig = getDefaultConfig()
	ValidateBaseOptions(theconfig.BaseOptions)
	assert.Equal(t, theconfig.BaseOptions.NumGoroutines, runtime.NumCPU())
	assert.Equal(t, theconfig.BaseOptions.OutputOptions.EvaluateInterval, theconfig.BaseOptions.Iterations)
}

func TestGetters(t *testing.T) {
	SetDefaultConfig()

	assert.Equal(t, GetBinSize(), 16)
}
