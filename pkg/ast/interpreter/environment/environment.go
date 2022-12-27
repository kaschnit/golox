package environment

import "fmt"

// Represents a block's environment.
type Environment struct {
	parent *Environment
	vars   map[string]interface{}
}

// Create a new Environment.
func NewEnvironment(parent *Environment) *Environment {
	return NewEnvironmentWithVars(parent, make(map[string]interface{}))
}

// Create a new environment with variables defined.
func NewEnvironmentWithVars(parent *Environment, vars map[string]interface{}) *Environment {
	return &Environment{parent, vars}
}

func (e *Environment) Get(varName string) (interface{}, bool) {
	val, exists := e.vars[varName]
	return val, exists
}

func (e *Environment) GetAt(varName string, distance int) interface{} {
	return e.ancestor(distance).vars[varName]
}

func (e *Environment) Replace(varName string, val interface{}) bool {
	if _, exists := e.Get(varName); exists {
		e.Set(varName, val)
		return true
	}
	return false
}

func (e *Environment) Set(varName string, val interface{}) {
	e.vars[varName] = val
}

func (e *Environment) SetAt(varName string, distance int, val interface{}) {
	e.ancestor(distance).vars[varName] = val
}

func (e *Environment) NewChild() *Environment {
	return NewEnvironment(e)
}

func (e *Environment) Print() {
	if e.parent != nil {
		e.parent.Print()
	}
	fmt.Printf(" --> %v (%p)", e.vars, e.vars)
}

func (e *Environment) ancestor(distance int) *Environment {
	currentEnv := e
	for i := 0; i < distance; i++ {
		currentEnv = currentEnv.parent
	}
	return currentEnv
}
