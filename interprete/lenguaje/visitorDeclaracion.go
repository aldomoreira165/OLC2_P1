package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// variables
func (l *Visitor) VisitTypedDeclstmt(ctx *parser.TypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)
	value := l.Visit(ctx.Expr())

	// Comprobar si el tipo de la expresión coincide con el tipo declarado
	if !validateType(value, declType) {
		//se  toma como error y obtiene el valor de nil para fines practicos
		nuevaVariable := Variable{
			Name:  varName,
			Type:  declType,
			Value: nil,
		}
		l.currentEnvironment.variables[varName] = nuevaVariable
		l.errores.InsertarError("Error de tipo en la declaración: " + varName, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error de tipo en la declaración: %s", varName)
	}

	if ctx.LET() != nil {
		variable := l.agregarVariable(varName, declType, true, value, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.simbolos.InsertarSimbolo(varName, "Variable",declType, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return variable
	} else {
		variable := l.agregarVariable(varName, declType, false, value, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.simbolos.InsertarSimbolo(varName, "Variable",declType, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return variable
	}
}

func (l *Visitor) VisitOptionalTypedDeclstmt(ctx *parser.OptionalTypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	declType := l.Visit(ctx.Tipo()).(string)

	// Verificar si la variable ya existe en el entorno
	if _, ok := l.currentEnvironment.variables[varName]; ok {
		l.errores.InsertarError("Error variable ya existente en el entorno actual: " + varName, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error variable ya existente en el entorno actual: %s", varName)
	}

	// Crear una instancia de Variable y almacenarla en el entorno actual
	nuevaVariable := Variable{
		Name:  varName,
		Type:  declType,
		Constante: false,
		Value: nil,
	}

	// Almacenar la información de la declaración en el entorno
	l.currentEnvironment.variables[varName] = nuevaVariable
	l.simbolos.InsertarSimbolo(varName, "Variable",declType, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return true

}

func (l *Visitor) VisitUntypedDeclstmt(ctx *parser.UntypedDeclstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())
	valueType := determineType(value)
	
	if ctx.LET() != nil {
		variable := l.agregarVariable(varName, valueType, true, value, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.simbolos.InsertarSimbolo(varName, "Variable",valueType, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return variable
	} else {
		variable := l.agregarVariable(varName, valueType, false, value, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		l.simbolos.InsertarSimbolo(varName, "Variable",valueType, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return variable
	}
}

// vectores
func (l *Visitor) VisitDeclvectorstmt(ctx *parser.DeclvectorstmtContext) interface{} {
	varName := ctx.ID().GetText()
	tipo := l.Visit(ctx.Tipo()).(string)
	valores := l.Visit(ctx.Defvectorstmt()).([]interface{})

	//verificar tipo de cada uno de los valores
	for _, valor := range valores {
		if !validateType(valor, tipo) {
			l.errores.InsertarError("Error de tipo en la declaración: " + varName, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error de tipo en la declaración: %s", varName)
		}
	}
	l.agregarVector(varName, tipo, valores, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	l.simbolos.InsertarSimbolo(varName, "Vector",tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return true
}

// definicicion de vector
func (l *Visitor) VisitDefVector(ctx *parser.DefVectorContext) interface{} {
	if ctx.Listaexpresiones() != nil {
		valores := l.Visit(ctx.Listaexpresiones()).([]interface{})
		return valores
	} else {
		valores := []interface{}{}
		return valores
	}
}

// asignar a vector los valores de otro vector
func (l *Visitor) VisitDefVectorID(ctx *parser.DefVectorIDContext) interface{} {
	id := ctx.ID().GetText()
	if vector, ok := l.currentEnvironment.Vectores[id]; ok {
		return vector.Valores
	}
	l.errores.InsertarError("Error vector no existente en el entorno actual: " + id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// lista de valores
func (l *Visitor) VisitListaexpresiones(ctx *parser.ListaexpresionesContext) interface{} {
	var list []interface{}
	for _, expr := range ctx.AllExpr() {
		list = append(list, l.Visit(expr))
	}
	fmt.Println("lista de valores: ", list)
	return list
}

// acceso vector
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
			l.errores.InsertarError("Error: Índice fuera de rango: " + vectorId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error: Índice fuera de rango: %d", index)
		}
		currentEnv = currentEnv.parent
	}
	l.errores.InsertarError("Error: Vector no encontrado: " + vectorId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error: Vector no encontrado: %s", vectorId)
}

// matrices
func (l *Visitor) VisitDeclmatrizstmt2(ctx *parser.Declmatrizstmt2Context) interface{} {
	matrizId := ctx.ID().GetText()
	tipo := l.Visit(ctx.Tipomatriz()).(string)
	valores := l.Visit(ctx.Listavaloresmatriz()).([][]interface{})
	fmt.Println("matrizId:", matrizId)
	fmt.Println("tipo:", tipo)
	fmt.Println("valores:", valores)

	for _, fila := range valores {
		for _, valor := range fila {
			if !validateType(valor, tipo) {
				l.errores.InsertarError("Error de tipo en la declaración: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return fmt.Sprintf("Error de tipo en la declaración: %s", matrizId)
			}
		}
	}

	matrizNueva := Matriz{
		Id:      matrizId,
		Tipo:    tipo,
		Dimension: 2,
		Valores: valores,
	}
	l.currentEnvironment.Matrices[matrizId] = matrizNueva
	l.simbolos.InsertarSimbolo(matrizId, "Matriz",tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return true
}

func (l *Visitor) VisitTipomatriz2(ctx *parser.Tipomatriz2Context) interface{} {
	return ctx.Tipo().GetText()
}

func (l *Visitor) VisitListavaloresmatriz2(ctx *parser.Listavaloresmatriz2Context) interface{} {
	matrizValores := [][]interface{}{}

	for _, filaExpresiones := range ctx.AllListaexpresiones() {
		fila := l.Visit(filaExpresiones).([]interface{})
		matrizValores = append(matrizValores, fila)
	}

	return matrizValores
}

// acceso matriz
func (l *Visitor) VisitAccesomatriz2(ctx *parser.Accesomatriz2Context) interface{} {
	matrizId := ctx.ID().GetText()
	filaIdx := l.Visit(ctx.Expr(0)).(int64)
	columnaIdx := l.Visit(ctx.Expr(1)).(int64)

	// si no se encuentra la lista en matrices se busca en matrices3D, si se encuentra en 3d se devuelve la lista de valores []interface{}

	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		if matriz, ok := currentEnv.Matrices[matrizId]; ok {
			if filaIdx >= 0 && filaIdx < int64(len(matriz.Valores)) && columnaIdx >= 0 && columnaIdx < int64(len(matriz.Valores[0])) {
				valor := matriz.Valores[filaIdx][columnaIdx]
				return valor
			} else {
				l.errores.InsertarError("Error: Índices fuera de rango", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return "Error: Índices fuera de rango"
			}
		} else if matriz3D, ok := currentEnv.Matrices3D[matrizId]; ok {
			if filaIdx >= 0 && filaIdx < int64(len(matriz3D.Valores)) && columnaIdx >= 0 && columnaIdx < int64(len(matriz3D.Valores[0])) {
				valor := matriz3D.Valores[filaIdx][columnaIdx]
				return valor
			} else {
				l.errores.InsertarError("Error: Índices fuera de rango", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return "Error: Índices fuera de rango"
			}
		}
		currentEnv = currentEnv.parent
	}
	l.errores.InsertarError("Error: Matriz no encontrada: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error: Matriz no encontrada: %s", matrizId)
}

// asignmatrizstmt
func (l *Visitor) VisitAsignmatrizstmt2(ctx *parser.Asignmatrizstmt2Context) interface{} {
	matrizId := ctx.ID().GetText()
	filaIdx := l.Visit(ctx.Expr(0)).(int64)
	columnaIdx := l.Visit(ctx.Expr(1)).(int64)
	valor := l.Visit(ctx.Expr(2))
	matriz := l.currentEnvironment.Matrices[matrizId]
	tipo := matriz.Tipo

	if !validateType(valor, tipo) {
		l.errores.InsertarError("Error de tipo en la asignacion: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error de tipo en la asignacion: %s", matrizId)
	}

	if filaIdx >= 0 && filaIdx < int64(len(matriz.Valores)) && columnaIdx >= 0 && columnaIdx < int64(len(matriz.Valores[0])) {
		matriz.Valores[filaIdx][columnaIdx] = valor
		l.currentEnvironment.Matrices[matrizId] = matriz
	} else {
		l.errores.InsertarError("Error: Índices fuera de rango", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return "Error: Índices fuera de rango"
	}
	return true
}

//matrices de 3 dimensiones

func (l *Visitor) VisitDeclmatrizstmt3(ctx *parser.Declmatrizstmt3Context) interface{} {
	matrizId := ctx.ID().GetText()
	tipo := l.Visit(ctx.Tipomatriz()).(string)
	valores := l.Visit(ctx.Listavaloresmatriz3()).([][][]interface{})
	fmt.Println("matrizId:", matrizId)
	fmt.Println("tipo:", tipo)
	fmt.Println("valores:", valores)

	for _, nivelMatriz := range valores {
		for _, fila := range nivelMatriz {
			for _, valor := range fila {
				if !validateType(valor, tipo) {
					l.errores.InsertarError("Error de tipo en la declaración: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
					return fmt.Sprintf("Error de tipo en la declaración: %s", matrizId)
				}
			}
		}
	}

	matrizNueva := Matriz3D{
		Id:      matrizId,
		Tipo:    tipo,
		Dimension: 3,
		Valores: valores,
	}
	l.currentEnvironment.Matrices3D[matrizId] = matrizNueva
	l.simbolos.InsertarSimbolo(matrizId, "Matriz",tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return true
}

func (l *Visitor) VisitTipomatriz3(ctx *parser.Tipomatriz3Context) interface{} {
	return ctx.Tipo().GetText()
}

func (l *Visitor) VisitListavaloresmatriz3(ctx *parser.Listavaloresmatriz3Context) interface{} {
	matrizValores := [][][]interface{}{}

	for _, nivelMatriz := range ctx.AllListavaloresmatriz() {
		nivel := l.Visit(nivelMatriz).([][]interface{})
		matrizValores = append(matrizValores, nivel)
	}

	return matrizValores
}

func (l *Visitor) VisitAccesomatriz3(ctx *parser.Accesomatriz3Context) interface{} {
	matrizId := ctx.ID().GetText()
	nivelIdx := l.Visit(ctx.Expr(0)).(int64)
	filaIdx := l.Visit(ctx.Expr(1)).(int64)
	columnaIdx := l.Visit(ctx.Expr(2)).(int64)

	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		if matriz3D, ok := currentEnv.Matrices3D[matrizId]; ok {
			if nivelIdx >= 0 && nivelIdx < int64(len(matriz3D.Valores)) && filaIdx >= 0 && filaIdx < int64(len(matriz3D.Valores[0])) && columnaIdx >= 0 && columnaIdx < int64(len(matriz3D.Valores[0][0])) {
				valor := matriz3D.Valores[nivelIdx][filaIdx][columnaIdx]
				return valor
			} else {
				l.errores.InsertarError("Error: Índices fuera de rango", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return "Error: Índices fuera de rango"
			}
		}
		currentEnv = currentEnv.parent
	}
	l.errores.InsertarError("Error: Matriz no encontrada: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error: Matriz no encontrada: %s", matrizId)
}


func (l *Visitor) VisitAsignmatrizstmt3(ctx *parser.Asignmatrizstmt3Context) interface{} {
	matrizId := ctx.ID().GetText()
	nivelIdx := l.Visit(ctx.Expr(0)).(int64)
	filaIdx := l.Visit(ctx.Expr(1)).(int64)
	columnaIdx := l.Visit(ctx.Expr(2)).(int64)
	valor := l.Visit(ctx.Expr(3))
	matriz := l.currentEnvironment.Matrices3D[matrizId]
	tipo := matriz.Tipo

	if !validateType(valor, tipo) {
		l.errores.InsertarError("Error de tipo en la asignacion: " + matrizId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error de tipo en la asignacion: %s", matrizId)
	}

	if nivelIdx >= 0 && nivelIdx < int64(len(matriz.Valores)) && filaIdx >= 0 && filaIdx < int64(len(matriz.Valores[0])) && columnaIdx >= 0 && columnaIdx < int64(len(matriz.Valores[0][0])) {
		matriz.Valores[nivelIdx][filaIdx][columnaIdx] = valor
		l.currentEnvironment.Matrices3D[matrizId] = matriz
	} else {
		l.errores.InsertarError("Error: Índices fuera de rango", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return "Error: Índices fuera de rango"
	}
	return true
}

