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
	"encoding/json"
	"github.com/craftslab/githubsearch/search"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	app    = kingpin.New("githubsearch", "GitHub Search").Author(Author).Version(Version)
	api    = app.Flag("api", "API type, type: "+strings.Join(search.Apis, " ")).Short('a').Required().String()
	config = app.Flag("config", "Config file, format: .json").Short('c').Required().String()
	output = app.Flag("output", "Output file, format: .json").Short('o').Required().String()
	query  = app.Flag("query", "Query type, format: "+strings.Join(search.Types, search.SyntaxSep+"{QUERY}"+
		search.TypeSep)+search.SyntaxSep+"{QUERY}").Short('q').Required().String()
)

// Run is search implementation for the CLI
func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	api, err := parseApi(*api)
	if err != nil {
		log.Fatal("api invalid: ", err.Error())
	}

	config, err := parseConfig(*config)
	if err != nil {
		log.Fatal("config invalid: ", err.Error())
	}

	if err = parseOutput(*output); err != nil {
		log.Fatal("output invalid: ", err.Error())
	}

	query, err := parseQuery(*query)
	if err != nil {
		log.Fatal("query invalid: ", err.Error())
	}

	result, err := runSearch(api, config, query)
	if err != nil {
		log.Fatal("search failed: ", err.Error())
	}

	if err := writeFile(*output, result); err != nil {
		log.Fatal("write failed: ", err.Error())
	}
}

func parseApi(data string) (string, error) {
	err := errors.New("data invalid")

	for _, item := range search.Apis {
		if item == data {
			err = nil
			break
		}
	}

	return data, err
}

func parseConfig(name string) (interface{}, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}

	buf, _ := ioutil.ReadAll(file)

	if err := file.Close(); err != nil {
		return nil, errors.Wrap(err, "close failed")
	}

	result := map[string]interface{}{}

	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, errors.Wrap(err, "unmarshal failed")
	}

	return result, nil
}

func parseOutput(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name null")
	}

	if _, err := os.Stat(name); err == nil {
		return errors.New("name exist")
	}

	return nil
}

func removeDuplicates(data interface{}) interface{} {
	var buf []interface{}
	key := map[string]bool{}

	for _, item := range data.([]interface{}) {
		if _, present := key[item.(string)]; !present {
			key[item.(string)] = true
			buf = append(buf, item.(string))
		}
	}

	return buf
}

func parseQuery(data string) (interface{}, error) {
	_parseQuery := func(data string) (string, string, bool) {
		if !strings.Contains(data, search.SyntaxSep) {
			return "", "", false
		}
		buf := strings.Split(data, search.SyntaxSep)
		key := strings.TrimSpace(buf[0])
		val := strings.TrimSpace(buf[1])
		if key == "" || val == "" {
			return "", "", false
		}
		found := false
		for _, item := range search.Types {
			if item == key {
				found = true
				break
			}
		}
		return key, val, found
	}

	query := map[string][]interface{}{}

	buf := strings.Split(data, search.TypeSep)
	for _, item := range buf {
		if item != "" {
			key, val, found := _parseQuery(item)
			if found {
				query[key] = append(query[key], val)
			}
		}
	}

	if len(query) == 0 {
		return query, errors.New("query null")
	}

	for key := range query {
		query[key] = removeDuplicates(query[key]).([]interface{})
	}

	return query, nil
}

func runSearch(api string, config, query interface{}) (interface{}, error) {
	cfg, present := config.(map[string]interface{})[api]
	if !present {
		return nil, errors.New("api invalid")
	}

	return search.Run(api, cfg, query)
}

func writeFile(name string, data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	if err := ioutil.WriteFile(name, buf, 0644); err != nil {
		return errors.Wrap(err, "write failed")
	}

	return nil
}
