package tests

import (
	"encoding/json"
	"fmt"
	"ford-lang-interpreter/interp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleBinaryExpression(t *testing.T) {

	input := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 8
        },
        "operator": "/",
        "right": {
          "type": "NumericLiteral",
          "value": 4
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

	actual := interp.Eval(ast)
	expected := float64(2)

	assert.Equal(t, expected, actual)
}
