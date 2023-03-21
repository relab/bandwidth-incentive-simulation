package experiments

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Experiment() {
	ymlData := ReadYamlFile()
	SetExperiment(ymlData)
}

func ReadYamlFile() Yml {
	yamlFile, err := os.ReadFile("run_experiments.yaml")

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

func SetExperiment(experiment Yml) {
	switch experiment.Experiment.ExperimentName {
	case "omega":
		fmt.Println("omega experiment chose")
		OmegaExperiment()
	case "custom":
		fmt.Println("custom experiment chose")
		CustomExperiment(experiment.Custom)
	default:
		fmt.Println("default")
	}
}
