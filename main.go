package main
import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) == 1) {
		fmt.Println("Sprout v0.1")
	} else {
		fmt.Println(os.Args[1])
	}
}
