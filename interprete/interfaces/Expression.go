package interfaces

import "interprete/environment"

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}) environment.Symbol
}
