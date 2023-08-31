package main

import (
	"fmt"
	interp "ford-lang-interpreter/interp"
)

func main() {

	input := `{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "x",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "declarations": [
        {
          "id": {
            "name": "y",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "declarations": [
        {
          "id": {
            "name": "z",
            "type": "Identifier"
          },
          "initializer": {
            "left": {
              "name": "x",
              "type": "Identifier"
            },
            "operator": "+",
            "right": {
              "name": "y",
              "type": "Identifier"
            },
            "type": "BinaryExpression"
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "expression": {
        "name": "z",
        "type": "Identifier"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	globalEnv := interp.NewEnvironment(nil)
	result := interp.Eval(interp.ParseAST(input), globalEnv)
	fmt.Printf("result: %+v\n", result)
}
