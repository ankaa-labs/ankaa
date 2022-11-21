// Copyright (c) 2022 The ankaa-labs Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ankaa-labs/ankaa/test"
	"gopkg.in/yaml.v3"
)

func transform(
	file string,
	unmarshal func(data []byte, out interface{}) error,
	marshal func(in interface{}) ([]byte, error),
) {
	srcData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("ERR: cannot read file %v, %v, skip it", file, err)
		return
	}

	var srcObj interface{}
	err = unmarshal(srcData, &srcObj)
	if err != nil {
		log.Printf("ERR: cannot unmarshal file %v to json, %v, skip it", file, err)
		return
	}

	destData, err := marshal(srcObj)
	if err != nil {
		log.Printf("ERR: cannot marshal file %v to json, %v, skip it", file, err)
		return
	}

	err = os.Rename(file, file+".bak")
	if err != nil {
		log.Printf("ERR: cannot rename file %v to backup, %v, skip it", file, err)
		return
	}

	err = ioutil.WriteFile(file, destData, 0644)
	if err != nil {
		log.Printf("ERR: cannot write to file %v, %v, skip it", file, err)
		return
	}

	err = os.Remove(file + ".bak")
	if err != nil {
		log.Printf("ERR: cannot remove file %v backup, %v, skip it", file, err)
		return
	}

	log.Printf("format file %v done", file)
}

func main() {
	dir := path.Join(test.CurrentProjectPath(), "test", "resources")
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0, len(dirEntries))
	for _, entry := range dirEntries {
		if entry.IsDir() {
			log.Printf("%s is directory, skip it", entry.Name())
			continue
		}

		files = append(files, path.Join(dir, entry.Name()))
	}

	for _, file := range files {
		if strings.HasSuffix(file, ".json") {
			transform(file, json.Unmarshal, func(in interface{}) ([]byte, error) {
				return json.MarshalIndent(in, "", "    ")
			})
		} else if strings.HasSuffix(file, ".yaml") {
			transform(file, yaml.Unmarshal, yaml.Marshal)
		} else {
			log.Printf("file %v is not json or yaml format, skip it", file)
		}
	}
}
