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
	"go.zenithar.org/pkg/types"
)

// Client describes a resource consumer in term of OIDC.
type Client struct {
	ID            string            `json:"id"`
	Label         string            `json:"label"`
	Secret        string            `json:"secret"`
	RedirectURIs  types.StringArray `json:"redirect_uris"`
	GrantTypes    types.StringArray `json:"grant_types"`
	ResponseTypes types.StringArray `json:"response_types"`
	Scopes        types.StringArray `json:"scopes"`
	Audience      types.StringArray `json:"audience"`
	Public        bool              `json:"public"`
	CreateAt      time.Time         `json:"created_at"`
	Active        bool              `json:"active"`
}

// NewClient returns an client instance.
func NewClient(label string) *Client {
	return &Client{
		ID:       helpers.IDGeneratorFunc(),
		Label:    label,
		Active:   false,
		CreateAt: helpers.TimeFunc(),
	}
}

// -----------------------------------------------------------------------------

// SetSecret updates the secret attribute value.
func (c *Client) SetSecret(secret string) error {
	encoded, err := helpers.PasswordEncodingFunc(secret)
	if err != nil {
		return err
	}

	c.Secret = encoded
	return nil
}

// URN returns the entity URN
func (c *Client) URN() string {
	return fmt.Sprintf("miam:v1::client:%s", c.ID)
}
