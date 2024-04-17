package service

import (
	"context"
	"testing"

	"github.com/flypay/go-kit/v4/pkg/projections"
)

var testcases = []struct {
	name    string
	context context.Context
	event   any
}{
	{"Writes data to db", context.Background(), 3},
	{"Maps event data to Projection", context.Background(), 4},
	{"third", context.Background(), 6},
}

func TestCreateUser(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			a := &eventHandler{
				db: &projections.MemCompositeDB{},
			}
			got := a.createUser(tc.context, tc.event)
			expected := 1 // work out expected
			if got != expected {
				t.Errorf("expected %v, got %v", expected, got)
			}
		})
	}
}
