package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func printFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	// fmt.Println(path)
	finfo, err := os.Stat(path)
	if err != nil && !finfo.IsDir() {
		// might be symlink
		if finfo.Size() < 50 {
			fmt.Printf("%s might be a symlink\n", path)
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	dir := os.Args[1]
	err := filepath.Walk(dir, printFile)
	if err != nil {
		log.Fatal(err)
	}
}