package commands

import (
	"context"

	"go.zenithar.org/miam/internal/models"
	"go.zenithar.org/miam/internal/reactor/pkg/application/commands/events"
	"go.zenithar.org/miam/internal/repositories"
	applicationv1 "go.zenithar.org/miam/pkg/gen/go/miam/application/v1"

	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
)

// CreateHandler returns a Create command event handler.
var CreateHandler = func(creator repositories.ApplicationCreator) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		// Check request type
		req, ok := r.(*applicationv1.CreateRequest)
		if !ok {
			return nil, errors.Newf(errors.InvalidArgument, nil, "invalid request type for handler 'application.create' (%T)", r)
		}

		// Validate request
		if err := req.Validate(); err != nil {
			return nil, errors.Newf(errors.FailedPrecondition, err, "request is not valid")
		}

		// Prepare model
		entity := models.NewApplication(req.Label)

		// Create in persistence
		if err := creator.Create(ctx, entity); err != nil {
			return nil, errors.Newf(errors.Internal, err, "unable to save entity")
		}

		// Publish event
		events.ApplicationCreated(entity.URN(), entity.Label)

		// Prepare result and return
		return nil, nil
	}
}
