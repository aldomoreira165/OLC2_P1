package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// funcion sin retorno y sin parametros
func (l *Visitor) VisitFuncionNormal(ctx *parser.FuncionNormalContext) interface{} {
	funcId := ctx.ID().GetText()
	funcBlock := ctx.Block()

	if _, ok := l.currentEnvironment.funciones[funcId]; ok {
		return fmt.Sprintf("Error funcion ya existente en el entorno actual: %s", funcId)
	}
	if ctx.Parametros() != nil {
		params := l.Visit(ctx.Parametros()).([]ParametroDef)
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       "nil",
			Sentencias: funcBlock,
			Parametros: params,
		}
	} else {
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       "nil",
			Sentencias: funcBlock,
			Parametros: nil,
		}
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
	if ctx.Parametros() != nil {
		params := l.Visit(ctx.Parametros()).([]ParametroDef)
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       tipo,
			Sentencias: funcBlock,
			Parametros: params,
		}
	} else {
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       tipo,
			Sentencias: funcBlock,
			Parametros: nil,
		}
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
				var argValues map[string]interface{}
				if ctx.Parametroscall() != nil {
					argValues = l.Visit(ctx.Parametroscall()).(map[string]interface{})
					previousEnvironment := CrearEntorno(l)
					defer EliminarEntorno(l, previousEnvironment)
					parametros := function.Parametros
					for paramName, paramValue := range argValues {
						for _, param := range parametros {
							fmt.Println(param)
							if param.Externo == paramName {
								if param.Tipo == "vector" {
									fmt.Println("entro a vector para la funcion XDDDDD")
									valores := paramValue.([]interface{})
									tipo := determineType(valores[0])
									l.agregarVector(param.Interno, tipo, valores)
								} else {
									l.agregarVariable(param.Interno, determineType(paramValue), false, paramValue)
								}
							}
						}
					}
					retValue := l.Visit(function.Sentencias)
					retornarInterface := interface{}(retValue)
					retorno := convertirExpresionD(function.Tipo, retornarInterface)
					return retorno
				} else {
					previousEnvironment := CrearEntorno(l)
					defer EliminarEntorno(l, previousEnvironment)
					retValue := l.Visit(function.Sentencias)
					retornarInterface := interface{}(retValue)
					retorno := convertirExpresionD(function.Tipo, retornarInterface)
					return retorno
				}
			} else {
				var argValues map[string]interface{}
				if ctx.Parametroscall() != nil {
					argValues = l.Visit(ctx.Parametroscall()).(map[string]interface{})
					previousEnvironment := CrearEntorno(l)
					defer EliminarEntorno(l, previousEnvironment)
					parametros := function.Parametros
					for paramName, paramValue := range argValues {
						for _, param := range parametros {
							if param.Externo == paramName {
								//validar si es una variable, vector o matriz
								if param.Tipo == "vector" {
									fmt.Println("entro a vector para la funcion XDDDDD")
									valores := paramValue.([]interface{})
									tipo := determineType(valores[0])
									l.agregarVector(param.Interno, tipo, valores)
								} else {
									fmt.Println("entro a variable para la funcion XDDDDD")
									l.agregarVariable(param.Interno, determineType(paramValue), false, paramValue)
								}
							}
						}
					}
					fmt.Println("entorno actual", l.currentEnvironment.variables)
					return l.Visit(function.Sentencias)
				} else {
					fmt.Println("dentro de funcion normal sin parametros")
					previousEnvironment := CrearEntorno(l)
					defer EliminarEntorno(l, previousEnvironment)
					return l.Visit(function.Sentencias)
				}
			}
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error funcion no encontrada: %s", funcId)
}

// parametros de la funcion
func (l *Visitor) VisitParametros(ctx *parser.ParametrosContext) interface{} {
	paramList := ctx.AllID()

	params := make([]ParametroDef, 0)

	for i := 0; i < len(paramList); i += 2 {
		externo := paramList[i].GetText()
		interno := paramList[i+1].GetText()
		tipo := l.Visit(ctx.Tipo(i / 2)).(string)


		param := ParametroDef{
			Externo: externo,
			Interno: interno,
			Tipo:    tipo,
		}
		params = append(params, param)
	}
	return params
}

func (l *Visitor) VisitParametroscall(ctx *parser.ParametroscallContext) interface{} {
	argValues := map[string]interface{}{}

	for i := 0; i < len(ctx.AllID()); i++ {
		interno := ctx.ID(i).GetText()
		argValue := l.Visit(ctx.Expr(i))
		argValues[interno] = argValue
	}
	return argValues
}
