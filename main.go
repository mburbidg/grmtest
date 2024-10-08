package main

import (
	"flag"
	"fmt"
	"github.com/mburbidg/grmtest/parser"
	"io"
	"os"
	"path/filepath"
)

var success, fail int

func main() {
	filename := flag.String("file", "", "Filename to read GQL-program from")
	dirname := flag.String("dir", "", "Directory name containing GQL-programs")
	flag.Parse()
	if *filename != "" {
		parseFile(*filename)
	}
	if *dirname != "" {
		parseDir(*dirname)
	}
	fmt.Printf("Success: %d, Fail: %d\n", success, fail)
}

func parseFile(filename string) {
	fail++
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open '%s': %v\n", filename, err)
		return
	}
	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read '%s': %v\n", filename, err)
		return
	}
	err = parser.Parse(string(b))
	if err != nil {
		fmt.Printf("Syntax error in '%s': %v\n", filename, err)
		return
	}
	fail--
	success++
	fmt.Printf("Success: %s\n", filename)
}

func parseDir(dirname string) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Printf("Failed to read directory '%s': %v\n", dirname, err)
		return
	}

	for _, file := range files {
		path := filepath.Join(dirname, file.Name())
		if !file.IsDir() {
			if filepath.Ext(path) == ".gql" {
				parseFile(path)
			}
		} else {
			parseDir(path)
		}
	}
}
