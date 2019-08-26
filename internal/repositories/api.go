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

package repositories

import (
	"context"

	"go.zenithar.org/miam/internal/models"
)

// ApplicationCreator describes application creator contract.
type ApplicationCreator interface {
	Create(ctx context.Context, entity *models.Application) error
	Delete(ctx context.Context, id string) error
}

// ApplicationUpdater describes application updator contract.
type ApplicationUpdater interface {
	Update(ctx context.Context, entity *models.Application) error
}

// ApplicationReader describes application reader contract.
type ApplicationReader interface {
	Get(ctx context.Context, id string) (*models.Application, error)
	FindByLabel(ctx context.Context, label string) (*models.Application, error)
}
