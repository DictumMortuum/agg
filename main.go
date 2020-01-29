package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
)

func isYml(env string) bool {
	re1 := regexp.MustCompile(`yml$`)
	re2 := regexp.MustCompile(`yaml$`)

	if re1.MatchString(env) || re2.MatchString(env) {
		return true
	} else {
		return false
	}
}

func main() {
	m := make(map[interface{}]interface{})

	files, _ := ioutil.ReadDir(".")

	for _, file := range files {
		name := file.Name()

		if isYml(name) {
			data, _ := ioutil.ReadFile(name)
			yaml.Unmarshal(data, &m)
		}
	}

	d, _ := yaml.Marshal(&m)
	fmt.Printf("%s", string(d))
}
