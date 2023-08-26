package lenguaje

import (
	"fmt"
	"interprete/Parser"
)


//variables
func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", varName)
	}

	// Comprobar si el tipo de la expresión coincide con el tipo declarado
	if !validateType(value, declType) {
		//se  toma como error y obtiene el valor de nil para fines practicos
		nuevaVariable := Variable{
			Name:  varName,
			Type:  declType,
			Value: nil,
		}
		l.currentEnvironment.variables[varName] = nuevaVariable
		return fmt.Sprintf("Error de tipo en la declaración: %s", varName)
	}

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  declType,
		Value: value,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", varName)
	}

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  declType,
		Value: nil,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true

}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", varName)
	}

	// Determinar el tipo de la expresión
	valueType := determineType(value)
	fmt.Println(valueType)

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  valueType, // Asignar el tipo de la expresión
		Value: value,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	return true
}

//vectores
func (l *Visitor) VisitDeclvectorstmt(ctx *parser.DeclvectorstmtContext) interface{} {
	varName := ctx.ID().GetText()
	tipo := l.Visit(ctx.Tipo()).(string)
	valores := l.Visit(ctx.Defvectorstmt()).([]interface{})

	//verificar tipo de cada uno de los valores
	for _, valor := range valores {
		if !validateType(valor, tipo) {
			return fmt.Sprintf("Error de tipo en la declaración: %s", varName)
		}
	}
	l.agregarVector(varName, tipo, valores)
	return true
}

//definicicion de vector
func (l *Visitor) VisitDefVector(ctx *parser.DefVectorContext) interface{} {
	if ctx.Listaexpresiones() != nil {
		valores := l.Visit(ctx.Listaexpresiones()).([]interface{})
		return valores
	}else{
		valores := []interface{}{}
		return valores
	}
}

func (l *Visitor) VisitDefVectorID(ctx *parser.DefVectorIDContext) interface{} {
	return nil
}

//lista de valores
func (l *Visitor) VisitListaexpresiones(ctx *parser.ListaexpresionesContext) interface{} {
	var list []interface{}
	for _, expr := range ctx.AllExpr() {
		list = append(list, l.Visit(expr))
	}
	return list
}

//acceso vector
func (l *Visitor) VisitAccesovectorstmt(ctx *parser.AccesovectorstmtContext) interface{} {
	vectorId := ctx.ID().GetText()
	index := l.Visit(ctx.Expr()).(int64)

	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		vector, ok := currentEnv.Vectores[vectorId]
		if ok {
			if index >= 0 && index < int64(len(vector.Valores)) {
				return vector.Valores[index]
			}
			return fmt.Sprintf("Error: Índice fuera de rango: %d", index)
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error: Vector no encontrado: %s", vectorId)
}