package reactor

import (
	"context"

	eventsv1 "go.zenithar.org/miam/pkg/gen/go/miam/events/v1"
)

// Publisher describes event publisher broker interface.
type Publisher interface {
	Publish(ctx context.Context, evt *eventsv1.Event)
}
