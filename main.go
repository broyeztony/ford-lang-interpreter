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
        "type": "NumericLiteral",
        "value": 0
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
	fmt.Printf("%+v\n", result)
}
