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

import "testing"

func TestRunRest(t *testing.T) {
	rest := &Rest{}

	config := map[string]interface{}{
		"page":     2,
		"per_page": 10,
	}

	query := map[string][]interface{}{
		"code":  {"runSearch"},
		"owner": {"craftslab"},
		"repo":  {"githubsearch"},
	}

	if err := rest.Init(config); err != nil {
		t.Error("FAIL:", err)
	}

	if _, err := rest.Run(query); err != nil {
		t.Error("FAIL:", err)
	}
}
