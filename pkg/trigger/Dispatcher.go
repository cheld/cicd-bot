package trigger

import (
	"fmt"

	"github.com/cheld/cicd-bot/pkg/config"
)

type Dispatcher struct {
	config config.Configuration
}

func NewDispatcher(cfg config.Configuration) Dispatcher {
	dispatcher := Dispatcher{}
	dispatcher.config = cfg
	return dispatcher
}

func (dispatcher *Dispatcher) Execute(tasks []config.Task) error {
	for _, task := range tasks {
		trigger := dispatcher.config.Trigger(task.Trigger)
		if trigger == nil {
			return fmt.Errorf("No trigger definition with name '%s' found\n", task.Trigger)
		}
		switch trigger.Type {
		case "debug":
			ExecuteDebug(trigger, task)
		case "http":
			ExecuteHttp(trigger, task)
		default:
			return fmt.Errorf("No implementation for trigger type '%s' found!\n", trigger.Type)
		}
	}
	return nil
}
