package tests

import (
	"ford-lang-interpreter/interp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	input := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "+",
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

	actual := interp.Eval(interp.ParseAST(input))
	expected := float64(3)

	assert.Equal(t, expected, actual)
}

func TestSubtraction(t *testing.T) {

	input := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 10
        },
        "operator": "-",
        "right": {
          "type": "NumericLiteral",
          "value": 1.15
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`
	actual := interp.Eval(interp.ParseAST(input))
	expected := float64(8.85)

	assert.Equal(t, expected, actual)
}

func TestMultiplication(t *testing.T) {

	input := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 8.2
        },
        "operator": "*",
        "right": {
          "type": "NumericLiteral",
          "value": 4.1
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`
	actual := interp.Eval(interp.ParseAST(input))
	expected := float64(33.62)

	assert.Equal(t, expected, actual)
}

func TestDivision(t *testing.T) {

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

	actual := interp.Eval(interp.ParseAST(input))
	expected := float64(2)

	assert.Equal(t, expected, actual)
}

// 3 * (2 + 1) / (5.12 - 1.006)
func TestComplexExpression(t *testing.T) {

	input := `{
  "body": [
    {
      "expression": {
        "left": {
          "left": {
            "type": "NumericLiteral",
            "value": 3
          },
          "operator": "*",
          "right": {
            "left": {
              "type": "NumericLiteral",
              "value": 2
            },
            "operator": "+",
            "right": {
              "type": "NumericLiteral",
              "value": 1
            },
            "type": "BinaryExpression"
          },
          "type": "BinaryExpression"
        },
        "operator": "/",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 5.12
          },
          "operator": "-",
          "right": {
            "type": "NumericLiteral",
            "value": 1.006
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

	actual := interp.Eval(interp.ParseAST(input))
	expected := 2.1876519202722413

	assert.Equal(t, expected, actual)
}
