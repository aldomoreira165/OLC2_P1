package main

import (
	"encoding/json"
	"fmt"
	parser "interprete/Parser"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// estructura para los entornos
type Environment struct {
	parent    *Environment
	variables map[string]Variable
}

// estructura para las variables
type Variable struct {
	Name  string
	Type  string
	Value interface{}
}

// estructura visitor
type Visitor struct {
	antlr.ParseTreeVisitor
	currentEnvironment *Environment
}

// constructor del entorno
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]Variable),
	}
}

// inicializando el visitor
func NewVisitor() parser.SwiftGrammarVisitor {
	//entorno global
	globalEnvironment := NewEnvironment(nil)
	return &Visitor{
		currentEnvironment: globalEnvironment,
	}
}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	return l.Visit(ctx.Block())
}

func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	//Se crea un nuevo entorno local para el bloque
	localEnvironment := NewEnvironment(l.currentEnvironment)
	fmt.Println("se ha creado un nuevo entorno")
	// Cambiar al nuevo entorno local
	previousEnvironment := l.currentEnvironment
	l.currentEnvironment = localEnvironment

	// Procesar las declaraciones y sentencias dentro del bloque
	var out string
	for i := 0; ctx.Stmt(i) != nil; i++ {
		stmtResult := l.Visit(ctx.Stmt(i))
		switch stmtResult.(type) {
		case int64:
			out += strconv.FormatInt(stmtResult.(int64), 10) + "\n"
		case string:
			out += stmtResult.(string) + "\n"
		}
	}

	// Restaurar el entorno anterior al salir del bloque
	l.currentEnvironment = previousEnvironment
	return out
}

func (l *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Printstmt() != nil {
		return l.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		return l.Visit(ctx.Ifstmt())
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
	return nil
}

func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	stringValue := fmt.Sprint(returnValue)
	return stringValue
}

func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	result := l.Visit(ctx.Expr())

	if result == true {
		return l.Visit(ctx.Block())
	}
	return true
}

func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Variable ya existente en el entorno actual: %s", varName)
	}

	// Comprobar si el tipo de la expresión coincide con el tipo declarado
	if !validateType(value, declType) {
		return fmt.Sprintf("Error de tipo en la declaración: %s", varName)
	}

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  declType,
		Value: value,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Variable ya existente en el entorno actual: %s", varName)
	}

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  declType,
		Value: nil,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true

}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Variable ya existente en el entorno actual: %s", varName)
	}

	// Determinar el tipo de la expresión
	valueType := determineType(value)
	fmt.Println(valueType)

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  valueType, // Asignar el tipo de la expresión
		Value: value,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true
}

func (l *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if variableExistente, ok := l.currentEnvironment.variables[varName]; ok {

		// Determinar el tipo de la expresión y el tipo de la variable existente
		exprType := determineType(value)
		existingVarType := variableExistente.Type

		if exprType == existingVarType {

			variableActualizada := Variable{
				Name:  varName,
				Type:  exprType, // Asignar el tipo de la expresión
				Value: value,
			}

			l.currentEnvironment.variables[varName] = variableActualizada
			return true // La asignación fue exitosa
		} else {
			fmt.Println("Error de tipo en la asignación para", varName)
			return false // La asignación falló debido a un error de tipo
		}
	} else {
		fmt.Println("Variable no encontrada:", varName)
		return false // La variable no existe en el entorno actual
	}
}

func (l *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) interface{} {
    for {
        expresion := l.Visit(ctx.Expr())

        // Se verifica si la expresión es verdadera
        if expresion == true {
            // Se ejecuta el bloque de código dentro del bucle
            l.Visit(ctx.Block())
        } else {
            break // Se sale del bucle si la expresion es falsa
        }
    }
    return nil
}

func determineType(value interface{}) string {
	switch value.(type) {
	case int64:
		return "int"
	case float64:
		//ARREGLARRRR
		return "float"
	case bool:
		//ARREGLARRRR
		return "bool"
	case string:
		if len(value.(string)) == 1 {
			return "character"
		}
		return "String"
	default:
		return "unknown"
	}
}

func validateType(value interface{}, declType string) bool {
	switch declType {
	case "int":
		_, isInt := value.(int64)
		return isInt
	case "float":
		//ARREGLARRRRRRRR
		_, isFloat := value.(float64)
		return isFloat
	case "bool":
		//ARREGLARRRRRRRR
		_, isBool := value.(bool)
		return isBool
	case "character":
		_, isChar := value.(string)
		return isChar && len(value.(string)) == 1
	case "String":
		_, isString := value.(string)
		return isString && len(value.(string)) > 1
	default:
		return false // Tipo desconocido o no admitido
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

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	m := l.Visit(ctx.GetLeft()).(int64)
	r := l.Visit(ctx.GetRight()).(int64)
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return m + r
	case "-":
		return m - r
	case "*":
		return m * r
	case "/":
		return m / r
	case "<":
		if m < r {
			return true
		} else {
			return false
		}
	}
	return true
}

func (l *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return l.Visit(ctx.Expr())
}

func (l *Visitor) VisitPrimExpr(ctx *parser.PrimExprContext) interface{} {
	return l.Visit(ctx.Primitivo())
}

func (l *Visitor) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return i
}

func (l *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()
	// Buscar la variable en los entornos, comenzando por el actual y luego ascendiendo
	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		value, ok := currentEnv.variables[id]
		if ok {
			return value.Value
		}
		currentEnv = currentEnv.parent
	}
	// Devolver un mensaje de error en lugar de lanzar una excepción
	return fmt.Sprintf("Variable no encontrada: %s", id)
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return value
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return value
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

type CodigoEnviado struct {
	Contenido string `json:"contenido"`
}

func manejarEnviarcodigo(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo JSON en una estructura CodigoEnviado
	var codigo CodigoEnviado
	err := json.NewDecoder(r.Body).Decode(&codigo)
	if err != nil {
		http.Error(w, "Error al leer el contenido JSON", http.StatusBadRequest)
		return
	}

	code := codigo.Contenido
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	visitor := NewVisitor()
	tree := p.S()
	out := visitor.Visit(tree)
	//fmt.Println(out)

	// enviando respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/enviar-codigo", manejarEnviarcodigo)
	http.Handle("/", fs)
	http.ListenAndServe("localhost:3000", nil)
}

//antlr4 -Dlanguage=Go -o parser -package parser -visitor *.g4
