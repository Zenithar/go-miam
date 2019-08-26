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

	"go.zenithar.org/miam/internal/reactor/internal/broker"
	"go.zenithar.org/miam/internal/reactor/pkg/application/commands/events"
	"go.zenithar.org/miam/internal/repositories"
	applicationv1 "go.zenithar.org/miam/pkg/gen/go/miam/application/v1"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
)

// DeleteHandler returns a Delete command event handler.
var DeleteHandler = func(reader repositories.ApplicationReader, creator repositories.ApplicationCreator, publisher broker.Publisher) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &applicationv1.DeleteResponse{}

		// Check request type
		req, ok := r.(*applicationv1.DeleteRequest)
		if !ok {
			return res, errors.Newf(errors.InvalidArgument, nil, "invalid request type for handler 'application.delete' (%T)", r)
		}

		// Validate request
		if err := req.Validate(); err != nil {
			return res, errors.Newf(errors.FailedPrecondition, err, "request is not valid")
		}

		// Check entity existence
		entity, err := reader.Get(ctx, req.Id)
		if err != nil && err != db.ErrNoResult {
			return res, errors.Newf(errors.Internal, err, "unable to query persistence")
		}
		if entity == nil {
			return res, errors.Newf(errors.NotFound, nil, "entity not found")
		}

		// Delete in persistence
		if err := creator.Delete(ctx, entity.ID); err != nil {
			return res, errors.Newf(errors.Internal, err, "unable to save entity")
		}

		// Publish event
		publisher.Publish(ctx, events.ApplicationDeleted(ctx, entity.URN()))

		return res, nil
	}
}
