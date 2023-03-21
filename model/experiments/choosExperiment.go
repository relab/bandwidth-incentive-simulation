package experiments

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Constant Constants

func ChooseExperiment() {
	yamlFile, err := os.ReadFile("run_experiments.yml")

	var experimentName string

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &experimentName)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	switch experimentName {
	case "noe":
		fmt.Println("noe")
		Constant = Experiment[1]
	default:
		fmt.Println("default")
		Constant = Experiment[2]
	}
}
