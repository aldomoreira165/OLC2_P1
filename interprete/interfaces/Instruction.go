package interfaces

import "interprete/environment"

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}) interface{}
}
