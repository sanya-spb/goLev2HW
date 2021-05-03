package main

import (
	"fmt"

	"github.com/sanya-spb/goLev2HW/srs_01/structurer"
)

type confDatabase struct {
	Host string
	Port int
	User string
	Pass string
	Ssl  bool
}

func main() {
	conf := &confDatabase{}
	filler := map[string]interface{}{
		"Host": "192.168.1.1",
		"Port": 8181,
		"User": "test",
		"Pass": "pwd45678",
		"Ssl":  true,
	}

	fmt.Printf("Result before: %+v\n", conf)
	err := structurer.FillStruct(conf, filler)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	fmt.Printf("Result after: %+v\n", conf)

}
