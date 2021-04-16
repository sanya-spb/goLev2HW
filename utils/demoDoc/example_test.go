package demoDoc_test

import (
	"fmt"

	. "github.com/sanya-spb/goLev2HW/demoDoc"
)

func ExampleDivByMyself() {
	if result, err := DivByMyself(4); err == nil {
		fmt.Println(result)
	}
	// Output:
	// 1
}
