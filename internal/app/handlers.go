package service

import (
	"context"

	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"github.com/flypay/go-kit/v4/pkg/log"
	"github.com/flypay/go-kit/v4/pkg/projections"
)

// eventHandler is defined to have a CompositeReadWriter db. This allows the programme
// to search to data with data that is not the identifier.
type eventHandler struct {
	db projections.CompositeReadWriter
}

// BirthdayRecordProjection maps the structure of the data in the database
// used to add data to the db.
type BirthdayRecordProjection struct {
	Identifier string // Dob
	Range      string // Name
}

// createUser is an eventHandler. It recieves the UserCreated event. Applies it to uc object.
// Maps the object to the matching projection. Writes the data in the projection to the db.
// New user data is added to the db.
func (e eventHandler) createUser(ctx context.Context, event any, headers ...eventbus.Header) error {
	logger := log.WithContext(ctx)
	uc := event.(*bootcamp.UserCreated)
	up := BirthdayRecordProjection{
		Identifier: uc.DateOfBirth,
		Range:      uc.Id,
	}
	logger.Debugf("createUser recieved %v", up)
	err := e.db.Write(ctx, up, nil)
	if err != nil {
		return err
	}
	return nil
}
