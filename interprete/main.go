package main

import (
	"encoding/json"
	"github.com/antlr4-go/antlr/v4"
	"interprete/environment"
	"interprete/interfaces"
	"interprete/parser"
	"net/http"
)

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
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
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST
	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Ejecutar(&Ast, nil)
	}
	//fmt.Println(Ast.GetPrint())

	// enviando respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Ast.GetPrint())
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/enviar-codigo", manejarEnviarcodigo)
	http.Handle("/", fs)
	http.ListenAndServe("localhost:3000", nil)
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}

//antlr4 -Dlanguage=Go -o parser -package parser *.g4
