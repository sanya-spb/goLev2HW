package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"

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
	db, err := sql.Open("sqlite3", ":memory:")
	// db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`
		CREATE TABLE "file_hash" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT, 
			"path" text not null, 
			"hash" text NOT NULL, 
			"size" integer NOT NULL);
	`)
	if err != nil {
		log.Fatal(err)
	}
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
		go fdouble.ScanDir(strings.TrimRight(root, string(filepath.Separator)), &n, fileHashes)
	}
	go func() {
		n.Wait()
		close(fileHashes)
	}()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`insert into file_hash(path, hash, size) values(?, ?, ?);`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
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
			_, err = stmt.Exec(fHash.Path(), fHash.Hash(), fHash.Size())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	tx.Commit()

	rows, err := db.Query(`
		select file_hash.path, file_hash.hash, file_hash.size
		from file_hash
		join (
			select 
				hash, size
			from file_hash
			group by hash, size
			having 
				count(id) > 1
				and count(id) > ?
			)tt on 
			tt.hash=file_hash.hash 
			and tt.size=file_hash.size
		order by 2, 3;
	`, MyApp.Config.DFactor)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			path, hash string
			size       int
		)
		if err := rows.Scan(&path, &hash, &size); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n  hash: %s\n  size: %d\n---\n", path, hash, size)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
