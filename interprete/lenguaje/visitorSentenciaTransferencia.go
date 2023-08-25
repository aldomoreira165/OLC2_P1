package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) interface{} {
    //l.shouldBreak = true // Establecer shouldBreak en verdadero
	return nil
}

func (l *Visitor) VisitReturnstmt(ctx *parser.ReturnstmtContext) interface{} {
    exp := l.Visit(ctx.Expr())
	return exp
}