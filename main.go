package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Sprout version 0.1")
		fmt.Println("Usage: sprout sourcepath")
	} else if sourcepath, err := filepath.Abs(os.Args[1]); err != nil {
		fmt.Println("Error command line parse")
	} else if output, err := ioutil.ReadFile(sourcepath); err != nil {
		fmt.Println("Error reading file:", sourcepath)
	} else {
		fmt.Println(string(output))
	}
}
