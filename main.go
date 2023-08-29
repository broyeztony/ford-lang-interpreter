package main

import (
	"encoding/json"
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
          "value": 10
        },
        "operator": "+",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 4
          },
          "operator": "-",
          "right": {
            "type": "NumericLiteral",
            "value": 2
          },
          "type": "BinaryExpression"
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	var ast map[string]interface{}
	err := json.Unmarshal([]byte(input), &ast)
	if err != nil {
		fmt.Println("Error parsing program AST:", err)
		return
	}

	result := interp.Eval(ast)
	fmt.Printf("result: %+v\n", result)
}
