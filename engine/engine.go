package engine
import (
	"fmt"
	"io/ioutil"
)

type Engine struct {

}

func NewEngine() *Engine {
	engine := new(Engine)
	return engine
}

func (engine *Engine) Run(sourcepath string) {
	if output, err := ioutil.ReadFile(sourcepath); err != nil {
		fmt.Println("Error reading file:", sourcepath)
	} else {
		fmt.Println(string(output))
	}
}