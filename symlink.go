package main

import (
	"fmt"
	"os"
)

func main() {
	from := "symlink.go"
	to := "thisisasymlink.link"
	err := os.Symlink(from, to)
	if (err != nil) {
		fmt.Printf("Failed to symlink: %v\n", err)
	}
}

