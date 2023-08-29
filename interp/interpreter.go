package interp

import "fmt"

type Interpreter struct {
}

func Eval(node map[string]interface{}) interface{} {

	fmt.Println("Eval node", node)

	nodeType := node["type"]
	fmt.Printf("Eval type: %+v\n", nodeType)

	if nodeType == "Program" {
		return evalBlock(node["body"].([]interface{}))
	}

	if nodeType == "ExpressionStatement" {
		return evalExpressionStatement(node)
	}

	if nodeType == "NumericLiteral" {
		return node["value"]
	}

	return nil
}

func evalBlock(block []interface{}) interface{} {

	var result interface{}

	for index, stmtNode := range block {
		fmt.Println("evalBlock index:", index, "stmtNode:", stmtNode, stmtNode.(map[string]interface{})["type"])
		result = Eval(stmtNode.(map[string]interface{}))
	}

	return result
}

func evalExpressionStatement(stmt map[string]interface{}) interface{} {

	fmt.Println("evalExpressionStatement stmt:", stmt)
	return Eval(stmt["expression"].(map[string]interface{}))
}
