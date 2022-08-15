package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ankaa-labs/ankaa/test"
	"gopkg.in/yaml.v3"
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
	for _, file := range files {
		if !strings.HasSuffix(file, ".json") {
			log.Printf("%s is not json format, skip it", file)
			continue
		}

		yamlFile := file[0:len(file)-len(".json")] + ".yaml"
		if _, err := os.Stat(yamlFile); err == nil {
			log.Printf("ERR: the target yaml file %v exists, skip it", yamlFile)
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			log.Printf("ERR: stat target yaml file %v, %v, skip it", yamlFile, err)
			continue
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("ERR: cannot read file %v, %v, skip it", file, err)
			continue
		}

		var obj interface{}
		err = json.Unmarshal(data, &obj)
		if err != nil {
			log.Printf("ERR: cannot unmarshal file %v to json, %v, skip it", file, err)
			continue
		}

		obj1 := convert(obj)
		data1, err := yaml.Marshal(obj1)
		if err != nil {
			log.Printf("ERR: cannot marshal fild %v data to yaml, %v, skip it", file, err)
			continue
		}

		err = ioutil.WriteFile(yamlFile, data1, fs.ModeAppend)
		if err != nil {
			log.Printf("ERR: cannot write to file %v, %v, skip it", yamlFile, err)
			continue
		}

		log.Printf("convert json %v to yaml %v done", file, yamlFile)
	}
}
