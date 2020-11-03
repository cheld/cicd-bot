package config

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	//"gopkg.in/yaml.v2"
	//sigs.k8s.io/yaml"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Serve struct {
		Secret string
		Port   int
	}
	Events   []Event
	Triggers []Trigger
}

//func (config *Configuration) getTrigger(name string) Trigger {
//	for _, trigger := range config.Triggers {
//		if trigger.Name == name {
//			return trigger
//		}
//	}
//	return Trigger{}
//}

type EventInput struct {
	Objectiv string
	Input    map[string]interface{}
}

type Event struct {
	Source      string
	Type        string
	If_contains string
	If_equals   string
	If_true     string
	Trigger     string
	Values      map[string]interface{}
}

func (event *Event) IsMatching(eventInput EventInput) bool {
	contains := true
	if event.If_contains != "" {
		contains = strings.Contains(eventInput.Objectiv, event.If_contains)
	}
	equals := true
	if event.If_equals != "" {
		equals = eventInput.Objectiv == event.If_equals
	}
	condition := true
	if event.If_true != "" {
		result, _ := ProcessTemplate(event.If_true, eventInput)
		condition, _ = strconv.ParseBool(result)
	}
	return contains && equals && condition
}

func (event *Event) Process(eventInput EventInput) TriggerInput {
	triggerInput := TriggerInput{}
	triggerInput.Name = event.Trigger
	triggerInput.Values = ProcessAllTemplates(event.Values, eventInput).(map[string]interface{})
	return triggerInput
}

type TriggerInput struct {
	Name   string
	Values map[string]interface{}
}

type Trigger struct {
	Name      string
	Type      string
	Arguments map[string]interface{}
}

func Load(filename string) Configuration {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	var yamlConfig Configuration
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return yamlConfig
}
