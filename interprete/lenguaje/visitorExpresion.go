package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	leftValue := l.Visit(ctx.GetLeft())
	fmt.Println("valor izquierdo ", leftValue)
	rightValue := l.Visit(ctx.GetRight())
	fmt.Println("valor derecho ", rightValue)
	op := ctx.GetOp().GetText()

	leftType := determineType(leftValue)
	fmt.Println("tipo izquierdo ", leftType)
	rightType := determineType(rightValue)
	fmt.Println("tipo derecho ", rightType)

	switch op {
	case "+":
		if leftType == "int" && rightType == "int" || rightType == "int" && leftType == "int" {
			return leftValue.(int64) + rightValue.(int64)
		} else if leftType == "int" && rightType == "float" {
			return float64(leftValue.(int64)) + rightValue.(float64)
		} else if leftType == "float" && rightType == "int" {
			return leftValue.(float64) + float64(rightValue.(int64))
		} else if leftType == "float" && rightType == "float" || rightType == "float" && leftType == "float" {
			return leftValue.(float64) + rightValue.(float64)
		} else if leftType == "String" && rightType == "String" || rightType == "String" && leftType == "String" {
			return leftValue.(string) + rightValue.(string)
		} else if leftType == "String" && rightType == "character" || rightType == "character" && leftType == "String" {
			return leftValue.(string) + rightValue.(string)
		} else if leftType == "character" && rightType == "character" || rightType == "character" && leftType == "character" {
			return leftValue.(string) + rightValue.(string)
		}
	case "-":
		if leftType == "int" && rightType == "int" || rightType == "int" && leftType == "int" {
			return leftValue.(int64) - rightValue.(int64)
		} else if leftType == "int" && rightType == "float" {
			return float64(leftValue.(int64)) - rightValue.(float64)
		} else if leftType == "float" && rightType == "int" {
			return leftValue.(float64) - float64(rightValue.(int64))
		} else if leftType == "float" && rightType == "float" || rightType == "float" && leftType == "float" {
			return leftValue.(float64) - rightValue.(float64)
		}
	case "*":
		if leftType == "int" && rightType == "int" || rightType == "int" && leftType == "int" {
			return leftValue.(int64) * rightValue.(int64)
		} else if leftType == "int" && rightType == "float" {
			return float64(leftValue.(int64)) * rightValue.(float64)
		} else if leftType == "float" && rightType == "int" {
			return leftValue.(float64) * float64(rightValue.(int64))
		} else if leftType == "float" && rightType == "float" || rightType == "float" && leftType == "float" {
			return leftValue.(float64) * rightValue.(float64)
		}
	case "/":
		//verificar que no se divida entre 0
		if leftType == "int" && rightType == "int" || rightType == "int" && leftType == "int" {
			return leftValue.(int64) / rightValue.(int64)
		} else if leftType == "int" && rightType == "float" {
			return float64(leftValue.(int64)) / rightValue.(float64)
		} else if leftType == "float" && rightType == "int" {
			return leftValue.(float64) / float64(rightValue.(int64))
		} else if leftType == "float" && rightType == "float" || rightType == "float" && leftType == "float" {
			return leftValue.(float64) / rightValue.(float64)
		}
	case "%":
		if leftType == "int" && rightType == "int" || rightType == "int" && leftType == "int" {
			return leftValue.(int64) % rightValue.(int64)
		}
	case "==":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) == rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) == rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "bool" && rightType == "bool" {
			if leftValue.(bool) == rightValue.(bool) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) == rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case "!=":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) != rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) != rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "bool" && rightType == "bool" {
			if leftValue.(bool) != rightValue.(bool) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) != rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case ">":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) > rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) > rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) > rightValue.(string) {
				return true
			} else {
				return false
			}
		} else if leftType == "character" && rightType == "character" {
			if leftValue.(string) > rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case ">=":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) >= rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) >= rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) >= rightValue.(string) {
				return true
			} else {
				return false
			}
		} else if leftType == "character" && rightType == "character" {
			if leftValue.(string) >= rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case "<":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) < rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) < rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) < rightValue.(string) {
				return true
			} else {
				return false
			}
		} else if leftType == "character" && rightType == "character" {
			if leftValue.(string) < rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case "<=":
		if leftType == "int" && rightType == "int" {
			if leftValue.(int64) <= rightValue.(int64) {
				return true
			} else {
				return false
			}
		} else if leftType == "float" && rightType == "float" {
			if leftValue.(float64) <= rightValue.(float64) {
				return true
			} else {
				return false
			}
		} else if leftType == "String" && rightType == "String" {
			if leftValue.(string) <= rightValue.(string) {
				return true
			} else {
				return false
			}
		} else if leftType == "character" && rightType == "character" {
			if leftValue.(string) <= rightValue.(string) {
				return true
			} else {
				return false
			}
		}
	case "&&":
		if leftType == "bool" && rightType == "bool" {
			valorA := leftValue.(bool)
			valorB := rightValue.(bool)

			if valorA && valorB {
				return true
			} else if valorA && !valorB {
				return false
			} else if !valorA && valorB {
				return false
			} else if !valorA && !valorB{
				return false
			}
		}

	case "||":
		if leftType == "bool" && rightType == "bool" {
			valorA := leftValue.(bool)
			valorB := rightValue.(bool)

			if valorA && valorB {
				return true
			} else if valorA && !valorB {
				return true
			} else if !valorA && valorB {
				return true
			} else if !valorA && !valorB{
				return false
			}
		}
	}

	//error, verificar como manejarlos
	l.errores.InsertarError("Error: operacion no valida", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error: operacion no valida")
}

func (l *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return l.Visit(ctx.Expr())
}

func (l *Visitor) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	numStr := ctx.GetText()

	intVal, intErr := strconv.ParseInt(numStr, 10, 64)
	if intErr == nil {
		return intVal
	}

	floatVal, floatErr := strconv.ParseFloat(numStr, 64)
	if floatErr == nil {
		return floatVal
	}
	return nil
}

func (l *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()
	// Buscar la variable en los entornos, comenzando por el actual y luego ascendiendo
	currentEnv := l.currentEnvironment
	for currentEnv != nil {
		value, ok := currentEnv.variables[id]
		if ok {
			return value.Value
		}
		currentEnv = currentEnv.parent
	}
	//buscar en la lista de vectores del entorno
	currentEnv = l.currentEnvironment
	for currentEnv != nil {
		vector, ok := currentEnv.Vectores[id]
		if ok {
			return vector.Valores
		}
		currentEnv = currentEnv.parent
	}

	//buscar en la lista de matrices
	currentEnv = l.currentEnvironment
	for currentEnv != nil {
		matriz, ok := currentEnv.Matrices[id]
		if ok {
			return matriz.Valores
		}
		currentEnv = currentEnv.parent
	}

	//buscar en lista de matrices 3d
	currentEnv = l.currentEnvironment
	for currentEnv != nil {
		matriz3d, ok := currentEnv.Matrices3D[id]
		if ok {
			return matriz3d.Valores
		}
		currentEnv = currentEnv.parent
	}

	//buscar en lista de structs
	currentEnv = l.currentEnvironment
	for currentEnv != nil {
		structInstance, ok := currentEnv.Instancias[id]
		if ok {
			return structInstance
		}
		currentEnv = currentEnv.parent
	}

	// Devolver un mensaje de error en lugar de lanzar una excepción
	l.errores.InsertarError("Error variable no encontrada: " + id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error variable no encontrada: %s", id)
}

func (l *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return value
}

func (l *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")

	if value == "true" {
		return true
	} else if value == "false" {
		return false
	}

	return false
}

func (l *Visitor) VisitUnariaExpr(ctx *parser.UnariaExprContext) interface{} {
	exp := l.Visit(ctx.Expr())
	tipo := determineType(exp)

	if tipo == "int" {
		valor, _ := exp.(int64)
		return -valor
	} else if tipo == "float" {
		valor, _ := exp.(float64)
		return -valor
	}
	return false
}

//not expr
func (l *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	exp := l.Visit(ctx.Expr())

	if exp == true {
		return false
	} else if exp == false {
		return true
	}
	return false
}

func (l *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return nil
}

func (l *Visitor) VisitAccFuncExpr(ctx *parser.AccFuncExprContext) interface{} {
	return l.Visit(ctx.Accfuncstm())
}

func (l *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return l.Visit(ctx.Intstmt())
}

func (l *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return l.Visit(ctx.Floatstmt())
}

func (l *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	return l.Visit(ctx.Stringstmt())
}

func (l *Visitor) VisitAccesoVectorExpr(ctx *parser.AccesoVectorExprContext) interface{} {
	return l.Visit(ctx.Accesovectorstmt())
}

func (l *Visitor) VisitCountVectorExpr(ctx *parser.CountVectorExprContext) interface{} {
	return l.Visit(ctx.Countvectorstmt())
}

func (l *Visitor) VisitIsEmptyVectorExpr(ctx *parser.IsEmptyVectorExprContext) interface{} {
	return l.Visit(ctx.Isemptyvectorstmt())
}

func (l *Visitor) VisitAccesoMatrizExpr(ctx *parser.AccesoMatrizExprContext) interface{} {
	return l.Visit(ctx.Accesomatriz())
}

func (l *Visitor) VisitAccesoStructExpr(ctx *parser.AccesoStructExprContext) interface{} {
	return l.Visit(ctx.Accesostructstmt())
}

func (l *Visitor) VisitAccesoValorStructExpr(ctx *parser.AccesoValorStructExprContext) interface{} {
	return l.Visit(ctx.Valor_struct_expr())
}

func (l *Visitor) VisitAccesoVectorStructExpr(ctx *parser.AccesoVectorStructExprContext) interface{} {
	return l.Visit(ctx.Acceso_vector_struct_stmt())
}

