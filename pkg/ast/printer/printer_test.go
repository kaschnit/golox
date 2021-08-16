package printer

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/kaschnit/golox/pkg/ast"
	"github.com/kaschnit/golox/pkg/token"
	"github.com/kaschnit/golox/pkg/token/tokentype"
	"github.com/stretchr/testify/assert"
)

var bangToken = token.Token{
	Type:    tokentype.BANG,
	Lexeme:  "!",
	Literal: nil,
	Line:    1,
}
var minusToken = token.Token{
	Type:    tokentype.MINUS,
	Lexeme:  "-",
	Literal: nil,
	Line:    1,
}
var equalToken = token.Token{
	Type:    tokentype.EQUAL,
	Lexeme:  "=",
	Literal: nil,
	Line:    1,
}
var starToken = token.Token{
	Type:    tokentype.STAR,
	Lexeme:  "*",
	Literal: nil,
	Line:    1,
}
var plusToken = token.Token{
	Type:    tokentype.PLUS,
	Lexeme:  "+",
	Literal: nil,
	Line:    1,
}

func verifyPrintedToStdout(t *testing.T, expected string, testCode func()) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the code that does the printing
	testCode()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	// Check that the epected value was printed to stdout
	assert.Equal(t, expected, string(out))
}

func TestAstPrinter_Literal(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, "123", func() {
		literalExpr := ast.LiteralExpr{Value: 123}
		literalExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `"hello"`, func() {
		literalExpr := ast.LiteralExpr{Value: "hello"}
		literalExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, "nil", func() {
		literalExpr := ast.LiteralExpr{Value: nil}
		literalExpr.Accept(printer)
	})
}

func TestAstPrinter_Grouped(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, "(group 123)", func() {
		literalExpr := ast.LiteralExpr{Value: 123}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		groupExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(group "hello")`, func() {
		literalExpr := ast.LiteralExpr{Value: "hello"}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		groupExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, "(group nil)", func() {
		literalExpr := ast.LiteralExpr{Value: nil}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		groupExpr.Accept(printer)
	})
}

func TestAstPrinter_GroupsOfGroups(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, `(group (group "hello"))`, func() {
		literalExpr := ast.LiteralExpr{Value: "hello"}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		groupGroupExpr := ast.GroupingExpr{Expression: &groupExpr}
		groupGroupExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(group (group (group (group (group "hello")))))`, func() {
		literalExpr := ast.LiteralExpr{Value: "hello"}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		ggroupExpr := ast.GroupingExpr{Expression: &groupExpr}
		gggroupExpr := ast.GroupingExpr{Expression: &ggroupExpr}
		ggggroupExpr := ast.GroupingExpr{Expression: &gggroupExpr}
		gggggroupExpr := ast.GroupingExpr{Expression: &ggggroupExpr}
		gggggroupExpr.Accept(printer)
	})
}

func TestAstPrinter_Unary(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, "(! 123)", func() {
		literalExpr := ast.LiteralExpr{Value: 123}
		unaryExpr := ast.UnaryExpr{Operator: &bangToken, Right: &literalExpr}
		unaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(- "abc")`, func() {
		literalExpr := ast.LiteralExpr{Value: "abc"}
		unaryExpr := ast.UnaryExpr{Operator: &minusToken, Right: &literalExpr}
		unaryExpr.Accept(printer)
	})
}

func TestAstPrinter_GroupingWithUnary(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, "(group (- (group 123)))", func() {
		literalExpr := ast.LiteralExpr{Value: 123}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		unaryExpr := ast.UnaryExpr{Operator: &minusToken, Right: &groupExpr}
		groupUnaryExpr := ast.GroupingExpr{Expression: &unaryExpr}
		groupUnaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, "(! (group (- (group 123))))", func() {
		literalExpr := ast.LiteralExpr{Value: 123}
		groupExpr := ast.GroupingExpr{Expression: &literalExpr}
		unaryExpr := ast.UnaryExpr{Operator: &minusToken, Right: &groupExpr}
		groupUnaryExpr := ast.GroupingExpr{Expression: &unaryExpr}
		parentUnaryEpr := ast.UnaryExpr{Operator: &bangToken, Right: &groupUnaryExpr}
		parentUnaryEpr.Accept(printer)
	})
}

func TestAstPrinter_BinaryExpr(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, `(* "hello" 123)`, func() {
		leftExpr := ast.LiteralExpr{Value: "hello"}
		rightExpr := ast.LiteralExpr{Value: 123}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &starToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, "(* nil 123)", func() {
		leftExpr := ast.LiteralExpr{Value: nil}
		rightExpr := ast.LiteralExpr{Value: 123}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &starToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, "(* nil nil)", func() {
		leftExpr := ast.LiteralExpr{Value: nil}
		rightExpr := ast.LiteralExpr{Value: nil}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &starToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
}

func TestAstPrinter_BinaryExprWithSubExpressions(t *testing.T) {
	printer := NewAstPrinter()
	verifyPrintedToStdout(t, `(* (group "hello") (group 123))`, func() {
		leftExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: "hello"}}
		rightExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: 123}}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &starToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(= "hello" (group 123))`, func() {
		leftExpr := ast.LiteralExpr{Value: "hello"}
		rightExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: 123}}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &equalToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(* (group "hello") (- 123))`, func() {
		leftExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: "hello"}}
		rightExpr := ast.UnaryExpr{Operator: &minusToken, Right: &ast.LiteralExpr{Value: 123}}
		binaryExpr := ast.BinaryExpr{Left: &leftExpr, Operator: &starToken, Right: &rightExpr}
		binaryExpr.Accept(printer)
	})
	verifyPrintedToStdout(t, `(+ (* (group 2) (- 3)) (* (group "hello") 123))`, func() {
		leftLeftExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: 2}}
		leftRightExpr := ast.UnaryExpr{Operator: &minusToken, Right: &ast.LiteralExpr{Value: 3}}
		leftBinaryExpr := ast.BinaryExpr{Left: &leftLeftExpr, Operator: &starToken, Right: &leftRightExpr}

		rightLeftExpr := ast.GroupingExpr{Expression: &ast.LiteralExpr{Value: "hello"}}
		rightRightExpr := ast.LiteralExpr{Value: 123}
		rightBinaryExpr := ast.BinaryExpr{Left: &rightLeftExpr, Operator: &starToken, Right: &rightRightExpr}

		binaryExpr := ast.BinaryExpr{Left: &leftBinaryExpr, Operator: &plusToken, Right: &rightBinaryExpr}
		binaryExpr.Accept(printer)
	})
}
