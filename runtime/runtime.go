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

package runtime

import (
	"net/url"
)

// Operation type for the runtime
type Operation func(req *Request) interface{}

// Request structure for the runtime
type Request struct {
	Url string
	Val url.Values
}

type bundle struct {
	req  *Request
	resp chan interface{}
}

// Run is runtime implementation for the runtime
func Run(op Operation, req []Request) (interface{}, error) {
	return runRuntime(op, req)
}

func runRuntime(op Operation, req []Request) (interface{}, error) {
	helper := func(op Operation) (chan *bundle, chan bool) {
		data := make(chan *bundle)
		quit := make(chan bool)

		go routine(op, data, quit)

		return data, quit
	}

	channel, quit := helper(op)

	bundles := make([]bundle, len(req))

	for i := 0; i < len(req); i++ {
		b := &bundles[i]
		b.req = &req[i]
		b.resp = make(chan interface{})
		channel <- b
	}

	buf := make([]interface{}, len(req))

	for i := 0; i < len(req); i++ {
		buf = append(buf, <-bundles[i].resp)
	}

	quit <- true

	return buf, nil
}

func routine(op Operation, data chan *bundle, quit chan bool) {
	for {
		select {
		case buf := <-data:
			go operation(op, buf)
		case <-quit:
			return
		}
	}
}

func operation(op Operation, data *bundle) {
	data.resp <- op(data.req)
}
