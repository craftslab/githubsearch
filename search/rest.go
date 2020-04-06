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
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/craftslab/githubsearch/runtime"
	"github.com/pkg/errors"
)

const (
	host = "https://api.github.com"
)

// Request structure for the runtime
type Request struct {
	Url string
	Val url.Values
}

// Rest is search structure for the REST API
type Rest struct {
	config map[string]interface{}
}

var (
	options = []string{"order", "page", "per_page", "sort"}
	urls    = []string{host + "/search/code", host + "/search/repositories"}
)

// Init is search initialization for the REST API
func (r *Rest) Init(config map[string]interface{}) error {
	r.config = config
	return nil
}

// Run is search implementation for the REST API
func (r Rest) Run(qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	return r.runRest(qualifier, srch)
}

func (r Rest) runRest(qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	req, err := r.request(qualifier, srch)
	if err != nil {
		return nil, errors.Wrap(err, "request invalid")
	}

	return runtime.Run(r.operation, req)
}

func (r Rest) request(qualifier, srch map[string][]interface{}) ([]interface{}, error) {
	helper := func(_type, srch, query, option string) (Request, error) {
		_url := ""
		for _, u := range urls {
			if strings.Contains(u, _type) {
				_url = u
				break
			}
		}
		if _url == "" {
			return Request{}, errors.New("url invalid")
		}
		req := Request{
			Url: _url + "?q=" + srch + query + option,
			Val: nil,
		}
		return req, nil
	}

	var req []interface{}
	qry := r.query(qualifier)
	opt := r.option(r.config)

	for key, val := range srch {
		for _, item := range val {
			r, err := helper(key, item.(string), qry, opt)
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

func (r Rest) query(data map[string][]interface{}) string {
	if len(data) == 0 {
		return ""
	}

	var ret []string

	for key, val := range data {
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
}

func (r Rest) option(data map[string]interface{}) string {
	if len(data) == 0 {
		return ""
	}

	var ret []string

	for _, item := range options {
		if _, present := data[item]; !present {
			continue
		}

		switch val := data[item].(type) {
		case int:
			ret = append(ret, "&"+item+"="+strconv.Itoa(val))
		case string:
			ret = append(ret, "&"+item+"="+val)
		default:
			// TODO
		}
	}

	return strings.Join(ret, "")
}

func (r Rest) operation(req interface{}) interface{} {
	// TODO: req.Val
	resp, err := http.Get(req.(Request).Url)
	if err != nil {
		return nil
	}

	defer func() { _ = resp.Body.Close() }()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return buf
}
