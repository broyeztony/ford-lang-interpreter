package main

import (
	"fmt"
	interp "ford-lang-interpreter/interp"
	"io/ioutil"
)

func main() {

	ast, err := ioutil.ReadFile("ast.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	astString := string(ast)

	globalEnv := interp.NewEnvironment(nil)

	// pre-install `print` as a native function using `fmt.Println` as the backend
	globalEnv.Define("print", func() func(a ...any) (n int, err error) {
		return fmt.Println
	})

	result := interp.Eval(interp.ParseAST(astString), globalEnv)
	fmt.Printf("result: %+v\n", result)
}
