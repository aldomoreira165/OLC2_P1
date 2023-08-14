package lenguaje

import (
	"fmt"
	parser "interprete/Parser"
	"log"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	currentEnvironment *Environment
}

func NewVisitor() parser.SwiftGrammarVisitor {
	globalEnvironment := NewEnvironment(nil)
	return &Visitor{
		currentEnvironment: globalEnvironment,
	}
}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	return l.Visit(ctx.Block())
}

// visit de block
func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	var out string
	for i := 0; ctx.Stmt(i) != nil; i++ {
		stmtResult := l.Visit(ctx.Stmt(i))
		switch stmtResult.(type) {
		case int64:
			out += strconv.FormatInt(stmtResult.(int64), 10) + "\n"
		case float64: // Nuevo caso para números decimales
			out += strconv.FormatFloat(stmtResult.(float64), 'f', -1, 64) + "\n"
		case string:
			out += stmtResult.(string) + "\n"
		}
	}
	return out
}

// visit de las sentencias
func (l *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Printstmt() != nil {
		return l.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		return l.Visit(ctx.Ifstmt())
	}
	if ctx.Switchstmt() != nil {
		return l.Visit(ctx.Switchstmt())
	}
	if ctx.TypedDeclstmt() != nil {
		return l.Visit(ctx.TypedDeclstmt())
	}
	if ctx.UntypedDeclstmt() != nil {
		return l.Visit(ctx.UntypedDeclstmt())
	}
	if ctx.OptionalTypedDeclstmt() != nil {
		return l.Visit(ctx.OptionalTypedDeclstmt())
	}
	if ctx.Asignstmt() != nil {
		return l.Visit(ctx.Asignstmt())
	}
	if ctx.Whilestmt() != nil {
		return l.Visit(ctx.Whilestmt())
	}
	if ctx.Opasignstmt() != nil {
		return l.Visit(ctx.Opasignstmt())
	}
	return nil
}

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	stringValue := fmt.Sprint(returnValue)
	return stringValue
}

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
		valorCase := convertirExpresion(caseSwitchType,caseSwitch)
		
		if valorExpresion == valorCase && exprType == caseSwitchType {
			return l.Visit(caseCtx.Block()) // Salir del switch después de ejecutar un caso válido
		}
	}

	// Si no se encontró un caso, ejecutar el caso por defecto si existe
	if ctx.DefaultCase() != nil {
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

// visit del while (arreglar)
func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
	// Crear un nuevo entorno antes de entrar al ciclo
	previousEnvironment := CrearEntorno(l)

	var out string

	for {
		expresion := l.Visit(ctx.Expr())

		// Se verifica si la expresión es verdadera
		if expresion == true {
			// Se ejecuta el bloque de código dentro del bucle
			resultado := l.Visit(ctx.Block())
			out += resultado.(string)
		} else {
			// Salir del bucle y eliminar el entorno creado
			EliminarEntorno(l, previousEnvironment)
			return out
		}
	}
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

// Funcion visitar
func (l *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		nodo := tree.Accept(l)
		return nodo
	}
}
