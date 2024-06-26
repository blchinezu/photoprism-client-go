// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package test

import (
	"testing"

	photoprism "github.com/blchinezu/photoprism-client-go"
)

// TestHappyLogin should succeed with the good password "missy"
func TestHappyLogin(t *testing.T) {
	client := photoprism.New(WellKnownSampleAppConnectionString)
	err := client.Auth(photoprism.NewClientAuthLogin(WellKnownUser, WellKnownPass))
	if err != nil {
		t.Errorf("expected login: %v", err)
		t.FailNow()
	}
}

// TestSadLogin should fail with the bad password "charlie"
func TestSadLogin(t *testing.T) {
	client := photoprism.New(WellKnownSampleAppConnectionString)
	err := client.Auth(photoprism.NewClientAuthLogin(WellKnownUser, BadPassword))
	if err == nil {
		t.Errorf("expecting error for known bad password")
		t.FailNow()
	}
}
