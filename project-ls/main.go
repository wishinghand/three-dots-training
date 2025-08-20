package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	fmt.Print(strings.Join(listFiles("./testdata"), " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
