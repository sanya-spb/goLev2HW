package task2_test

import (
	"fmt"

	. "github.com/sanya-spb/goLev2HW/task2"
)

func ExampleDivByMyself() {
	if result, err := DivByMyself(4); err == nil {
		fmt.Println(result)
	}
	// Output:
	// 1
}
