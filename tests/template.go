package tests

import (
	"encoding/json"
	"fmt"
	"ford-lang-interpreter/interp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	input := ``

	var ast map[string]interface{}
	err := json.Unmarshal([]byte(input), &ast)
	if err != nil {
		fmt.Println("Error parsing program AST:", err)
		return
	}

	actual := interp.Eval(ast)
	expected := "<>"

	assert.Equal(t, expected, actual)
}
