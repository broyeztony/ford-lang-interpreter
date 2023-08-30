package interp

import (
	"encoding/json"
	"fmt"
)

func ParseAST(input string) map[string]interface{} {
	var ast map[string]interface{}
	err := json.Unmarshal([]byte(input), &ast)
	if err != nil {
		fmt.Println("Error parsing program AST:", err)
		return nil
	}
	return ast
}
