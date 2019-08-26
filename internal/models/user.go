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

// User describes a user identity.
type User struct {
	ID        string    `json:"id"`
	Principal string    `json:"principal"`
	Active    bool      `json:"active"`
	CreateAt  time.Time `json:"created_at"`
}

// NewUser returns an user instance.
func NewUser(principal string) *User {
	return &User{
		ID:        helpers.IDGeneratorFunc(),
		Principal: principal,
		Active:    false,
		CreateAt:  helpers.TimeFunc(),
	}
}

// -----------------------------------------------------------------------------

// URN returns the entity URN
func (u *User) URN() string {
	return fmt.Sprintf("miam:v1::user:%s", u.ID)
}
