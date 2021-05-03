package main

import (
	"errors"
	"fmt"
	"reflect"
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
	err := FillStruct(conf, filler)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	fmt.Printf("Result after: %+v\n", conf)

}

func FillStruct(in interface{}, filler map[string]interface{}) error {
	if in == nil {
		return errors.New("arg is nil")
	}

	if filler == nil {
		return errors.New("filler is nil")
	}

	sVal := reflect.ValueOf(in)
	if sVal.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
	}

	if sVal.Type().Kind() != reflect.Struct {
		return fmt.Errorf("arg is not struct; kind=%s", sVal.Type().Kind())
	}

	for key, v := range filler {
		value := sVal.FieldByName(key)
		newVal := reflect.ValueOf(v)
		if value.Type().Kind() != newVal.Type().Kind() {
			return errors.New("type field is not correct")
		}
		value.Set(newVal)
	}
	return nil
}
