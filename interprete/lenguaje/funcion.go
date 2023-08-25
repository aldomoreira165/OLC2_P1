package lenguaje

import (
	"github.com/antlr4-go/antlr/v4"
)

type Funcion struct {
	Id  string
	Tipo string
	Sentencias antlr.ParserRuleContext
	Parametros []ParametroDef
}

type ParametroDef struct {
	Externo string
	Interno string
	Tipo    string
}

