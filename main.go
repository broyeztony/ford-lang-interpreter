package main

import (
	"fmt"
	interp "ford-lang-interpreter/interp"
)

func main() {

	input := `{
  "body": [
    {
      "expression": {
        "type": "StringLiteral",
        "value": "hello"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	result := interp.Eval(interp.ParseAST(input))
	fmt.Printf("result: %+v\n", result)
}
