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

package models

import (
	"fmt"
	"time"

	"go.zenithar.org/miam/internal/helpers"
)

// Application describes a resource provider in term of OIDC.
type Application struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	Active   bool      `json:"active"`
	CreateAt time.Time `json:"created_at"`
}

// NewApplication returns an application instance.
func NewApplication(label string) *Application {
	return &Application{
		ID:       helpers.IDGeneratorFunc(),
		Label:    label,
		Active:   false,
		CreateAt: helpers.TimeFunc(),
	}
}

// -----------------------------------------------------------------------------

// URN returns the entity URN
func (app *Application) URN() string {
	return fmt.Sprintf("miam:v1::application:%s", app.ID)
}
