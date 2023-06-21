package config

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

// Variables This is the one that gets changed in setup
var Variables = GetDefaultVariables()

func InitConfigs() {
	ymlData := ReadYamlFile()
	SetConfOptions(ymlData.ConfOptions)
	SetExperiment(ymlData)
}

func InitConfigsWithId(id string) {
	InitConfigs()
	Variables.confOptions.OutputOptions.ExperimentId = id
}

func SetMaxPO(maxPO int) {
	Variables.confOptions.MaxProximityOrder = maxPO
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

	switch yml.Experiment.Name {
	case "omega":
		fmt.Println("omega experiment chosen")
		OmegaExperiment()

		// TODO: Add more experiments here

	case "custom":
		fmt.Println("custom experiment chosen")
		CustomExperiment(yml.CustomExperiment)

	default:
		fmt.Println("default experiment chosen")
	}
}

func SetConfOptions(configOptions confOptions) {
	Variables.confOptions = configOptions
	SetNumGoroutines(configOptions.NumGoroutines)
}

func SetNumGoroutines(numGoroutines int) {
	if numGoroutines == -1 {
		Variables.confOptions.NumGoroutines = runtime.NumCPU()
	} else {
		Variables.confOptions.NumGoroutines = numGoroutines
	}
}
