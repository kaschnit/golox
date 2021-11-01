package ast

type Program struct {
	Statements []Stmt
}

func (p *Program) Accept(v AstVisitor) (interface{}, error) {
	return v.VisitProgram(p)
}
