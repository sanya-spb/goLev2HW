package main

import "testing"

func TestFillStruct(t *testing.T) {
	testCase := []struct {
		n    uint8
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, test := range testCase {
		if result := FillStruct(test.n); result != test.want {
			t.Errorf("Invalid value for N: %d, got: %d, want: %d", test.n, result, test.want)
		}
	}
}
