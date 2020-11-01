package destination

import (
	"fmt"

	"github.com/cheld/cicd-bot/pkg/config"
)

type Dispatcher struct {
	config  config.Configuration
	targets map[string]interface{}
}

func NewDispatcher(config config.Configuration) *Dispatcher {
	dispatcher := Dispatcher{}
	dispatcher.config = config
	dispatcher.targets = make(map[string]interface{})
	for _, trigger := range config.Destinations.Debug.Stdout {
		dispatcher.targets[trigger.Name] = trigger
	}
	fmt.Println(dispatcher.targets)
	return &dispatcher
}

func (dispatcher *Dispatcher) Execute(eventData []config.EventData) {
	for _, data := range eventData {
		target := dispatcher.targets[data.Name]
		switch v := target.(type) {
		case config.DebugStdout:
			Stdout(target.(config.DebugStdout), data)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}

}
