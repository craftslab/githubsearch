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

import "github.com/craftslab/githubsearch/runtime"

const (
	urlApi  = "https://api.github.com"
	urlCode = urlApi + "/search/code"
	urlRepo = urlApi + "/search/repositories"
)

const (
	optionOrder   = "order"
	optionPage    = "page"
	optionPerPage = "per_page"
	optionSort    = "sort"
)

// Rest is search structure for the REST API
type Rest struct {
	config interface{}
}

// Init is search initialization for the REST API
func (r *Rest) Init(config interface{}) error {
	r.config = config
	return nil
}

// Run is search implementation for the REST API
func (r Rest) Run(qualifier, srch interface{}) (interface{}, error) {
	return r.runRest(qualifier, srch)
}

func (r Rest) runRest(qualifier, srch interface{}) (interface{}, error) {
	// TODO
	req := []runtime.Request{{Url: "", Val: nil}}
	return runtime.Run(operation, req)
}

func operation(req *runtime.Request) interface{} {
	return nil
}
