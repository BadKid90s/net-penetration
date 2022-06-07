package test

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net-penetration/conf"
	"testing"
)

func TestUnMarshal(t *testing.T) {
	s := new(conf.Server)
	b, err := ioutil.ReadFile("../conf/server.yml")
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
}
