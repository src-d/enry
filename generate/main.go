package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

const (
	languagesYAML = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
)

func main() {
	res, err := http.Get(languagesYAML)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var out yaml.MapSlice
	if err := yaml.Unmarshal(buf, &out); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", out)
}
