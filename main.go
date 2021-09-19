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

type conf struct {
	Currencies []Param `yaml:"currencies"`
}

func main() {
	var c conf
	data, err := ioutil.ReadFile("data_file.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
