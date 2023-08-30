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

// vectores
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
			return fmt.Sprintf("Error: Índice fuera de rango: %d", index)
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error: Vector no encontrado: %s", vectorId)
}

// matrices
/*
5.2. Matrices
Las matrices en T-Swift nos permiten almacenar solamente datos de tipo primitivo, la
diferencia principal entre el vector y la matriz es que esta última organiza sus elementos en
n dimensiones y la manipulación de datos es con la notación [ ] además que su tamaño
no puede cambiar en tiempo de ejecución.
5.2.1. Creación de matrices
Las matrices en T-Swift pueden ser de 2 a n dimensiones pero solo de un tipo específico,
además su tamaño será constante y será definido durante su declaración.
Consideraciones:
● La declaración del tamaño puede ser explícita o en base a se definición.
● Si la declaración es explícita pero su definición no es acorde a esta declaración se
debe marcar como un error. Por lo tanto se deben verificar que la cantidad de
dimensiones sea acorde a la definida.
● La asignación y lectura valores se realizará con la notación [ ]
● Los índices de declaración comienzan a partir de 1
● Los índices de acceso comienzan a partir de 0
● Las matrices no van a cambiar su tamaño durante la ejecución.
● Si se hace un acceso con índices en fuera de rango se devuelve nil y se debe
notificar como un error.
● Si se declara una matriz con índices negativos o 0, será considerado un error
● El atributo count solo recibirá número enteros en forma de literales, no podrán ser
asignadas ni variables ni elementos de otras estructuras a este atributo.
*/

/*
gramatica
declmatrizstmt
    : VAR ID (DOSPUNTOS tipomatriz)? IG definicionmatriz
    ;

tipomatriz
    : CORCHIZQ tipomatriz CORCHDER
    | CORCHIZQ tipo CORCHDER
    ;

definicionmatriz
    : listavaloresmatriz
    ;

listavaloresmatriz
    : CORCHIZQ listavaloresmatriz2 CORCHDER
    ;

listavaloresmatriz2
    : listavaloresmatriz2 COMA listavaloresmatriz
    | listavaloresmatriz
    | listaexpresiones
    ;

*/

/*
	struct Matriz {
		Id      string
		Tipo    string
		Dimensiones []int64
		Valores [][]interface{}
	}
*/

/*
ejemplos de expresiones a evaluar
var mtx1 : [[Int]] = [[1,2,3],[4,5,6],[7,8,9]]

var mtx2 : [[[Int]]] = [[[1,2,3],[4,5,6],[7,8,9]],
[[10,11,12],[13,14,15],[16,17,18]],
[[19,20,21],[22,23,24],[25,26,2print(mtx1)7]]]

acceso a matriz
print(mtx1[0][0])
print(mtx2[0][0][0])
*/

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
				return fmt.Sprintf("Error de tipo en la declaración: %s", matrizId)
			}
		}
	}

	matrizNueva := Matriz{
		Id:      matrizId,
		Tipo:    tipo,
		Valores: valores,
	}
	l.currentEnvironment.Matrices[matrizId] = matrizNueva
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

//acceso matriz 
func (l *Visitor) VisitAccesomatriz2(ctx *parser.Accesomatriz2Context) interface{} {
	matrizId := ctx.ID().GetText()
	filaIdx := l.Visit(ctx.Expr(0)).(int64)
	columnaIdx := l.Visit(ctx.Expr(1)).(int64)

	// Acceder a la matriz y obtener el valor en la posición indicada
	matriz := l.currentEnvironment.Matrices[matrizId]
	if filaIdx >= 0 && filaIdx <  int64(len(matriz.Valores)) && columnaIdx >= 0 && columnaIdx < int64(len(matriz.Valores[0])) {
		valor := matriz.Valores[filaIdx][columnaIdx]
		return valor
	} else {
		// Manejar el caso de índices fuera de rango
		return "Error: Índices fuera de rango"
	}
}

//asignmatrizstmt
func (l *Visitor) VisitAsignmatrizstmt2(ctx *parser.Asignmatrizstmt2Context) interface{} {
	matrizId := ctx.ID().GetText()
	filaIdx := l.Visit(ctx.Expr(0)).(int64)
	columnaIdx := l.Visit(ctx.Expr(1)).(int64)
	valor := l.Visit(ctx.Expr(2))
	matriz := l.currentEnvironment.Matrices[matrizId]
	tipo := matriz.Tipo

	if !validateType(valor, tipo) {
		return fmt.Sprintf("Error de tipo en la asignacion: %s", matrizId)
	}

	if filaIdx >= 0 && filaIdx <  int64(len(matriz.Valores)) && columnaIdx >= 0 && columnaIdx < int64(len(matriz.Valores[0])) {
		matriz.Valores[filaIdx][columnaIdx] = valor
		l.currentEnvironment.Matrices[matrizId] = matriz
	} else {
		return "Error: Índices fuera de rango"
	}
	return true
}