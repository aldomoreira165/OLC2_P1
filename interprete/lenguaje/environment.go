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
	Structs map[string]Struct
	Instancias map[string]StructInstance
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]Variable),
		Vectores:  make(map[string]Vector),
		Matrices: make(map[string]Matriz),
		Matrices3D: make(map[string]Matriz3D),
		funciones: make(map[string]Funcion),
		Structs: make(map[string]Struct),
		Instancias: make(map[string]StructInstance),
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

func (l *Visitor) agregarVariable(name string, tipo string, constante bool,value interface{}) interface{} {
	if _, ok := l.currentEnvironment.variables[name]; ok {
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", name)
	}

	if validateType(value,tipo) {
		fmt.Println("Se ha agregado una nueva variable")
		nuevaVariable := Variable{
			Name:  name,
			Type:  tipo,
			Constante: constante,
			Value: value,
		}
		l.currentEnvironment.variables[name] = nuevaVariable
	}
	return true
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
