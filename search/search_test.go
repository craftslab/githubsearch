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

func TestInitApis(t *testing.T) {
	if buf := initApis(); len(buf) == 0 {
		t.Error("FAIL")
	}
}

func TestRunSearch(t *testing.T) {
	config := map[string]interface{}{
		"page":     2,
		"per_page": 10,
	}

	query := map[string][]interface{}{
		"code":  {"runSearch"},
		"owner": {"craftslab"},
		"repo":  {"githubsearch"},
	}

	graphql := &GraphQl{}
	if _, err := runSearch(graphql, config, query); err != nil {
		t.Error("FAIL:", err)
	}

	rest := &Rest{}
	if _, err := runSearch(rest, config, query); err != nil {
		t.Error("FAIL:", err)
	}
}
