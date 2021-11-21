package environment

// Represents a block's environment.
type Environment struct {
	parent *Environment
	vars   map[string]interface{}
}

// Create a new Environment.
func NewEnvironment(parent *Environment) *Environment {
	return NewEnvironmentWithVars(parent, make(map[string]interface{}))
}

// Create a new environment with the variables defined.
func NewEnvironmentWithVars(parent *Environment, vars map[string]interface{}) *Environment {
	return &Environment{parent, vars}
}

// Find the variable with name varName by traversing the environments upwards.
func (e *Environment) GetTraverse(varName string) (interface{}, bool) {
	for currentEnv := e; currentEnv != nil; currentEnv = currentEnv.parent {
		if value, exists := currentEnv.Get(varName); exists {
			return value, true
		}
	}
	return nil, false
}

func (e *Environment) Get(varName string) (interface{}, bool) {
	val, exists := e.vars[varName]
	return val, exists
}

func (e *Environment) GetAt(varName string, distance int) interface{} {
	currentEnv := e
	for i := 0; i < distance; i++ {
		currentEnv = currentEnv.parent
	}
	return currentEnv.vars[varName]
}

// Update the existing variable with name varName by traversing the environments upwards.
func (e *Environment) ReplaceTraverse(varName string, val interface{}) bool {
	for currentEnv := e; currentEnv != nil; currentEnv = currentEnv.parent {
		if _, exists := currentEnv.Get(varName); exists {
			currentEnv.Set(varName, val)
			return true
		}
	}
	return false
}

func (e *Environment) Set(varName string, val interface{}) {
	e.vars[varName] = val
}

func (e *Environment) Fork() *Environment {
	return NewEnvironment(e)
}
