// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"context"

	"go.zenithar.org/miam/internal/models"
	"go.zenithar.org/miam/internal/reactor/internal/broker"
	"go.zenithar.org/miam/internal/reactor/pkg/application/commands/events"
	"go.zenithar.org/miam/internal/reactor/pkg/application/mapper"
	"go.zenithar.org/miam/internal/repositories"
	applicationv1 "go.zenithar.org/miam/pkg/gen/go/miam/application/v1"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
)

// CreateHandler returns a Create command event handler.
var CreateHandler = func(creator repositories.ApplicationCreator, reader repositories.ApplicationReader, publisher broker.Publisher) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &applicationv1.CreateResponse{}

		// Check request type
		req, ok := r.(*applicationv1.CreateRequest)
		if !ok {
			return res, errors.Newf(errors.InvalidArgument, nil, "invalid request type for handler 'application.create' (%T)", r)
		}

		// Validate request
		if err := req.Validate(); err != nil {
			return res, errors.Newf(errors.FailedPrecondition, err, "request is not valid")
		}

		// Check if label is not used
		saved, err := reader.FindByLabel(ctx, req.Label)
		if err != nil && err != db.ErrNoResult {
			return res, errors.Newf(errors.Internal, err, "unable to query persistence")
		}
		if saved != nil {
			return res, errors.Newf(errors.AlreadyExists, nil, "entity label '%s' already used", req.Label)
		}

		// Prepare model
		entity := models.NewApplication(req.Label)

		// Create in persistence
		if err := creator.Create(ctx, entity); err != nil {
			return res, errors.Newf(errors.Internal, err, "unable to save entity")
		}

		// Publish event
		publisher.Publish(ctx, events.ApplicationCreated(ctx, entity.URN(), entity.Label))

		// Prepare result and return
		res.Application = mapper.FromEntity(entity)

		return res, nil
	}
}
