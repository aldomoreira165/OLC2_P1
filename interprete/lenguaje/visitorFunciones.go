package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit de las sentencias
func (l *Visitor) VisitFuncdclstmt(ctx *parser.FuncdclstmtContext) interface{} {
	funcId := ctx.ID().GetText()
	funcBlock := ctx.Block()

	if _, ok := l.currentEnvironment.funciones[funcId]; ok {
		return fmt.Sprintf("Error funcion ya existente en el entorno actual: %s", funcId)
	}

	l.currentEnvironment.funciones[funcId] = Funcion{
		Id:  funcId,
		Sentencias: funcBlock,
	}
	return nil
}

func (l *Visitor) VisitAccfuncstm(ctx *parser.AccfuncstmContext) interface{} {
	funcId := ctx.ID().GetText()
	
	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		function, ok := currentEnv.funciones[funcId]
		if ok {
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			return l.Visit(function.Sentencias)
		}
		currentEnv = currentEnv.parent
	} 
	return fmt.Sprintf("Error funcion no encontrada: %s", funcId)
}