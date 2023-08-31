package lenguaje

import (
	"fmt"
)

type Environment struct {
	parent    *Environment
	variables map[string]Variable
	Vectores  map[string]Vector
	Matrices map[string]Matriz
	Matrices3D map[string]Matriz3D
	funciones map[string]Funcion
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]Variable),
		Vectores:  make(map[string]Vector),
		Matrices: make(map[string]Matriz),
		Matrices3D: make(map[string]Matriz3D),
		funciones: make(map[string]Funcion),
	}
}

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

func (l *Visitor) agregarVariable(name string, tipo string, value interface{}) interface{} {
	if _, ok := l.currentEnvironment.variables[name]; ok {
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", name)
	}

	if validateType(value,tipo) {
		fmt.Println("Se ha agregado una nueva variable")
		nuevaVariable := Variable{
			Name:  name,
			Type:  tipo,
			Value: value,
		}
		l.currentEnvironment.variables[name] = nuevaVariable
	}
	return nil
}

func (l *Visitor) agregarVector(id string, tipo string, valores []interface{}) interface{} {
	if _, ok := l.currentEnvironment.Vectores[id]; ok {
		return fmt.Sprintf("Error vector ya existente en el entorno actual: %s",id)
	}

	nuevoVector := Vector{
		Id:  id,
		Tipo:  tipo,
		Valores: valores,
	}
	l.currentEnvironment.Vectores[id] = nuevoVector
	return true
}
