//
// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtcontainers

import (
	"testing"
)

func TestNoopShimStart(t *testing.T) {
	s := &noopShim{}
	sandbox := Sandbox{}
	params := ShimParams{}
	expected := 0

	pid, err := s.start(sandbox, params)
	if err != nil {
		t.Fatal(err)
	}

	if pid != expected {
		t.Fatalf("PID should be %d", expected)
	}
}