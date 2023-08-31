package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

// visit de asignaciones
func (l *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Buscar la variable en los entornos, comenzando por el actual y luego ascendiendo
	currentEnv := l.currentEnvironment

	for currentEnv != nil {
		// Verificar si la variable ya existe en el entorno
		if variableExistente, ok := currentEnv.variables[varName]; ok {
			if !variableExistente.Constante {
				// Determinar el tipo de la expresión y el tipo de la variable existente
				exprType := determineType(value)
				existingVarType := variableExistente.Type

				if exprType == existingVarType {

					variableActualizada := Variable{
						Name:  varName,
						Type:  exprType, // Asignar el tipo de la expresión
						Value: value,
					}

					currentEnv.variables[varName] = variableActualizada
					return true // La asignación fue exitosa
				} else {
					return fmt.Sprintf("Error de tipo en la asignacion para la variable: %s", varName)
				}
			} else {
				return fmt.Sprintf("Error la variable: %s es una constante", varName)
			}
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error no existe la variable: %s", varName)
}

//operadores de asignacion += -=

func (l *Visitor) VisitIncremento(ctx *parser.IncrementoContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Buscar la variable en los entornos, comenzando por el actual y luego ascendiendo
	currentEnv := l.currentEnvironment

	for currentEnv != nil {
		// Verificar si la variable ya existe en el entorno
		if variableExistente, ok := currentEnv.variables[varName]; ok {

			if !variableExistente.Constante {
				// Determinar el tipo de la expresión y el tipo de la variable existente
				exprType := determineType(value)
				existingVarType := variableExistente.Type

				if exprType == existingVarType {

					if existingVarType == "int" || existingVarType == "float" || existingVarType == "String" {
						if existingVarType == "int" {
							nuevoValor := value.(int64) + variableExistente.Value.(int64)
							variableActualizada := Variable{
								Name:  varName,
								Type:  exprType, // Asignar el tipo de la expresión
								Value: nuevoValor,
							}
							currentEnv.variables[varName] = variableActualizada
							return true // La asignación fue exitosa
						} else if existingVarType == "float" {
							nuevoValor := value.(float64) + variableExistente.Value.(float64)
							variableActualizada := Variable{
								Name:  varName,
								Type:  exprType, // Asignar el tipo de la expresión
								Value: nuevoValor,
							}
							currentEnv.variables[varName] = variableActualizada
							return true // La asignación fue exitosa
						} else if existingVarType == "String" {
							nuevoValor := value.(string) + variableExistente.Value.(string)
							variableActualizada := Variable{
								Name:  varName,
								Type:  exprType, // Asignar el tipo de la expresión
								Value: nuevoValor,
							}
							currentEnv.variables[varName] = variableActualizada
							return true // La asignación fue exitosa
						}
					} else {
						return fmt.Sprintf("Error: Operacion invalida")
					}
				} else {
					return fmt.Sprintf("Error de tipo en la asignación para: %s", varName) // La asignación falló debido a un error de tipo
				}
			} else {
				return fmt.Sprintf("Error la variable: %s es una constante", varName)
			}
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error Variable no encontrada: %s", varName) // La variable no se encontró en ningún entorno
}

func (l *Visitor) VisitDecremento(ctx *parser.DecrementoContext) interface{} {
	varName := ctx.ID().GetText()
	value := l.Visit(ctx.Expr())

	// Buscar la variable en los entornos, comenzando por el actual y luego ascendiendo
	currentEnv := l.currentEnvironment

	for currentEnv != nil {
		// Verificar si la variable ya existe en el entorno
		if variableExistente, ok := currentEnv.variables[varName]; ok {

			if !variableExistente.Constante {
				// Determinar el tipo de la expresión y el tipo de la variable existente
				exprType := determineType(value)
				existingVarType := variableExistente.Type

				if exprType == existingVarType {

					if existingVarType == "int" || existingVarType == "float" {
						if existingVarType == "int" {
							nuevoValor := -(value.(int64) - variableExistente.Value.(int64))
							fmt.Print(nuevoValor)
							variableActualizada := Variable{
								Name:  varName,
								Type:  exprType, // Asignar el tipo de la expresión
								Value: nuevoValor,
							}
							currentEnv.variables[varName] = variableActualizada
							return true // La asignación fue exitosa
						} else if existingVarType == "float" {
							nuevoValor := value.(float64) - variableExistente.Value.(float64)
							variableActualizada := Variable{
								Name:  varName,
								Type:  exprType, // Asignar el tipo de la expresión
								Value: nuevoValor,
							}
							currentEnv.variables[varName] = variableActualizada
							return true // La asignación fue exitosa
						}
					} else {
						return fmt.Sprintf("Error: Operacion invalida") // La asignación falló debido a un error de tipo
					}

				} else {
					return fmt.Sprintf("Error de tipo en la asignación para: %s", varName) // La asignación falló debido a un error de tipo
				}
			}else{
				return fmt.Sprintf("Error la variable: %s es una constante", varName)
			}
		}
		currentEnv = currentEnv.parent
	}
	return fmt.Sprintf("Error variable no encontrada: %s", varName) // La variable no se encontró en ningún entorno
}
