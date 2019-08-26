package meta

import (
	"context"

	eventsv1 "go.zenithar.org/miam/pkg/gen/go/miam/events/v1"
)

type mdCtxKeyType string

const (
	userCtxKey    mdCtxKeyType = "user_id"
	requestCtxKey mdCtxKeyType = "request_id"
)

// FromContext extracts metadat from context.
func FromContext(ctx context.Context) *eventsv1.Meta {
	m := &eventsv1.Meta{}

	// Extract userid
	user, ok := ctx.Value(userCtxKey).(string)
	if ok {
		m.UserId = user
	}

	return m
}

// WithUser updates the user_id attributes of the context.
func WithUser(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userCtxKey, userID)
}
