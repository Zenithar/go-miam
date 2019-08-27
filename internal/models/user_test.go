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

package models_test

import (
	"fmt"
	"testing"
	"time"

	"go.zenithar.org/miam/internal/models"

	. "github.com/onsi/gomega"
)

func TestUser_Creation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewUser("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Principal).To(Equal("foo"), "Entity should have the matching principal")
	g.Expect(obj.Active).To(BeFalse(), "Entity must be deactivated by default")
	g.Expect(obj.CreateAt).Should(BeTemporally("~", time.Now(), time.Second), "Entity should have creation date")
	g.Expect(obj.URN()).To(Equal(fmt.Sprintf("miam:v1::user:%s", obj.ID)), "Entity urn should be as expected")

	obj.SetSecret("foobar")
	g.Expect(obj.Secret).ToNot(BeEmpty(), "Entity secret should not be empty")
}

func TestUser_Validate(t *testing.T) {
	tcl := []struct {
		name      string
		principal string
		wantErr   bool
	}{
		{
			name:      "Invalid principal: empty input",
			principal: "",
			wantErr:   true,
		},
		{
			name:      "Invalid principal: invalid charset",
			principal: "日本語",
			wantErr:   true,
		},
	}

	for _, tc := range tcl {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			underTest := models.NewUser(tt.principal)

			err := underTest.Validate()
			if tt.wantErr && err == nil {
				t.Errorf("Error expected, but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Error not expected, but got %v", err)
			}
		})
	}
}
