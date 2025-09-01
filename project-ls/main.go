package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Your solution goes here. Good luck!
	for _, file := range listFiles("./testdata") {
		fmt.Println(file)
	}

	// fmt.Print(strings.Join(listFiles("./testdata"), " "))
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
