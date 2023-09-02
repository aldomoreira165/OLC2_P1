package lenguaje

import (
	"interprete/Parser"
)

func (l *Visitor) VisitDefStruct(ctx *parser.DefStructContext) interface{} {
	print("VisitDefStruct\n")
	return nil
}

func (l *Visitor) VisitAtributoStruct(ctx *parser.AtributoStructContext) interface{} {
	print("VisitAtributoStruct\n")
	return nil
}

func (l *Visitor) VisitStructExpr(ctx *parser.StructExprContext) interface{} {
	print("VisitStructExpr\n")
	return nil
}

func (l *Visitor) VisitDuplastruct(ctx *parser.DuplastructContext) interface{} {
	print("VisitDuplastruct\n")
	return nil
}

func (l *Visitor) VisitAccesoStruct(ctx *parser.AccesoStructContext) interface{} {
	print("VisitAccesoStruct\n")
	return nil
}






