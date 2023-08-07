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

type Visitor struct {
	antlr.ParseTreeVisitor
	memory map[string]interface{}
}

func NewVisitor() parser.SwiftGrammarVisitor {
	return &Visitor{
		ParseTreeVisitor: &parser.BaseSwiftGrammarVisitor{},
	}

}

func (l *Visitor) VisitS(ctx *parser.SContext) interface{} {
	return l.Visit(ctx.Block())
}

func (l *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	var out string
	for i := 0; ctx.Stmt(i) != nil; i++ {
		stmtResult := l.Visit(ctx.Stmt(i))
		switch stmtResult.(type) {
		case int64:
			out += strconv.FormatInt(stmtResult.(int64), 10) + " "
		case string:
			out += stmtResult.(string) + " "
		}
	}
	return out
}

func (l *Visitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.Printstmt() != nil {
		return l.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		// return l.Visit(ctx.Ifstmt())
		return "If Execute!"
	}
	return nil
}

func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	fmt.Println(returnValue)
	return returnValue
}

func (l *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
	fmt.Println("If Implements..")
	return true
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
	value, ok := l.memory[id]
	if ok {
		return value
	} else {
		panic("no existe la variable " + id)
	}
	return nil
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
	fmt.Println(out)
	//fmt.Println(code)
	
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