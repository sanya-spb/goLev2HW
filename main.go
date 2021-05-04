package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sanya-spb/goLev2HW/utils/config"
	"github.com/sanya-spb/goLev2HW/utils/fdouble"
	"github.com/sanya-spb/goLev2HW/utils/version"
)

type APP struct {
	Version version.AppVersion
	Config  config.Config
}

var MyApp *APP = new(APP)

func main() {
	MyApp.Version = *version.Version
	MyApp.Config = *config.NewConfig()
	log.Printf("version: %+v\n", MyApp.Version)
	log.Printf("config: %+v\n", MyApp.Config)

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(fdouble.Done)
	}()

	// Traverse each root of the file tree in parallel.
	fileHashes := make(chan fdouble.FDescr)
	var n sync.WaitGroup
	for _, root := range MyApp.Config.Dirs {
		n.Add(1)
		go fdouble.ScanDir(root, &n, fileHashes)
	}
	go func() {
		n.Wait()
		close(fileHashes)
	}()

loop:
	for {
		select {
		case <-fdouble.Done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileHashes {
				// Do nothing.
			}
			return
		case fHash, ok := <-fileHashes:
			if !ok {
				break loop // fileHashes was closed
			}
			fmt.Fprintf(os.Stdout, "%+v\n", fHash)
		}
	}
}
