package environment

// Represents a block's environment.
type Environment struct {
	parent *Environment
	vars   map[string]interface{}
}

// Create a new environment with variables defined.
func NewEnvironment(vars map[string]interface{}) *Environment {
	return &Environment{
		parent: nil,
		vars:   vars,
	}
}

func (e *Environment) Get(varName string) (value interface{}, exists bool) {
	val, exists := e.vars[varName]
	return val, exists
}

func (e *Environment) TraverseGet(varName string) (value interface{}, exists bool) {
	envWithVar := e.findEnvContainingName(varName)
	if envWithVar == nil {
		return nil, false
	}
	return envWithVar.vars[varName], true
}

func (e *Environment) Replace(varName string, val interface{}) (exists bool) {
	envWithVar := e.findEnvContainingName(varName)
	if envWithVar == nil {
		return false
	}
	envWithVar.vars[varName] = val
	return true
}

func (e *Environment) NewChild() *Environment {
	child := NewEnvironment(make(map[string]interface{}))
	child.parent = e
	return child
}

func (e *Environment) WithValue(varName string, value interface{}) *Environment {
	child := e.NewChild()
	child.vars[varName] = value
	return child
}

func (e *Environment) WithValues(values map[string]interface{}) *Environment {
	child := e.NewChild()
	for varName, value := range values {
		child.vars[varName] = value
	}
	return child
}

func (e *Environment) findEnvContainingName(varName string) *Environment {
	currentEnv := e
	for currentEnv != nil {
		if _, exists := currentEnv.vars[varName]; exists {
			break
		}
		currentEnv = currentEnv.parent
	}
	return currentEnv
}
