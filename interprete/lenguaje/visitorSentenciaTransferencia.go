package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) interface{} {
    l.shouldBreak = true // Establecer shouldBreak en verdadero
	return nil
}