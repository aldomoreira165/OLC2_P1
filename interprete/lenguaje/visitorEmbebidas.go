package lenguaje

import (
	"fmt"
	"interprete/Parser"
	"math"
	"strconv"
)

// visit del print
func (l *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {
	returnValue := l.Visit(ctx.Expr())
	stringValue := fmt.Sprint(returnValue)
	fmt.Println(stringValue + "\n")
	return stringValue + "\n"
}

// visit del int
func (l *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
	value := l.Visit(ctx.Expr())
	switch val := value.(type) {
	case string:
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return intVal
		}
		floatVal, err := strconv.ParseFloat(val, 64)
		if err == nil {
			intVal := int64(floatVal)
			return intVal
		}
		fmt.Println("Error: No se puede convertir a int64:", val)
		return nil
	case float64:
		intVal := int64(math.Trunc(val))
		intRet := interface{}(intVal)
		return intRet
	default:
		fmt.Println("Error: Tipo no válido para conversión a int64")
		return nil
	}
}

// visit del float
func (l *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) interface{} {
	value := l.Visit(ctx.Expr())
	switch val := value.(type) {
	case string:
		floatVal, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return floatVal
		}
		fmt.Println("Error: No se puede convertir a float64:", val)
		return nil
	case int64:
		floatVal := float64(val)
		floatRet := interface{}(floatVal)
		return floatRet
	default:
		fmt.Println("Error: Tipo no válido para conversión a float64")
		return nil
	}
}

// visit del string
func (l *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) interface{} {
	value := l.Visit(ctx.Expr())
	switch val := value.(type) {
	case string:
		return val
	case int64:
		stringVal := strconv.FormatInt(val, 10)
		return stringVal
	case float64:
		stringVal := strconv.FormatFloat(val, 'f', -1, 64)
		fmt.Println("tipoooo str:", determineType(stringVal))
		return stringVal
	case bool:
		stringVal := strconv.FormatBool(val)
		fmt.Println("tipoooo str:", determineType(stringVal))
		return stringVal
	default:
		fmt.Println("Error: Tipo no válido para conversión a string")
		return nil
	}
}