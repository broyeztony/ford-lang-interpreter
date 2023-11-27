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
            "name": "a",
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
            "name": "b",
            "type": "Identifier"
          },
          "initializer": {
            "type": "NumericLiteral",
            "value": 0
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    },
    {
      "alternate": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "b",
                "type": "Identifier"
              },
              "operator": "=",
              "right": {
                "type": "NumericLiteral",
                "value": 30
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "consequent": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "b",
                "type": "Identifier"
              },
              "operator": "=",
              "right": {
                "type": "NumericLiteral",
                "value": 20
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "test": {
        "left": {
          "name": "a",
          "type": "Identifier"
        },
        "operator": "<",
        "right": {
          "type": "NumericLiteral",
          "value": 10
        },
        "type": "BinaryExpression"
      },
      "type": "IfStatement"
    },
    {
      "expression": {
        "name": "b",
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
