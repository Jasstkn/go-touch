package main

import (
	"fmt"
	"os"
	"strings"
)

const Usage = `Usage:
%s /new/file.txt
`

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(Usage, os.Args[0])
		os.Exit(-1)
	}

	path := os.Args[1]
	dirs := path[:strings.LastIndex(path, "/")]
	os.MkdirAll(dirs, 0755)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Create(path)
	}
}
