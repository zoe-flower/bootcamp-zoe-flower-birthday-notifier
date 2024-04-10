package service

import (

	// "github.com/flypay/events/pkg/flyt"
	"fmt"

	"github.com/flypay/bootcamp-zoe-flower-birthday-notifier/internal/httpservice"
	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"github.com/flypay/go-kit/v4/pkg/log"
	"github.com/flypay/go-kit/v4/pkg/projections"
	"github.com/labstack/echo/v4"
)

// eventHandler is defined to have a CompositeReadWriter db. This allows the programme to search to data with data that is not the identifier.
// Needs http to be able to call on the end point to run the BirthdayToday function. // Needs to consume the UserCreated event // Needs to produce
// the BirthdayToday event

// passed the db property on the eh instance of the eventHandler struct an actual database value.
// then calls the handler to handle the event with the incoming event data.
func RunService(
	compositeDatabase projections.CompositeReadWriter,
	http *echo.Echo,
	consumer eventbus.Consumer,
	producer eventbus.Producer,
) error {
	// NOTE: For each SNS topic you produce or consume you will need to modify
	// the service.json file to specify the topic in the 'eventbus.consumes_topics' list.
	// E.g., for &flyt.UserCreated, add "flyt-user-created" (dashes instead of dots in "flyt.user.created" from the proto) to 'eventbus.consumes_topics' in service.json
	httpservice.RegisterHandlers(http,
		httpservice.HTTPHandler{Logger: log.DefaultLogger, CompositeDatabase: compositeDatabase, Producer: producer},
	)
	eh := eventHandler{db: compositeDatabase}
	err := consumer.On(&bootcamp.UserCreated{}, eh.createUser)
	if err != nil {
		fmt.Println("HERE", err)
	}
	// TO FINISH: Why is the listening defered if the its starts handling the event already above. How does this not error?
	// thought you would need to listen for the event, to then be able to take and use it.
	defer consumer.Listen()
	return nil
}
