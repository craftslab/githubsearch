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
	"github.com/craftslab/githubsearch/runtime"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	host = "https://api.github.com"
)

// Rest is search structure for the REST API
type Rest struct {
	config interface{}
}

var (
	options = []string{"order", "page", "per_page", "sort"}
	urls    = []string{host + "/search/code", host + "/search/repositories"}
)

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
	req, err := r.request(qualifier, srch)
	if err != nil {
		return nil, errors.Wrap(err, "request invalid")
	}

	return runtime.Run(r.operation, req)
}

func (r Rest) request(qualifier, srch interface{}) ([]runtime.Request, error) {
	helper := func(_type, srch, query, option string) (runtime.Request, error) {
		url := ""
		for _, u := range urls {
			if strings.Contains(u, _type) {
				url = u
				break
			}
		}
		if url == "" {
			return runtime.Request{}, errors.New("url invalid")
		}
		req := runtime.Request{
			Url: url + "?q=" + srch + query + option,
			Val: nil,
		}
		return req, nil
	}

	option := func(data interface{}) string {
		if _, present := data.(map[string]interface{}); !present {
			return ""
		}
		buf := data.(map[string]interface{})
		if len(buf) == 0 {
			return ""
		}
		var ret []string
		for _, item := range options {
			if _, present := buf[item]; !present {
				continue
			}
			switch val := buf[item].(type) {
			case int:
				ret = append(ret, "&"+item+"="+strconv.Itoa(val))
			case string:
				ret = append(ret, "&"+item+"="+val)
			default:
				// TODO
			}
		}
		return strings.Join(ret, "")
	}(r.config)

	query := func(data interface{}) string {
		if _, present := data.(map[string][]interface{}); !present {
			return ""
		}
		buf := data.(map[string][]interface{})
		if len(buf) == 0 {
			return ""
		}
		var ret []string
		for key, val := range buf {
			for _, item := range val {
				if _, present := item.(string); present {
					ret = append(ret, "+"+key+":"+item.(string))
				} else {
					// TODO
					continue
				}
			}
		}
		return strings.Join(ret, "")
	}(qualifier)

	var req []runtime.Request

	for key, val := range srch.(map[string][]interface{}) {
		for _, item := range val {
			r, err := helper(key, item.(string), query, option)
			if err != nil {
				return nil, err
			}
			req = append(req, r)
		}
	}

	if len(req) == 0 {
		return nil, errors.New("request null")
	}

	return req, nil
}

func (r Rest) operation(req *runtime.Request) interface{} {
	// TODO
	return nil
}
