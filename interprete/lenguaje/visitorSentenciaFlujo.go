package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit del if
func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	result := l.Visit(ctx.Expr())

	if result == true {
		previousEnvironment := CrearEntorno(l)
		defer EliminarEntorno(l, previousEnvironment) // Esto asegura que el entorno se elimine al salir del bloque
		return l.Visit(ctx.Block(0))
	} else {
		// Recorrer los bloques else if
		for _, elseifCtx := range ctx.AllElseifstmt() {
			elseifValue := l.Visit(elseifCtx.Expr())
			if elseifValue == true {
				previousEnvironment := CrearEntorno(l)
				defer EliminarEntorno(l, previousEnvironment)
				return l.Visit(elseifCtx.Block())
			}
		}

		if ctx.ELSE() != nil {
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			return l.Visit(ctx.Block(1))
		}

	}
	return nil
}

func (l *Visitor) VisitElseifstmt(ctx *parser.ElseifstmtContext) interface{} {
	elseifValue := l.Visit(ctx.Expr()).(bool)

	if elseifValue {
		l.Visit(ctx.Block())
	}
	return nil
}

// switch
func (l *Visitor) VisitSwitchstmt(ctx *parser.SwitchstmtContext) interface{} {
	expresion := l.Visit(ctx.Expr())
	exprType := determineType(expresion)
	valorExpresion := convertirExpresion(exprType, expresion)
	for _, caseCtx := range ctx.AllCaseStmt() {
		caseSwitch := l.Visit(caseCtx.Expr())
		caseSwitchType := determineType(caseSwitch)
		valorCase := convertirExpresion(caseSwitchType, caseSwitch)

		if valorExpresion == valorCase && exprType == caseSwitchType {
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			contenido := l.Visit(caseCtx.Block())  
			if l.shouldBreak {
				l.shouldBreak = false
				return contenido
			}
			return contenido
		}
	}

	// Si no se encontró un caso, ejecutar el caso por defecto si existe
	if ctx.DefaultCase() != nil {
		previousEnvironment := CrearEntorno(l)
		defer EliminarEntorno(l, previousEnvironment)
		contenido := l.Visit(ctx.DefaultCase().Block())
		if l.shouldBreak {
			l.shouldBreak = false
			return contenido
		}
		return contenido
	}
	return nil
}

func (l *Visitor) VisitCaseStmt(ctx *parser.CaseStmtContext) interface{} {
	l.Visit(ctx.Block())
	return nil
}

func (l *Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	l.Visit(ctx.Block())
	return nil
}

// visit del while (arreglar) break no jala todavia
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	var out string

	for {
		expresion := l.Visit(ctx.Expr())

		if expresion == true {
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			resultado := l.Visit(ctx.Block())
			out += resultado.(string)
			if l.shouldBreak {
				l.shouldBreak = false
				return out
			}
			if l.shouldContinue {
				l.shouldContinue = false
				continue
			}
		} else {
			return out
		}
	}
}

//for
func (l *Visitor) VisitForstmt(ctx *parser.ForstmtContext) interface{} {
	var out string
	if ctx.Rangostmt() != nil {
		posiciones := l.Visit(ctx.Rangostmt()).([]int64)
		inicio := posiciones[0]
		fin := posiciones[1]
		fmt.Println(inicio, fin)

		if inicio > fin {
			return ("Error: el inicio del rango es mayor que el fin del rango \n")
		}

		//crear la variable a iterar en el entorno
		varName := ctx.ID().GetText()
		l.agregarVariable(varName, "nil", true,nil)

		for i := inicio; i <= fin; i++ {
			variable := Variable{
				Name: varName,
				Type:   determineType(i),
				Constante: true,
				Value:  i,
			}
			l.currentEnvironment.variables[varName] = variable

			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			resultado := l.Visit(ctx.Block())
			out += resultado.(string)
			if l.shouldBreak {
				l.shouldBreak = false
				return out
			}else if l.shouldContinue {
				l.shouldContinue = false
				continue		
			}
		}
		return out
	}

	if ctx.Expr() != nil {
		expresion := l.Visit(ctx.Expr())
		if determineType(expresion) == "String" {
			//crear la variable a iterar en el entorno
			varName := ctx.ID().GetText()
			l.agregarVariable(varName, "nil", true,nil)

			for _, char := range expresion.(string) {
				variable := Variable{
					Name: varName,
					Type:   "character",
					Value:  string(char),
				}
				l.currentEnvironment.variables[varName] = variable

				previousEnvironment := CrearEntorno(l)
				defer EliminarEntorno(l, previousEnvironment)
				resultado := l.Visit(ctx.Block())
				out += resultado.(string)
				if l.shouldBreak {
					l.shouldBreak = false
					return out
				}else if l.shouldContinue {
					l.shouldContinue = false
					continue
				}
			}
			return out
		}
		//buscar si es un vector buscando el id de la variable en la lista de vectores del entorno
		if vector, ok := l.currentEnvironment.Vectores[ctx.Expr().GetText()]; ok {
			//crear la variable a iterar en el entorno
			varName := ctx.ID().GetText()
			l.agregarVariable(varName, "nil", true,nil)

			for _, valor := range vector.Valores {
				variable := Variable{
					Name: varName,
					Type:   determineType(valor),
					Value:  valor,
				}
				l.currentEnvironment.variables[varName] = variable

				previousEnvironment := CrearEntorno(l)
				defer EliminarEntorno(l, previousEnvironment)
				resultado := l.Visit(ctx.Block())
				out += resultado.(string)
				if l.shouldBreak {
					l.shouldBreak = false
					return out
				}else if l.shouldContinue {
					l.shouldContinue = false
					continue
				}
			}
			return out
		}
	}
	return nil
}

//guard
func (l *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
	condicion := l.Visit(ctx.Expr())
	if condicion == false{
		if ctx.Block() != nil{
			return l.Visit(ctx.Block())
		}else{
			return nil
		}
	}
	return nil
}

//rango
func (l *Visitor) VisitRangostmt(ctx *parser.RangostmtContext) interface{} {
	inicio := l.Visit(ctx.Expr(0)).(int64)
	fin := l.Visit(ctx.Expr(1)).(int64)
	//agregar en un arreglo de dos posiciones
	recorrido := []int64{inicio, fin}
	return recorrido
}

func (l *Visitor) VisitTipo(ctx *parser.TipoContext) interface{} {
	if ctx.INT() != nil {
		return "int"
	} else if ctx.FLOAT() != nil {
		return "float"
	} else if ctx.BOOL() != nil {
		return "bool"
	} else if ctx.CHARACTER() != nil {
		return "character"
	} else if ctx.PSTRING() != nil {
		return "String"
	}
	return nil
}