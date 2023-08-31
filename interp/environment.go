package interp

import "fmt"

type Environment struct {
	record map[string]interface{}
	parent *Environment
}

func NewEnvironment(parent *Environment) *Environment {
	env := &Environment{}
	env.record = make(map[string]interface{})
	env.parent = parent
	return env
}

func (e *Environment) define(name string, value interface{}) interface{} {
	e.record[name] = value
	fmt.Println("@Environment::define name", name, "value:", value)
	return value
}

func (e *Environment) lookup(name string) interface{} {
	value, ok := e.record[name]
	if !ok {
		panic(fmt.Sprintf("ReferenceError. Variable '%v' is not defined.", name))
	}
	return value
}
