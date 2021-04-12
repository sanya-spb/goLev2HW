package task2

import (
	"testing"
)

func TestDivByMyself(t *testing.T) {
	type testCase struct {
		value  int
		result int
	}
	var tests = []testCase{
		{
			value:  1,
			result: 1,
		},
	}

	for _, test := range tests {
		if result, err := DivByMyself(test.value); !(err == nil && result == test.result) {
			t.Errorf("For %d expected %d, got %d", test.value, test.result, result)
		}
	}

	if _, err := DivByMyself(0); err == nil {
		t.Errorf("For %d expected error, got %v", 0, err)
	}

}
