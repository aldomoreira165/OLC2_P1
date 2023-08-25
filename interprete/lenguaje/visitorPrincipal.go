package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"log"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	currentEnvironment *Environment
	shouldBreak bool
}

func NewVisitor() parser.SwiftGrammarVisitor {
	globalEnvironment := NewEnvironment(nil)
	return &Visitor{
		currentEnvironment: globalEnvironment,
		shouldBreak: false,
	}
}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	return l.Visit(ctx.Block())
}

// visit de block
func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	var out string

	for i := 0; ctx.Stmt(i) != nil; i++ {
		stmtResult := l.Visit(ctx.Stmt(i))
		switch stmtResult.(type) {
		case int64:
			out += strconv.FormatInt(stmtResult.(int64), 10) + "\n"
		case float64: // Nuevo caso para números decimales
			out += strconv.FormatFloat(stmtResult.(float64), 'f', -1, 64) + "\n"
		case string:
			out += stmtResult.(string) + "\n"
		}
	}
	return out
}

// visit de las sentencias
func (l *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Printstmt() != nil {
		return l.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		return l.Visit(ctx.Ifstmt())
	}
	if ctx.Switchstmt() != nil {
		return l.Visit(ctx.Switchstmt())
	}
	if ctx.TypedDeclstmt() != nil {
		return l.Visit(ctx.TypedDeclstmt())
	}
	if ctx.UntypedDeclstmt() != nil {
		return l.Visit(ctx.UntypedDeclstmt())
	}
	if ctx.OptionalTypedDeclstmt() != nil {
		return l.Visit(ctx.OptionalTypedDeclstmt())
	}
	if ctx.Asignstmt() != nil {
		return l.Visit(ctx.Asignstmt())
	}
	if ctx.Whilestmt() != nil {
		return l.Visit(ctx.Whilestmt())
	}
	if ctx.Opasignstmt() != nil {
		return l.Visit(ctx.Opasignstmt())
	}
	if ctx.Breakstmt() != nil {
		return l.Visit(ctx.Breakstmt())
	}
	if ctx.Returnstmt() != nil {
		return l.Visit(ctx.Returnstmt())
	}
	if ctx.Funcdclstmt() != nil {
		return l.Visit(ctx.Funcdclstmt())
	}
	if ctx.Accfuncstm() != nil {
		return l.Visit(ctx.Accfuncstm())
	}
	return nil
}

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	stringValue := fmt.Sprint(returnValue)
	return stringValue
}

// Funcion visitar
func (l *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		nodo := tree.Accept(l)
		return nodo
	}
}

