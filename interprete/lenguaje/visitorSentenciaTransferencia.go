package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) interface{} {
	l.shouldBreak = true
	return nil
}

func (l *Visitor) VisitReturnstmt(ctx *parser.ReturnstmtContext) interface{} {
    exp := l.Visit(ctx.Expr())
	return exp
}

func (l *Visitor) VisitContinuestmt(ctx *parser.ContinuestmtContext) interface{} {
	l.shouldContinue = true
	return nil
}