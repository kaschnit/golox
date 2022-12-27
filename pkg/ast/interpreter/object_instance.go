package interpreter

import "fmt"

type LoxClassInstance struct {
	Class *LoxClass
}

func NewLoxClassInstance(cls *LoxClass) *LoxClassInstance {
	return &LoxClassInstance{Class: cls}
}

func (c *LoxClassInstance) String() string {
	return fmt.Sprintf("<instance of %s [%p]>", c.Class.declaration.Name.Lexeme, c)
}
