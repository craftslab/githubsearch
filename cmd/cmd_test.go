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

func checkDuplicate(data []interface{}) bool {
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

func TestRemoveDuplicate(t *testing.T) {
	buf := []interface{}{"code:runSearch", "code:runSearch"}
	buf = removeDuplicate(buf)

	if found := checkDuplicate(buf); found {
		t.Error("FAIL")
	}
}

func TestParseQualifier(t *testing.T) {
	if _, err := parseQualifier(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseQualifier("in"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseQualifier("in:"); err == nil {
		t.Error("FAIL")
	}

	qualifier, err := parseQualifier("in:file,in:file")
	if err != nil {
		t.Error("FAIL:", err)
	}

	if len(qualifier) != 1 || len(qualifier["in"]) != 1 {
		t.Error("FAIL")
	}

	qualifier, err = parseQualifier("in:file,in:path")
	if err != nil {
		t.Error("FAIL:", err)
	}

	if len(qualifier) != 1 || len(qualifier["in"]) != 2 {
		t.Error("FAIL")
	}

	if _, err := parseQualifier("in:file,language:go,repo:githubsearch,user:craftslab"); err != nil {
		t.Error("FAIL:", err)
	}
}

func TestParseSearch(t *testing.T) {
	if _, err := parseSearch(""); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseSearch("code"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseSearch("code:"); err == nil {
		t.Error("FAIL")
	}

	if _, err := parseSearch("code:runSearch,repo:githubsearch"); err == nil {
		t.Error("FAIL")
	}

	srch, err := parseSearch("code:runSearch,code:runSearch")
	if err != nil {
		t.Error("FAIL:", err)
	}

	if len(srch) != 1 || len(srch["code"]) != 1 {
		t.Error("FAIL")
	}

	srch, err = parseSearch("code:parseSearch,code:runSearch")
	if err != nil {
		t.Error("FAIL:", err)
	}

	if len(srch) != 1 || len(srch["code"]) != 2 {
		t.Error("FAIL")
	}
}

func TestRunSearch(t *testing.T) {
	config := map[string]interface{}{
		"rest": map[string]interface{}{
			"order":    "desc",
			"page":     2,
			"per_page": 10,
			"sort":     "stars",
		},
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

	if _, err := runSearch("rest", config, qualifier, srch); err != nil {
		t.Error("FAIL:", err)
	}
}

func TestWriteFile(t *testing.T) {
	if err := writeFile("rest", "", []interface{}{}); err == nil {
		t.Error("FAIL")
	}

	name := "tmp.json"
	buf := []interface{}{
		[]byte(`{"key": "value"}`),
	}

	err := writeFile("rest", name, buf)
	defer func(name string) { _ = os.Remove(name) }(name)
	if err != nil {
		t.Error("FAIL:", err)
	}
}
