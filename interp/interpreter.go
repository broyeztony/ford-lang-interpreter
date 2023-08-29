package interp

import "fmt"

type Interpreter struct {
}

func Eval(node map[string]interface{}) interface{} {

	fmt.Printf("Eval: %+v\n", node["type"])

	return nil
}
