package main

import (
	"encoding/json"
	"fmt"
	"interprete/Parser"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// entornos
type Environment struct {
	parent    *Environment
	variables map[string]interface{}
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]interface{}),
	}
}

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

func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// Crear un nuevo entorno local para el bloque
	localEnvironment := NewEnvironment(l.currentEnvironment)
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
	if ctx.Declstmt() != nil {
		return l.Visit(ctx.Declstmt())
	}
	return nil
}

func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	stringValue := fmt.Sprint(returnValue)
	//fmt.Println(returnValue)
	return stringValue
}

func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	result := l.Visit(ctx.Expr())

	if result == true {
		return l.Visit(ctx.Block())
	}
	return true
}

func (l *Visitor) VisitDeclstmt(ctx *parser.DeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Variable ya existente en el entorno actual: %s", varName)
	} else {
		l.currentEnvironment.variables[varName] = value
	}
	return true
}

func (l *Visitor) VisitPtipo(ctx *parser.PtipoContext) interface{} {
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
			return value
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

