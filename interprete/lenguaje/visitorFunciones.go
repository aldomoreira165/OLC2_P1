package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit de las sentencias
func (l *Visitor) VisitFuncionNormal(ctx *parser.FuncionNormalContext) interface{} {
	funcId := ctx.ID().GetText()
	funcBlock := ctx.Block()

	if _, ok := l.currentEnvironment.funciones[funcId]; ok {
		return fmt.Sprintf("Error funcion ya existente en el entorno actual: %s", funcId)
	}

	l.currentEnvironment.funciones[funcId] = Funcion{
		Id:  funcId,
		Tipo: "nil",
		Sentencias: funcBlock,
	}
	return nil
}

func (l *Visitor) VisitFuncionRetorno(ctx *parser.FuncionRetornoContext) interface{} {
	funcId := ctx.ID().GetText()
	funcBlock := ctx.Block()
	tipo := ctx.Tipo().GetText()

	if _, ok := l.currentEnvironment.funciones[funcId]; ok {
		return fmt.Sprintf("Error funcion ya existente en el entorno actual: %s", funcId)
	}

	l.currentEnvironment.funciones[funcId] = Funcion{
		Id:  funcId,
		Tipo: tipo,
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
			if function.Tipo != "nil" {
				previousEnvironment := CrearEntorno(l)
				defer EliminarEntorno(l, previousEnvironment)
				retValue := l.Visit(function.Sentencias)
				return retValue
			}else{
				previousEnvironment := CrearEntorno(l)
				defer EliminarEntorno(l, previousEnvironment)
				return l.Visit(function.Sentencias)
			}
		}
		currentEnv = currentEnv.parent
	} 
	return fmt.Sprintf("Error funcion no encontrada: %s", funcId)
}