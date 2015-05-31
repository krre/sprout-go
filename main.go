package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krre/sprout/engine"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Sprout version 0.1")
		fmt.Println("Usage: sprout sourcepath")
	} else if sourcepath, err := filepath.Abs(os.Args[1]); err != nil {
		fmt.Println("Error command line parse")
	} else {
		e := engine.NewEngine()
		e.Run(sourcepath)
	}
}
