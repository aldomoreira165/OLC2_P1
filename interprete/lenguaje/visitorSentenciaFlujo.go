package lenguaje

import (
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
				defer EliminarEntorno(l, previousEnvironment) // Esto asegura que el entorno se elimine al salir del bloque
				return l.Visit(elseifCtx.Block())
			}
		}

		if ctx.ELSE() != nil {
			previousEnvironment := CrearEntorno(l)
			defer EliminarEntorno(l, previousEnvironment) // Esto asegura que el entorno se elimine al salir del bloque
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
			defer EliminarEntorno(l, previousEnvironment) // Esto asegura que el entorno se elimine al salir del bloque
			return l.Visit(caseCtx.Block())               // Salir del switch después de ejecutar un caso válido
		}
	}

	// Si no se encontró un caso, ejecutar el caso por defecto si existe
	if ctx.DefaultCase() != nil {
		previousEnvironment := CrearEntorno(l)
		defer EliminarEntorno(l, previousEnvironment) // Esto asegura que el entorno se elimine al salir del bloque
		return l.Visit(ctx.DefaultCase().Block())
	}
	return nil
}

func (l *Visitor) VisitCaseStmt(ctx *parser.CaseStmtContext) interface{} {
	// No se necesita implementación adicional aquí, ya que los bloques se ejecutan en VisitSwitchstmt
	l.Visit(ctx.Block())
	return nil
}

func (l *Visitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	// No se necesita implementación adicional aquí, ya que los bloques se ejecutan en VisitSwitchstmt
	l.Visit(ctx.Block())
	return nil
}

// visit del while (arreglar) break no jala todavia
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	previousEnvironment := CrearEntorno(l)
	var out string

	for {
		expresion := l.Visit(ctx.Expr())

		if expresion == true && !l.shouldBreak {
			resultado := l.Visit(ctx.Block())
			out += resultado.(string)

			if l.shouldBreak {
				break
			}//implementar en for, while, switch, if
		} else {
			EliminarEntorno(l, previousEnvironment)
			l.shouldBreak = false
			return out
		}
	}
	EliminarEntorno(l, previousEnvironment)
	l.shouldBreak = false
	return out
}
