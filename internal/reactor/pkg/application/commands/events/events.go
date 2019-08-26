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

package events

import (
	"context"

	"go.zenithar.org/miam/internal/helpers"
	"go.zenithar.org/miam/internal/reactor/internal/meta"
	eventsv1 "go.zenithar.org/miam/pkg/gen/go/miam/events/v1"
)

const (
	aggregateType = "application"
)

// ApplicationCreated is raised when an application domain has been created.
func ApplicationCreated(ctx context.Context, id, label string) *eventsv1.Event {
	// Build event
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_APPLICATION_CREATED,
		EventId:       helpers.EventIDGeneratorFunc(),
		AggregateType: aggregateType,
		AggregateId:   id,
		Meta:          meta.FromContext(ctx),
		Payload: &eventsv1.Event_ApplicationCreated{
			ApplicationCreated: &eventsv1.ApplicationCreated{
				Id:    id,
				Label: label,
			},
		},
	}
}

// ApplicationActivated is raised when an application domain has been activated.
func ApplicationActivated(ctx context.Context, id string) interface{} {
	// Build event
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_APPLICATION_ACTIVATED,
		EventId:       helpers.EventIDGeneratorFunc(),
		AggregateType: aggregateType,
		AggregateId:   id,
		Meta:          meta.FromContext(ctx),
		Payload: &eventsv1.Event_ApplicationActivated{
			ApplicationActivated: &eventsv1.ApplicationActivated{
				Id: id,
			},
		},
	}
}

// ApplicationDeactivated is raised when an application domain has been deactivated.
func ApplicationDeactivated(ctx context.Context, id string) interface{} {
	// Build event
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_APPLICATION_DEACTIVATED,
		EventId:       helpers.EventIDGeneratorFunc(),
		AggregateType: aggregateType,
		AggregateId:   id,
		Meta:          meta.FromContext(ctx),
		Payload: &eventsv1.Event_ApplicationDeactivated{
			ApplicationDeactivated: &eventsv1.ApplicationDeactivated{
				Id: id,
			},
		},
	}
}

// ApplicationDeleted is raised when an application entity has been deleted.
func ApplicationDeleted(ctx context.Context, id string) interface{} {
	// Build event
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_APPLICATION_DELETED,
		EventId:       helpers.EventIDGeneratorFunc(),
		AggregateType: aggregateType,
		AggregateId:   id,
		Meta:          meta.FromContext(ctx),
		Payload: &eventsv1.Event_ApplicationDeleted{
			ApplicationDeleted: &eventsv1.ApplicationDeleted{
				Id: id,
			},
		},
	}
}

// ApplicationLabelChanged is raised when an application label attribute value has changed.
func ApplicationLabelChanged(ctx context.Context, id, old, new string) interface{} {
	// Build event
	return &eventsv1.Event{
		EventType:     eventsv1.EventType_EVENT_TYPE_APPLICATION_LABEL_UPDATED,
		EventId:       helpers.EventIDGeneratorFunc(),
		AggregateType: aggregateType,
		AggregateId:   id,
		Meta:          meta.FromContext(ctx),
		Payload: &eventsv1.Event_ApplicationLabelChanged{
			ApplicationLabelChanged: &eventsv1.ApplicationLabelChanged{
				Id:  id,
				Old: old,
				New: new,
			},
		},
	}
}
