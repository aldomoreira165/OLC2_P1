package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	leftValue := l.Visit(ctx.GetLeft())
	rightValue := l.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	leftType := determineType(leftValue)
	rightType := determineType(rightValue)

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
	// Devolver un mensaje de error en lugar de lanzar una excepción
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

func (l *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return nil
}
