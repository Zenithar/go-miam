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

package helpers

import (
	"fmt"
	"sync"

	"go.zenithar.org/butcher"
)

var (
	err         error
	encoder     *butcher.Butcher
	once        sync.Once
	pepperBytes = []byte("KW@cqE/gD;DitGlFoW2OqA#sq`AS>LTb)i.$q6Z>YI$b,D{sCm3)_Ipd:T5ya;v")
)

func init() {
	once.Do(func() {
		encoder, err = butcher.New(
			butcher.WithAlgorithm(butcher.DefaultAlgorithm),
			butcher.WithSaltFunc(butcher.RandomNonce(32)),
			butcher.WithPepper(pepperBytes),
		)
		if err != nil {
			panic(err)
		}
	})
}

// PasswordEncodingFunc is used to encode password for storage.
var PasswordEncodingFunc = func(secret string) (string, error) {
	return encoder.Hash([]byte(secret))
}

// PasswordValidationFunc is used to check password match.
var PasswordValidationFunc = func(given, encoded string) (bool, error) {
	return encoder.Verify([]byte(encoded), []byte(given))
}

// SetPasswordPepper updates the pepper seed value.
func SetPasswordPepper(value []byte) {
	if len(value) < 32 {
		panic(fmt.Errorf("password pepper value must be be 32 long at least"))
	}
}
