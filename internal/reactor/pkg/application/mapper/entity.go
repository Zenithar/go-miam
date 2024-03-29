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

package mapper

import (
	"go.zenithar.org/miam/internal/models"
	applicationv1 "go.zenithar.org/miam/pkg/gen/go/miam/application/v1"
)

// FromEntity return a value object from entity.
func FromEntity(entity *models.Application) *applicationv1.Application {
	return &applicationv1.Application{
		Id:    entity.ID,
		Urn:   entity.URN(),
		Label: entity.Label,
	}
}
