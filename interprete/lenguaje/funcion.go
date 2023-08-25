package lenguaje

import (
	"github.com/antlr4-go/antlr/v4"
)

type Funcion struct {
	Id  string
	Tipo string
	Sentencias antlr.ParserRuleContext
}

