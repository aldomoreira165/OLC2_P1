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
		l.simbolos.InsertarSimbolo(funcId, "Funcion", "nil", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	} else {
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       "nil",
			Sentencias: funcBlock,
			Parametros: nil,
		}
		l.simbolos.InsertarSimbolo(funcId, "Funcion", "nil", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
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
		l.simbolos.InsertarSimbolo(funcId, "Funcion", tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	} else {
		l.currentEnvironment.funciones[funcId] = Funcion{
			Id:         funcId,
			Tipo:       tipo,
			Sentencias: funcBlock,
			Parametros: nil,
		}
		l.simbolos.InsertarSimbolo(funcId, "Funcion", tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
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
									valores := paramValue.(ParametroCall).Valor.([]interface{})
									tipo := determineType(valores[0])
									l.agregarVector(param.Interno, tipo, valores)
								} else {
									l.agregarVariable(param.Interno, determineType(paramValue.(ParametroCall).Valor), false, paramValue.(ParametroCall).Valor)
								}
							}
						}
					}
					retValue := l.Visit(function.Sentencias)
					retornarInterface := interface{}(retValue)
					retorno := convertirExpresionD(function.Tipo, retornarInterface)
					for _, param := range parametros {
						if param.Inout {
							fmt.Println("es inout", param)
							nombreBuscar := param.Externo
							//buscar en argvalues
							for paramName, paramValue := range argValues {
								if paramName == nombreBuscar {
									nombreVariable := paramValue.(ParametroCall).NameVar
									//buscar variable y asignar valor de param
									currentEnv := l.currentEnvironment
									for currentEnv != nil {
										if param.Tipo == "vector" {
											fmt.Println("valores: ", l.currentEnvironment.Vectores[param.Interno].Valores)
											if _, ok := currentEnv.Vectores[nombreVariable]; ok {
												nuevoVector := Vector{
													Id:      nombreVariable,
													Tipo:    determineType(l.currentEnvironment.Vectores[param.Interno].Valores[0]),
													Valores: l.currentEnvironment.Vectores[param.Interno].Valores,
												}
												currentEnv.Vectores[nombreVariable] = nuevoVector
												break
											}
										} else {
											if _, ok := currentEnv.variables[nombreVariable]; ok {

												nuevaVariable := Variable{
													Name:      nombreVariable,
													Type:      param.Tipo,
													Constante: false,
													Value:     l.currentEnvironment.variables[param.Interno].Value,
												}
												currentEnv.variables[nombreVariable] = nuevaVariable
												break
											}
										}
										currentEnv = currentEnv.parent
									}
								}
							}
						}
					}
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
					fmt.Println("argValues: ", argValues)
					previousEnvironment := CrearEntorno(l)
					defer EliminarEntorno(l, previousEnvironment)
					parametros := function.Parametros
					for paramName, paramValue := range argValues {
						for _, param := range parametros {
							if param.Externo == paramName {
								fmt.Println("TIPPPPPPPPPPPPPPOOOOOOOOO", param.Tipo)
								//validar si es una variable, vector o matriz
								if param.Tipo == "vector" {
									valores := paramValue.(ParametroCall).Valor.([]interface{})
									tipo := determineType(valores[0])
									l.agregarVector(param.Interno, tipo, valores)
								} else if param.Tipo == "matriz2" {
									fmt.Println("es matriz2")
									valores := paramValue.(ParametroCall).Valor.([][]interface{})
									tipo := determineType(valores[0][0])
									matrizNueva := Matriz{
										Id:        param.Interno,
										Tipo:      tipo,
										Dimension: 2,
										Valores:   valores,
									}
									l.currentEnvironment.Matrices[param.Interno] = matrizNueva
								} else if param.Tipo == "matriz3" {
									fmt.Println("es matriz3")
									valores := paramValue.(ParametroCall).Valor.([][][]interface{})
									tipo := determineType(valores[0][0][0])
									matrizNueva := Matriz3D{
										Id:        param.Interno,
										Tipo:      tipo,
										Dimension: 3,
										Valores:   valores,
									}
									l.currentEnvironment.Matrices3D[param.Interno] = matrizNueva
								} else {
									l.agregarVariable(param.Interno, determineType(paramValue.(ParametroCall).Valor), false, paramValue.(ParametroCall).Valor)
								}

							}
						}
					}
					instrucciones := l.Visit(function.Sentencias)
					//actualizar variables inout en los entornos anteriores
					for _, param := range parametros {
						if param.Inout {
							fmt.Println("es inout", param)
							nombreBuscar := param.Externo
							//buscar en argvalues
							for paramName, paramValue := range argValues {
								if paramName == nombreBuscar {
									nombreVariable := paramValue.(ParametroCall).NameVar
									//buscar variable y asignar valor de param
									currentEnv := l.currentEnvironment
									for currentEnv != nil {
										if param.Tipo == "vector" {
											fmt.Println("valores: ", l.currentEnvironment.Vectores[param.Interno].Valores)
											if _, ok := currentEnv.Vectores[nombreVariable]; ok {
												nuevoVector := Vector{
													Id:      nombreVariable,
													Tipo:    determineType(l.currentEnvironment.Vectores[param.Interno].Valores[0]),
													Valores: l.currentEnvironment.Vectores[param.Interno].Valores,
												}
												currentEnv.Vectores[nombreVariable] = nuevoVector
												break
											}
										} else if param.Tipo == "matriz2" {
											return true
										} else if param.Tipo == "matriz3" {
											return true
										}else{
											if _, ok := currentEnv.variables[nombreVariable]; ok {

												nuevaVariable := Variable{
													Name:      nombreVariable,
													Type:      param.Tipo,
													Constante: false,
													Value:     l.currentEnvironment.variables[param.Interno].Value,
												}
												currentEnv.variables[nombreVariable] = nuevaVariable
												break
											}
										}
										currentEnv = currentEnv.parent
									}
								}
							}
						}
					}
					return instrucciones
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

		if ctx.INOUT(i/2) != nil {
			param := ParametroDef{
				Externo: externo,
				Interno: interno,
				Inout:   true,
				Tipo:    tipo,
			}
			params = append(params, param)
		} else {
			param := ParametroDef{
				Externo: externo,
				Interno: interno,
				Inout:   false,
				Tipo:    tipo,
			}
			params = append(params, param)
		}
		fmt.Println(params)
	}
	return params
}

func (l *Visitor) VisitParametroscall(ctx *parser.ParametroscallContext) interface{} {
	argValues := map[string]interface{}{}

	for i := 0; i < len(ctx.AllID()); i++ {
		externo := ctx.ID(i).GetText()

		interno := ParametroCall{
			NameVar: ctx.Expr(i).GetText(),
			Valor:   l.Visit(ctx.Expr(i)),
		}
		fmt.Println("valorrrrrr", interno)
		argValues[externo] = interno
	}
	return argValues
}
