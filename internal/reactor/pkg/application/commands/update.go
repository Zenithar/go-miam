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
	"go.zenithar.org/miam/internal/reactor/pkg/application/mapper"
	"go.zenithar.org/miam/internal/repositories"
	applicationv1 "go.zenithar.org/miam/pkg/gen/go/miam/application/v1"
	eventsv1 "go.zenithar.org/miam/pkg/gen/go/miam/events/v1"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/errors"
	"go.zenithar.org/pkg/reactor"
)

// UpdateHandler returns a Update command event handler.
var UpdateHandler = func(reader repositories.ApplicationReader, updater repositories.ApplicationUpdater, publisher broker.Publisher) reactor.HandlerFunc {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		res := &applicationv1.CreateResponse{}

		// Check request type
		req, ok := r.(*applicationv1.UpdateRequest)
		if !ok {
			return res, errors.Newf(errors.InvalidArgument, nil, "invalid request type for handler 'application.update' (%T)", r)
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

		// Prepare updates
		eventList := []*eventsv1.Event{}

		// Check request attributes
		if req.Label != nil {
			if entity.Label != req.Label.Value {
				// Check if label is not used
				saved, err := reader.FindByLabel(ctx, req.Label.Value)
				if err != nil && err != db.ErrNoResult {
					return res, errors.Newf(errors.Internal, err, "unable to query persistence")
				}
				if saved != nil {
					return res, errors.Newf(errors.AlreadyExists, nil, "entity label '%s' already used", req.Label.Value)
				}

				// Add event to publication list
				eventList = append(eventList, events.ApplicationLabelChanged(ctx, entity.URN(), entity.Label, req.Label.Value))

				// Assign entity attribute value
				entity.Label = req.Label.Value
			}
		}

		if req.Active != nil {
			if entity.Active != req.Active.Value {
				// Add event to publication list
				if req.Active.Value {
					eventList = append(eventList, events.ApplicationActivated(ctx, entity.URN()))
				} else {
					eventList = append(eventList, events.ApplicationDeactivated(ctx, entity.URN()))
				}

				// Assign entity attribute value
				entity.Active = req.Active.Value
			}
		}

		// If updated do the update
		if len(eventList) > 0 {
			// Create in persistence
			if err := updater.Update(ctx, entity); err != nil {
				return res, errors.Newf(errors.Internal, err, "unable to update entity")
			}

			// Publish all events
			for _, evt := range eventList {
				publisher.Publish(ctx, evt)
			}
		}

		// Return result
		res.Application = mapper.FromEntity(entity)

		return res, nil
	}
}
