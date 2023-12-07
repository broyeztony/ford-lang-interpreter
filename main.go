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
	result := interp.Eval(interp.ParseAST(astString), globalEnv)
	fmt.Printf("result: %+v\n", result)
}
