package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Sprout version 0.1")
		fmt.Println("Usage: sprout sourcepath")
	} else {
		fmt.Println(os.Args[1])
	}
}
