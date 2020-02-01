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

package cmd

import (
	"os"
	"testing"
)

func TestParseApi(t *testing.T) {
	if _, err := parseApi(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseApi("graphql"); err != nil {
		t.Error("FAIL:", err)
	}

	if _, err := parseApi("rest"); err != nil {
		t.Error("FAIL:", err)
	}

	if _, err := parseApi("invalid"); err == nil {
		t.Error("FAIL")
	}
}

func TestParseConfig(t *testing.T) {
	if _, err := parseConfig(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseConfig("../config/search.json"); err != nil {
		t.Error("FAIL:", err)
	}
}

func TestParseOutput(t *testing.T) {
	if err := parseOutput(""); err == nil {
		t.Error("FAIL")
	}

	if err := parseOutput("../config/search.json"); err == nil {
		t.Error("FAIL")
	}

	if err := parseOutput("output.json"); err != nil {
		t.Error("FAIL:", err)
	}
}

func checkDuplicates(data []interface{}) bool {
	found := false
	key := map[string]bool{}

	for _, item := range data {
		if _, isPresent := key[item.(string)]; isPresent {
			found = true
			break
		}
	}

	return found
}

func TestRemoveDuplicates(t *testing.T) {
	buf := []interface{}{"code:runSearch", "code:runSearch"}
	buf = removeDuplicates(buf).([]interface{})

	if found := checkDuplicates(buf); found {
		t.Error("FAIL")
	}
}

func TestParseQuery(t *testing.T) {
	if _, err := parseQuery(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseQuery("code"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseQuery("code:"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseQuery("code:runSearch,owner:craftslab,repo:githubsearch"); err != nil {
		t.Error("FAIL:", err)
	}

	query, err := parseQuery("code:runSearch,code:runSearch")
	if err != nil {
		t.Error("FAIL:", err)
	}

	buf := query.(map[string][]interface{})
	if len(buf) != 1 || len(buf["code"]) != 1 {
		t.Error("FAIL")
	}

	query, err = parseQuery("code:runGraphQl,code:runRest")
	if err != nil {
		t.Error("FAIL:", err)
	}

	buf = query.(map[string][]interface{})
	if len(buf) != 1 || len(buf["code"]) != 2 {
		t.Error("FAIL")
	}
}

func TestRunSearch(t *testing.T) {
	config := map[string]interface{}{
		"graphql": map[string]interface{}{
			"page":     2,
			"per_page": 10,
		},
		"rest": map[string]interface{}{
			"page":     2,
			"per_page": 10,
		},
	}

	query := map[string][]interface{}{
		"code":  {"runSearch"},
		"owner": {"craftslab"},
		"repo":  {"githubsearch"},
	}

	if _, err := runSearch("", config, query); err == nil {
		t.Error("FAIL")
	}

	if _, err := runSearch("invalid", config, query); err == nil {
		t.Error("FAIL")
	}

	if _, err := runSearch("graphql", config, query); err != nil {
		t.Error("FAIL:", err)
	}

	if _, err := runSearch("rest", config, query); err != nil {
		t.Error("FAIL:", err)
	}
}

func TestWriteFile(t *testing.T) {
	buf := map[string]interface{}{
		"k": "key",
		"v": "value",
	}

	name := "tmp.json"

	if err := writeFile("", buf); err == nil {
		t.Error("FAIL")
	}

	if err := writeFile(name, func() {}); err == nil {
		t.Error("FAIL")
	}

	err := writeFile(name, buf)
	defer func(name string) { _ = os.Remove(name) }(name)
	if err != nil {
		t.Error("FAIL:", err)
	}
}
