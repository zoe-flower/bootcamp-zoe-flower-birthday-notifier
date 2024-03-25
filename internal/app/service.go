package service

import (

	// "github.com/flypay/events/pkg/flyt"
	"fmt"

	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"github.com/flypay/go-kit/v4/pkg/projections"
)

func RunService(
	defaultDatabase projections.ReadWriter,
	eventbus eventbus.Consumer,
) error {
	// consumer.On(...)
	// defer consumer.Listen()

	// NOTE: For each SNS topic you produce or consume you will need to modify
	// the service.json file to specify the topic in the 'eventbus.consumes_topics' list.
	// E.g., for &flyt.UserCreated, add "flyt-user-created" (dashes instead of dots in "flyt.user.created" from the proto) to 'eventbus.consumes_topics' in service.json

	eh := eventHandler{db: defaultDatabase}
	err := eventbus.On(&bootcamp.UserCreated{}, eh.xxx)
	if err != nil {
		fmt.Println("HERE", err)
	}
	defer eventbus.Listen()
	return nil
}
