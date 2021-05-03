package structurer

import (
	"errors"
	"fmt"
	"reflect"
)

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
