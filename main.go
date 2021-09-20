package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Param struct {
	Name  string `yaml:"name"`
	Value int64  `yaml:"value"`
}

type config struct {
	Currencies []Param `yaml:"currencies"`
}

func main() {
	yamlFile := yamlParse()
	fmt.Printf("%v", yamlFile)
}

func yamlParse() config {
	var c config
	data, err := ioutil.ReadFile("data_file.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}
