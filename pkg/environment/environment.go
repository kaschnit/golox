package environment

// Represents a block's environment.
type Environment struct {
	parent *Environment
	vars   map[string]interface{}
}

// Create a new Scope.
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent: parent,
		vars:   make(map[string]interface{}),
	}
}

// Find the variable with name varName by traversing the scope upwards.
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

// Update the existing variable with name varName by traversing the scope upwards.
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
