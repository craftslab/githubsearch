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
	SyntaxSep = ":"
	TypeSep   = ","
)

// Search interface for the API
type Search interface {
	Init(config interface{}) error
	Run(query interface{}) (interface{}, error)
}

// Search APIs and types for the CLI
var (
	Apis  = initApis()
	Types = []string{"code", "owner", "repo"}
)

var (
	searches = map[string]Search{
		"graphql": &GraphQl{},
		"rest":    &Rest{},
	}
)

// Run is search implementation for the API
func Run(api string, config, query interface{}) (interface{}, error) {
	return runSearch(searches[api], config, query)
}

func initApis() []string {
	var buf []string

	for key := range searches {
		buf = append(buf, key)
	}

	return buf
}

func runSearch(s Search, config, query interface{}) (interface{}, error) {
	if err := s.Init(config); err != nil {
		return nil, errors.Wrap(err, "init failed")
	}

	return s.Run(query)
}
