package event

import (
	"encoding/json"
	"fmt"

	"github.com/cheld/cicd-bot/pkg/config"
)

func (handler *Handler) HandleHttp(body []byte, path string) []config.Task {

	// parse payload
	var payload interface{}
	if len(body) > 0 {
		err := json.Unmarshal(body, &payload)
		if err != nil {
			fmt.Printf("Not possible to parse request body %s", string(body))
			return []config.Task{}
		}
	} else {
		payload = ""
	}
	source := config.Source{
		Value:   string(body),
		Payload: payload,
		Environ: handler.env,
	}

	// handle event
	event := handler.config.FindEvent("http", path, source)
	if event == nil {
		fmt.Printf("No event found for value %s\n", source.Value)
		return []config.Task{}
	}

	// build execution task
	task, err := event.BuildTask(source)
	if err != nil {
		fmt.Printf("Cannot handle event: %v. Error: %v", event.Trigger, err)
		return []config.Task{}
	}
	return []config.Task{task}
}
