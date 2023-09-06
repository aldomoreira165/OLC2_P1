package main

import (
	"encoding/json"
	"interprete/Parser"
	"interprete/lenguaje"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
)

var nodeCount int

type CodigoEnviado struct {
	Contenido string `json:"contenido"`
}

/*func generarCST(code string) (string, error) {
	// Genera el CST aquí utilizando ANTLR4
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()

	// Generar el código DOT del CST
	dot := generateDOT(tree)

	return dot, nil
}

func generateDOT(tree antlr.Tree) string {
	var dotBuilder strings.Builder
	dotBuilder.WriteString("digraph CST {\n")
	visitDOT(tree, &dotBuilder)
	dotBuilder.WriteString("}\n")
	return dotBuilder.String()
}

func visitDOT(tree antlr.Tree, dotBuilder *strings.Builder) {
	if tree == nil {
		return
	}

	// Obtener el texto del nodo
	nodeText := tree.GetText()

	// Escapar comillas dobles en el texto del nodo
	nodeText = strings.ReplaceAll(nodeText, "\"", "\\\"")

	// Agregar el nodo al código DOT
	dotBuilder.WriteString("\"")
	dotBuilder.WriteString(nodeText)
	dotBuilder.WriteString("\";\n")

	// Conectar el nodo con su padre si lo tiene
	if parent := tree.GetParent(); parent != nil {
		parentText := parent
		dotBuilder.WriteString("\"")
		dotBuilder.WriteString(parentText)
		dotBuilder.WriteString("\" -> \"")
		dotBuilder.WriteString(nodeText)
		dotBuilder.WriteString("\";\n")
	}

	// Recorrer los hijos
	for i := 0; i < tree.GetChildCount(); i++ {
		visitDOT(tree.GetChild(i), dotBuilder)
	}
}*/

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
	visitor := lenguaje.NewVisitor()
	tree := p.S()
	out := visitor.Visit(tree)
	
	// Generar el código DOT del CST
	/*cst, err := generarCST(code)
	if err != nil {
		http.Error(w, "Error al generar el CST", http.StatusInternalServerError)
		return
	}
	fmt.Println(cst) // Imprimir el código DOT en la consola*/

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
