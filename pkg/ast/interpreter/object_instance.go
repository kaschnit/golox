package interpreter

import (
	"fmt"

	loxerr "github.com/kaschnit/golox/pkg/errors"
	"github.com/kaschnit/golox/pkg/token"
)

type LoxClassInstance struct {
	Class      *LoxClass
	properties map[string]interface{}
}

func NewLoxClassInstance(cls *LoxClass) *LoxClassInstance {
	return &LoxClassInstance{
		Class:      cls,
		properties: make(map[string]interface{}),
	}
}

func (c *LoxClassInstance) GetProperty(propertyName *token.Token) (interface{}, error) {
	if prop, ok := c.properties[propertyName.Lexeme]; ok {
		return prop, nil
	}

	if prop, ok := c.Class.methods[propertyName.Lexeme]; ok {
		return prop.Bind(c), nil
	}

	return nil, loxerr.Runtime(propertyName, fmt.Sprintf("Property '%s' is not defined on %s", propertyName.Lexeme, c))
}

func (c *LoxClassInstance) SetProperty(propertyName *token.Token, value interface{}) {
	c.properties[propertyName.Lexeme] = value
}

func (c *LoxClassInstance) String() string {
	return fmt.Sprintf("<instance of %s [%p]>", c.Class.declaration.Name.Lexeme, c)
}
