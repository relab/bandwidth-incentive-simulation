package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ChooseExperiment() {
	ymlData := ReadYamlFile()
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

func SetExperiment(experiment Yml) {
	//ConfOptions(experiment.ConfigOptions)

	switch experiment.Experiment.ExperimentName {
	case "omega":
		fmt.Println("omega experiment chose")
		OmegaExperiment()
	case "k20p20":
		fmt.Println("bucket size 20 and 20% originators experiment chose")
		BucketSize20And20pOriginators()
	case "k20p100":
		fmt.Println("bucket size 20 and 100% originators experiment chose")
		BucketSize20And100pOriginators()
	case "k16p20":
		fmt.Println("bucket size 16 and 20% originators experiment chose")
		BucketSize16And20pOriginators()
	case "k16p100":
		fmt.Println("bucket size 16 and 100% originators experiment chose")
		BucketSize16And100pOriginators()
	case "k8p20":
		fmt.Println("bucket size 8 and 20% originators experiment chose")
		BucketSize8And20pOriginators()
	case "k8p100":
		fmt.Println("bucket size 8 and 100% originators experiment chose")
		BucketSize8And100pOriginators()
	case "k4p20":
		fmt.Println("bucket size 4 and 20% originators experiment chose")
		BucketSize4And20pOriginators()
	case "k4p100":
		fmt.Println("bucket size 4 and 100% originators experiment chose")
		BucketSize4And100pOriginators()
	case "custom":
		fmt.Println("custom experiment chose")
		CustomExperiment(experiment.Custom)

	default:
		fmt.Println("default experiment chose")
	}
}
