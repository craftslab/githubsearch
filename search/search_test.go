// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"testing"
)

func TestRun(t *testing.T) {
	config := map[string]interface{}{
		"rest": map[string]interface{}{
			"order":    "desc",
			"page":     2,
			"per_page": 10,
			"sort":     "stars",
		},
	}

	if _, err := Run("invalid", config, nil, nil); err == nil {
		t.Error("FAIL")
	}
}

func TestInitApi(t *testing.T) {
	if buf := initApi(); len(buf) == 0 {
		t.Error("FAIL")
	}
}

func TestRunSearch(t *testing.T) {
	rest := &Rest{}

	config := map[string]interface{}{
		"order":    "desc",
		"page":     2,
		"per_page": 10,
		"sort":     "stars",
	}

	qualifier := map[string][]interface{}{
		"in":       {"file"},
		"language": {"go"},
		"repo":     {"githubsearch"},
		"user":     {"craftslab"},
	}

	srch := map[string][]interface{}{
		"code": {"runSearch"},
		"repo": {"githubsearch"},
	}

	if _, err := runSearch(rest, config, qualifier, srch); err != nil {
		t.Error("FAIL:", err)
	}
}
