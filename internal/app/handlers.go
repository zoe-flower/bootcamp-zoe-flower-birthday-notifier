package service

import (
	"context"

	"github.com/flypay/events/pkg/bootcamp"
	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"github.com/flypay/go-kit/v4/pkg/log"
	"github.com/flypay/go-kit/v4/pkg/projections"
)

type eventHandler struct {
	db projections.ReadWriter
}

type UserCreatedProjection struct {
	Identifier string // primary key
	Dob        string
}

func (e eventHandler) xxx(ctx context.Context, event any, headers ...eventbus.Header) error {
	logger := log.WithContext(ctx)
	uc := event.(*bootcamp.UserCreated)
	up := UserCreatedProjection{
		Identifier: uc.Id,
		Dob:        uc.DateOfBirth,
	}
	logger.Debugf("xxx recieved %v", event)
	err := e.db.Write(ctx, up, nil)
	if err != nil {
		return err
	}
	return nil
}
