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
	"reflect"
	"strings"
)

var (
	app       = kingpin.New("githubsearch", "GitHub Search").Author(Author).Version(Version)
	api       = app.Flag("api", "API type, format: "+strings.Join(search.Api, " or ")).Short('a').Required().String()
	config    = app.Flag("config", "Config file, format: .json").Short('c').Required().String()
	output    = app.Flag("output", "Output file, format: .json").Short('o').Required().String()
	qualifier = app.Flag("qualifier", "Qualifier list, format: "+
		"{qualifier}"+search.SyntaxSep+"{query}"+search.QualifierSep+
		"{qualifier}"+search.SyntaxSep+"{query}"+search.QualifierSep+"...").Short('q').Required().String()
	srch = app.Flag("search", "Search list, format: "+
		strings.Join(search.Type, search.SyntaxSep+"{text}"+" or ")+
		search.SyntaxSep+"{text}").Short('s').Required().String()
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

	qualifier, err := parseQualifier(*qualifier)
	if err != nil {
		log.Fatal("qualifier invalid: ", err.Error())
	}

	srch, err := parseSearch(*srch)
	if err != nil {
		log.Fatal("search invalid: ", err.Error())
	}

	result, err := runSearch(api, config, qualifier, srch)
	if err != nil {
		log.Fatal("search failed: ", err.Error())
	}

	if err := writeFile(*output, result); err != nil {
		log.Fatal("write failed: ", err.Error())
	}
}

func parseApi(data string) (string, error) {
	err := errors.New("data invalid")

	for _, item := range search.Api {
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

func removeDuplicate(data interface{}) interface{} {
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

func parseQualifier(data string) (interface{}, error) {
	helper := func(data string) (string, string, bool) {
		if !strings.Contains(data, search.SyntaxSep) {
			return "", "", false
		}
		buf := strings.Split(data, search.SyntaxSep)
		key := strings.TrimSpace(buf[0])
		val := strings.TrimSpace(buf[1])
		if key == "" || val == "" {
			return "", "", false
		}
		return key, val, true
	}

	qualifier := map[string][]interface{}{}

	buf := strings.Split(data, search.QualifierSep)
	for _, item := range buf {
		if item != "" {
			key, val, found := helper(item)
			if found {
				qualifier[key] = append(qualifier[key], val)
			}
		}
	}

	if len(qualifier) == 0 {
		return qualifier, errors.New("qualifier null")
	}

	for key := range qualifier {
		qualifier[key] = removeDuplicate(qualifier[key]).([]interface{})
	}

	return qualifier, nil
}

func parseSearch(data string) (interface{}, error) {
	helper := func(data string) (string, string, bool) {
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
		for _, item := range search.Type {
			if item == key {
				found = true
				break
			}
		}
		return key, val, found
	}

	srch := map[string][]interface{}{}

	buf := strings.Split(data, search.SearchSep)
	for _, item := range buf {
		if item != "" {
			key, val, found := helper(item)
			if found {
				srch[key] = append(srch[key], val)
			}
		}
	}

	if len(srch) == 0 {
		return srch, errors.New("search null")
	}

	if len(reflect.ValueOf(srch).MapKeys()) == len(search.Type) {
		return srch, errors.New("type inconsistent")
	}

	for key := range srch {
		srch[key] = removeDuplicate(srch[key]).([]interface{})
	}

	return srch, nil
}

func runSearch(api string, config, qualifier, srch interface{}) (interface{}, error) {
	return search.Run(api, config, qualifier, srch)
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
