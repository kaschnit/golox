package ast

// Represents an entire program.
// The root of the AST.
type Program struct {
	Statements []Stmt
}

func (p *Program) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitProgram(p)
}
