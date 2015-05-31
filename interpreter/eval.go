package interpreter
import (
	"fmt"
	"io/ioutil"
)

func Eval(sourcepath string) {
	if output, err := ioutil.ReadFile(sourcepath); err != nil {
		fmt.Println("Error reading file:", sourcepath)
	} else {
		fmt.Println(string(output))
	}
}
