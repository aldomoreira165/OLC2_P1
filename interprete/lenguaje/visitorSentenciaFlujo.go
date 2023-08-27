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
			return l.Visit(caseCtx.Block())               // Salir del switch después de ejecutar un caso válido
		}
	}

	// Si no se encontró un caso, ejecutar el caso por defecto si existe
	if ctx.DefaultCase() != nil {
		previousEnvironment := CrearEntorno(l)
		defer EliminarEntorno(l, previousEnvironment)
		return l.Visit(ctx.DefaultCase().Block())
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
		} else {
			l.shouldBreak = false
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
		for i := inicio; i <= fin; i++ {
			fmt.Println("en ciclo")
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment)
			resultado := l.Visit(ctx.Block())
			out += resultado.(string)
		}
		return out
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