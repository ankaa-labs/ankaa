package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/ankaa-labs/ankaa/test"
)

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func transform(
	files []string,
	srcFormat string,
	destFormat string,
	unmarshal func(data []byte, out interface{}) error,
	marshal func(in interface{}) ([]byte, error),
) {
	for _, srcFile := range files {
		if !strings.HasSuffix(srcFile, srcFormat) {
			log.Printf("%s is not %s format, skip it", srcFile, srcFormat)
			continue
		}

		destFile := srcFile[0:len(srcFile)-len(srcFormat)] + destFormat
		if _, err := os.Stat(destFile); err == nil {
			log.Printf("ERR: the target file %v exists, skip it", destFile)
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			log.Printf("ERR: stat target file %v, %v, skip it", destFile, err)
			continue
		}

		srcData, err := ioutil.ReadFile(srcFile)
		if err != nil {
			log.Printf("ERR: cannot read file %v, %v, skip it", srcFile, err)
			continue
		}

		var srcObj interface{}
		err = unmarshal(srcData, &srcObj)
		if err != nil {
			log.Printf("ERR: cannot unmarshal file %v to %s, %v, skip it", srcFile, srcFormat, err)
			continue
		}

		destObj := convert(srcObj)
		destData, err := marshal(destObj)
		if err != nil {
			log.Printf("ERR: cannot marshal fild %v data to %v, %v, skip it", srcFile, destFormat, err)
			continue
		}

		err = ioutil.WriteFile(destFile, destData, 0644)
		if err != nil {
			log.Printf("ERR: cannot write to file %v, %v, skip it", destFile, err)
			continue
		}

		log.Printf("convert %v to %v done", srcFile, destFile)
	}
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

	log.Printf("found %v files", len(files))

	// First, convert all json format files to yaml
	log.Printf("start to convert all json format files to yaml format")
	transform(files, ".json", ".yaml", json.Unmarshal, yaml.Marshal)

	// Second, convert all yaml format files to json
	log.Printf("start to convert all yaml format files to json format")
	transform(files, ".yaml", ".json", yaml.Unmarshal, func(in interface{}) ([]byte, error) {
		return json.MarshalIndent(in, "", "    ")
	})
}
