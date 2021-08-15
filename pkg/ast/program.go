package ast

type Program struct {
	Statements []Stmt
}

func (p *Program) Accept(v AstVisitor) interface{} {
	return v.VisitProgram(p)
}
