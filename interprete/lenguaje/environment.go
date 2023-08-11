package lenguaje

import "fmt"

type Environment struct {
	parent    *Environment
	variables map[string]Variable
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]Variable),
	}
}

// Crear un nuevo entorno local
func CrearEntorno(visit *Visitor) *Environment {
	localEnvironment := NewEnvironment(visit.currentEnvironment)
	fmt.Println("Se ha creado un nuevo entorno")
	previousEnvironment := visit.currentEnvironment
	visit.currentEnvironment = localEnvironment
	return previousEnvironment
}

func EliminarEntorno(visit *Visitor, previousEnvironment *Environment) {
	fmt.Println("Se ha eliminado el entorno")
	visit.currentEnvironment = previousEnvironment
}

