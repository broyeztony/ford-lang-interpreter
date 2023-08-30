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

	if nodeType == "BinaryExpression" {
		return evalBinaryExpression(node)
	}

	if nodeType == "NumericLiteral" {
		return node["value"]
	}

	panic("Not Implemented.")
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

func evalBinaryExpression(expression map[string]interface{}) interface{} {

	operator := expression["operator"]
	left := expression["left"]
	right := expression["right"]

	if operator == "+" {
		leftResult := Eval(left.(map[string]interface{}))
		rightResult := Eval(right.(map[string]interface{}))
		return leftResult.(float64) + rightResult.(float64)
	}

	if operator == "-" {
		leftResult := Eval(left.(map[string]interface{}))
		rightResult := Eval(right.(map[string]interface{}))
		return leftResult.(float64) - rightResult.(float64)
	}

	if operator == "*" {
		leftResult := Eval(left.(map[string]interface{}))
		rightResult := Eval(right.(map[string]interface{}))
		return leftResult.(float64) * rightResult.(float64)
	}

	if operator == "/" {
		leftResult := Eval(left.(map[string]interface{}))
		rightResult := Eval(right.(map[string]interface{}))
		return leftResult.(float64) / rightResult.(float64)
	}

	panic(fmt.Sprintf("evalBinaryExpression unknown operator: %v", operator))
}
