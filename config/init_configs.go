package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"runtime"
)

// Variables This is the one that gets changed in setup
var Variables = defaultVariables

func InitConfigs() {
	ymlData := ReadYamlFile()
	SetNumGoroutines(ymlData.ConfOptions.NumGoroutines)
	SetConfOptions(ymlData.ConfOptions)
	SetExperiment(ymlData)
}

func ReadYamlFile() Yml {
	yamlFile, err := os.ReadFile("config.yaml")

	var yamlData Yml

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlData)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return yamlData
}

func SetExperiment(yml Yml) {

	switch yml.Experiment.ExperimentName {
	case "omega":
		fmt.Println("omega experiment chosen")
		OmegaExperiment()

		// TODO: Add more experiments here

	case "custom":
		fmt.Println("custom experiment chosen")
		CustomExperiment(yml.Custom)

	default:
		fmt.Println("default experiment chosen")
	}
}

func SetConfOptions(configOptions ConfVariables) {
	Variables.Iterations = configOptions.Iterations
	Variables.Bits = configOptions.Bits
	Variables.NetworkSize = configOptions.NetworkSize
	Variables.BinSize = configOptions.BinSize
	Variables.RangeAddress = configOptions.RangeAddress
	Variables.Originators = configOptions.Originators
	Variables.RefreshRate = configOptions.RefreshRate
	Variables.Threshold = configOptions.Threshold
	Variables.RandomSeed = configOptions.RandomSeed
	Variables.MaxProximityOrder = configOptions.MaxProximityOrder
	Variables.Price = configOptions.Price
	Variables.RequestsPerSecond = configOptions.RequestsPerSecond
	Variables.EdgeLock = configOptions.EdgeLock
	Variables.SameOriginator = configOptions.SameOriginator
	Variables.PrecomputeRespNodes = configOptions.PrecomputeRespNodes
	Variables.WriteRoutesToFile = configOptions.WriteRoutesToFile
	Variables.WriteStatesToFile = configOptions.WriteStatesToFile
	Variables.IterationMeansUniqueChunk = configOptions.IterationMeansUniqueChunk
	Variables.DebugPrints = configOptions.DebugPrints
	Variables.DebugInterval = configOptions.DebugInterval
}

func SetNumGoroutines(numGoroutines int) {
	if numGoroutines == -1 {
		Variables.NumGoroutines = runtime.NumCPU()
	} else {
		Variables.NumGoroutines = numGoroutines
	}
}
