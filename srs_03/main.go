package main

import (
	"fmt"
	"os"
)

const (
	dir    = "./"
	prefix = "out_"
	count  = 1000
)

func main() {
	for i := 0; i < count; i++ {
		func() {
			fileName := fmt.Sprintf(prefix+"%06d", i)
			if fH, err := os.Create(dir + fileName); err != nil {
				fmt.Printf("Error: %v on create %s\n", err, fileName)
			} else {
				defer fH.Close()
				fmt.Printf("%s\n", fileName)
				fH.WriteString(fileName)
			}
		}()
	}
}
