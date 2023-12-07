package interp

import "fmt"

func Eval(node map[string]interface{}, env *Environment) interface{} {

	fmt.Println("Eval node", node)

	nodeType := node["type"]
	fmt.Printf("Eval type: %+v\n", nodeType)

	if nodeType == "Program" {
		return evalBlock(node["body"].([]interface{}), env)
	}

	if nodeType == "BlockStatement" {
		return evalBlock(node["body"].([]interface{}), env)
	}

	if nodeType == "ExpressionStatement" {
		return evalExpressionStatement(node, env)
	}
	if nodeType == "IfStatement" {
		return evalIfStatement(node, env)
	}

	if nodeType == "BinaryExpression" { // +, -, * ...
		return evalBinaryExpression(node, env)
	}

	if nodeType == "BinaryExpression" {
		return evalBinaryExpression(node, env)
	}

	if nodeType == "NumericLiteral" {
		return node["value"]
	}
	if nodeType == "StringLiteral" {
		return node["value"]
	}
	if nodeType == "BooleanLiteral" {
		return node["value"]
	}

	if nodeType == "VariableStatement" {
		return evalVariableStatement(node["declarations"].([]interface{}), env)
	}
	if nodeType == "VariableDeclaration" {
		return evalVariableDeclaration(node, env)
	}
	if nodeType == "AssignmentExpression" {
		return evalAssignmentExpression(node, env)
	}

	// variable access
	if isVariableAccess(node) {
		return env.lookup(node["name"].(string))
	}

	panic(fmt.Sprintf("Not Implemented: %v", node["type"]))
}

func evalBlock(block []interface{}, env *Environment) interface{} {

	var result interface{}

	for _, stmtNode := range block {
		// fmt.Println("evalBlock index:", index, "stmtNode:", stmtNode, stmtNode.(map[string]interface{})["type"])
		result = Eval(stmtNode.(map[string]interface{}), env)
	}

	return result
}

func evalExpressionStatement(stmt map[string]interface{}, env *Environment) interface{} {

	// fmt.Println("evalExpressionStatement stmt:", stmt)
	return Eval(stmt["expression"].(map[string]interface{}), env)
}

func evalIfStatement(ifStmt map[string]interface{}, env *Environment) interface{} {
	test := ifStmt["test"].(map[string]interface{})
	consequent := ifStmt["consequent"].(map[string]interface{})
	alternate := ifStmt["alternate"].(map[string]interface{})

	if Eval(test, env).(bool) {
		return Eval(consequent, env)
	}
	return Eval(alternate, env)
}

func evalBinaryExpression(expression map[string]interface{}, env *Environment) interface{} {

	operator := expression["operator"]
	left := expression["left"]
	right := expression["right"]

	leftResult := Eval(left.(map[string]interface{}), env)
	rightResult := Eval(right.(map[string]interface{}), env)

	if operator == "+" {
		return leftResult.(float64) + rightResult.(float64)
	}

	if operator == "-" {
		return leftResult.(float64) - rightResult.(float64)
	}

	if operator == "*" {
		return leftResult.(float64) * rightResult.(float64)
	}

	if operator == "/" {
		return leftResult.(float64) / rightResult.(float64)
	}

	// Logical operator
	if operator == ">" {
		return leftResult.(float64) > rightResult.(float64)
	}
	if operator == ">=" {
		return leftResult.(float64) >= rightResult.(float64)
	}
	if operator == "<" {
		return leftResult.(float64) < rightResult.(float64)
	}
	if operator == "<=" {
		return leftResult.(float64) <= rightResult.(float64)
	}
	if operator == "==" {
		return leftResult.(float64) == rightResult.(float64)
	}

	panic(fmt.Sprintf("evalBinaryExpression unknown operator: %v", operator))
}

func evalVariableStatement(declarations []interface{}, env *Environment) interface{} {
	var result interface{}
	for _, declaration := range declarations {
		// fmt.Println("evalBlock index:", index, "stmtNode:", stmtNode, stmtNode.(map[string]interface{})["type"])
		result = Eval(declaration.(map[string]interface{}), env)
	}
	return result
}

func evalVariableDeclaration(varDeclaration map[string]interface{}, env *Environment) interface{} {

	variableName := varDeclaration["id"].(map[string]interface{})["name"].(string)
	var variableValue interface{}

	if varDeclaration["initializer"] != nil {
		variableValue = Eval(varDeclaration["initializer"].(map[string]interface{}), env)
	}

	return env.define(variableName, variableValue)
}

func evalAssignmentExpression(assignmentExpr map[string]interface{}, env *Environment) interface{} {

	fmt.Println("@ evalAssignmentExpression: ", assignmentExpr)

	left := assignmentExpr["left"].(map[string]interface{})
	right := assignmentExpr["right"].(map[string]interface{})
	// we currently only treat simple assignment (i.e `=`)
	// TODO: handle complex assignment (`+=`, `-=`, ...)
	// operator := assignmentExpr["operator"].(string)

	// TODO: assignment to property
	if left["type"] == "MemberExpression" {
		panic("Not implemented!")
	}

	// simple assignment
	if left["type"] == "Identifier" {
		return env.assign(left["name"].(string), Eval(right, env))
	}
	panic(fmt.Sprintf("AssignmentExpression can only be performed on Identifier node or MemberExpression node but got %v", left["type"]))
}

func isVariableAccess(node map[string]interface{}) bool {
	return node["type"] == "Identifier"
}
