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
        "left": {
          "type": "NumericLiteral",
          "value": 4
        },
        "operator": "+",
        "right": {
          "type": "NumericLiteral",
          "value": 3
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    },
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "-",
        "right": {
          "type": "NumericLiteral",
          "value": 1
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	result := interp.Eval(interp.ParseAST(input))
	fmt.Printf("result: %+v\n", result)
}
