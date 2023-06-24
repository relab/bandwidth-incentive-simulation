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

func InitConfig() {
	config, err := ReadYamlFile("config.yaml")
	if err != nil {
		log.Panicln("Unable to read config file: config.yaml")
	}
	theconfig = config
	ValidateBaseOptions(theconfig.BaseOptions)
	SetExperiment(theconfig)
}

func SetExperimentId(id string) {
	theconfig.BaseOptions.OutputOptions.ExperimentId = id
}

func SetMaxPO(maxPO int) {
	theconfig.BaseOptions.MaxProximityOrder = maxPO
}

func ReadYamlFile(filename string) (Config, error) {
	yamlFile, err := os.ReadFile(filename)

	var yamlData Config

	if err != nil {
		log.Printf("yamlFile.Get err :%v ", err)
		return yamlData, err
	}
	err = yaml.Unmarshal(yamlFile, &yamlData)
	if err != nil {
		log.Panicf("Unmarshal: %v", err)
	}
	return yamlData, nil
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
