package main

import (
	"fmt"

	. "github.com/sanya-spb/goLev2HW/task2"
)

func main() {
	for i := 3; i >= -3; i-- {
		if result, err := DivByMyself(i); err != nil {
			fmt.Printf("Error on: %[1]d / %[1]d\n%[2]s\n", i, err)
		} else {
			fmt.Printf("divByMyself: %[1]d / %[1]d = %[2]d\n", i, result)
		}
	}
}
