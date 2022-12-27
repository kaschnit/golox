package interpreter

import "fmt"

// Return is implemented as an error so that all execution of AST nodes that are
// parent to a node that used a return statement will propagate the return statement up.
// This allows return propagation without any explicit handling of the return.
// For nodes that must handle a return (e.g., CallExpr), the return can be explicitly handled
// by checking if the error is of type *Return.
type Return struct {
	Value interface{}
}

func NewReturn(value interface{}) *Return {
	return &Return{Value: value}
}

func (r *Return) Error() string {
	return fmt.Sprintf("RETURN %v", r.Value)
}
