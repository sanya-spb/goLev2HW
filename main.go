package main

import (
	"fmt"

	. "github.com/sanya-spb/goLev2HW/utils/demoDoc"
	"github.com/sanya-spb/goLev2HW/utils/version"
)

type APP struct {
	Version version.AppVersion
}

var MyApp *APP = new(APP)

func main() {
	MyApp.Version = *version.Version
	fmt.Printf("version: %+v\n", MyApp.Version)

	for i := 1; i >= -1; i-- {
		if result, err := DivByMyself(i); err != nil {
			fmt.Printf("Error on: %[1]d / %[1]d\n%[2]s\n", i, err)
		} else {
			fmt.Printf("divByMyself: %[1]d / %[1]d = %[2]d\n", i, result)
		}
	}
}
