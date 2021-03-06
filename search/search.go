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
	"github.com/pkg/errors"
)

// Search separator for the CLI
const (
	QualifierSep = ","
	SearchSep    = ","
	SyntaxSep    = ":"
)

// Search interface for the API
type Search interface {
	Init(config map[string]interface{}) error
	Run(qualifier, srch map[string][]interface{}) ([]interface{}, error)
}

// Search APIs and types for the CLI
var (
	Api  = initApi()
	Type = []string{"code", "repo"}
)

var (
	searches = map[string]Search{
		"graphql": &GraphQl{},
		"rest":    &Rest{},
	}
)

// Run is search implementation for the API
func Run(api string, config map[string]interface{}, qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	cfg, present := config[api]
	if !present {
		return nil, errors.New("config invalid")
	}

	return runSearch(searches[api], cfg.(map[string]interface{}), qualifier, srch)
}

func initApi() []string {
	var buf []string

	for key := range searches {
		buf = append(buf, key)
	}

	return buf
}

func runSearch(s Search, config map[string]interface{}, qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	if err := s.Init(config); err != nil {
		return nil, errors.Wrap(err, "init failed")
	}

	return s.Run(qualifier, srch)
}
