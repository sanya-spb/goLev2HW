package main

import (
	"fmt"

	"github.com/sanya-spb/goLev2HW/utils/version"
	_ "gopkg.in/yaml.v2"
)

type APP struct {
	Version version.AppVersion
}

var MyApp *APP = new(APP)

func main() {
	MyApp.Version = *version.Version
	fmt.Printf("version: %+v\n", MyApp.Version)
}
