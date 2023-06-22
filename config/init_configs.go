package config

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

// theconfig This is the current configuration.
var theconfig Config

func InitConfigs() {
	theconfig := ReadYamlFile("config.yaml")
	ValidateBaseOptions(theconfig.BaseOptions)
	SetExperiment(theconfig)
}

func InitConfigsWithId(id string) {
	InitConfigs()
	theconfig.BaseOptions.OutputOptions.ExperimentId = id
}

func SetMaxPO(maxPO int) {
	theconfig.BaseOptions.MaxProximityOrder = maxPO
}

func ReadYamlFile(filename string) Config {
	yamlFile, err := os.ReadFile(filename)

	var yamlData Config

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlData)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return yamlData
}

func SetExperiment(yml Config) {

	switch yml.Experiment.Name {
	case "omega":
		fmt.Println("omega experiment chosen")
		OmegaExperiment()

	case "custom":
		fmt.Println("custom experiment chosen")
		CustomExperiment(yml.ExperimentOptions)

	default:
		fmt.Println("default experiment chosen")
	}
}

func ValidateBaseOptions(configOptions baseOptions) {
	SetNumGoroutines(configOptions.NumGoroutines)
	SetEvaluateInterval(configOptions.OutputOptions.EvaluateInterval)
}

func SetNumGoroutines(numGoroutines int) {
	if numGoroutines == -1 {
		theconfig.BaseOptions.NumGoroutines = runtime.NumCPU()
	}
}

func SetEvaluateInterval(interval int) {
	if interval <= 0 {
		theconfig.BaseOptions.OutputOptions.EvaluateInterval = theconfig.BaseOptions.Iterations
	}
}
