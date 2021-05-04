package main

import (
	"errors"
	"testing"
)

func TestFillStruct(t *testing.T) {

	type tCase struct {
		p1   confDatabase
		p2   map[string]interface{}
		want error
	}
	type tCaseErr struct {
		p1   interface{}
		p2   map[string]interface{}
		want error
	}

	testCaseNoErr := []tCase{
		{
			confDatabase{},
			map[string]interface{}{
				"Host": "192.168.1.1",
				"Port": 8183,
				"User": "test",
				"Pass": "pwd45678",
				"Ssl":  true,
			},
			nil,
		},
	}

	testCaseErr := []tCaseErr{
		{
			nil,
			map[string]interface{}{},
			errors.New("arg is not struct"),
		},
		{
			confDatabase{},
			nil,
			errors.New("filler is nil"),
		},
		{
			confDatabase{},
			map[string]interface{}{
				"Host": "192.168.1.1",
				"Port": "8182",
				"User": "test",
				"Pass": "pwd45678",
				"Ssl":  true,
			},
			errors.New("type field is not correct"),
		},
	}

	for _, test := range testCaseNoErr {
		if result := FillStruct(&test.p1, test.p2); result != nil {
			t.Errorf("Invalid value for param2: %v, got: %v", test.p2, result)
		}
	}
	for _, test := range testCaseErr {
		if result := FillStruct(&test.p1, test.p2); result == nil {
			t.Errorf("Invalid value for param2: %v, got: %v, want: %s", test.p2, result, test.want.Error())
		}
	}
}
