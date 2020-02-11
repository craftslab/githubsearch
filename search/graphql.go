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

// GraphQl is search structure for the GraphQL API
type GraphQl struct {
	config map[string]interface{}
}

// Init is search initialization for the GraphQL API
func (g *GraphQl) Init(config map[string]interface{}) error {
	g.config = config
	return nil
}

// Run is search implementation for the GraphQL API
func (g GraphQl) Run(qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	return g.runGraphQl(qualifier, srch)
}

func (g GraphQl) runGraphQl(qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	// TODO
	return nil, nil
}
