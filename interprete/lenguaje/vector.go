package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

type Vector struct {
	Id      string
	Tipo    string
	Valores []interface{}
}

// append
func (l *Visitor) VisitAppendvectorstmt(ctx *parser.AppendvectorstmtContext) interface{} {
	id := ctx.ID().GetText()
	currentEnvironment := l.currentEnvironment

	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			valor := l.Visit(ctx.Expr())

			//verificar que el valor sea del mismo tipo que el vector
			valorInsetar := determineType(valor)
			if valorInsetar != vector.Tipo {
				return fmt.Sprintf("Error: el valor a insertar no es del mismo tipo que el vector: %s", id)
			} else {
				//sustituir vector por el vector actualizado
				vector.Valores = append(vector.Valores, valor)
				currentEnvironment.Vectores[id] = vector
				return true
			}
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// asignacion o modificacion de posicion de vector
func (l *Visitor) VisitAsignvectorstmt(ctx *parser.AsignvectorstmtContext) interface{} {
	id := ctx.ID().GetText()
	currentEnvironment := l.currentEnvironment
	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			valor := l.Visit(ctx.Expr(0))
			if valor.(int64) < int64(len(vector.Valores)) {
				vector.Valores[valor.(int64)] = l.Visit(ctx.Expr(1))
				currentEnvironment.Vectores[id] = vector
				return true
			} else {
				return fmt.Sprintf("Error: el indice esta fuera de rango: %s", id)
			}
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// remove last
func (l *Visitor) VisitRemovelastvectorstmt(ctx *parser.RemovelastvectorstmtContext) interface{} {
	id := ctx.ID().GetText()

	currentEnvironment := l.currentEnvironment
	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			if len(vector.Valores) > 0 {
				vector.Valores = vector.Valores[:len(vector.Valores)-1]
				currentEnvironment.Vectores[id] = vector
				return true
			} else {
				return fmt.Sprintf("Error: el vector esta vacio: %s", id)	
			}
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// remove at
func (l *Visitor) VisitRemoveatvectorstmt(ctx *parser.RemoveatvectorstmtContext) interface{} {
	id := ctx.ID().GetText()

	currentEnvironment := l.currentEnvironment
	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			valor := l.Visit(ctx.Expr())
			if valor.(int64) < int64(len(vector.Valores)) {
				vector.Valores = append(vector.Valores[:valor.(int64)], vector.Valores[valor.(int64)+1:]...)
				currentEnvironment.Vectores[id] = vector
				return true
			} else {
				return fmt.Sprintf("Error: el indice esta fuera de rango: %s", id)
			}
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// count
func (l *Visitor) VisitCountvectorstmt(ctx *parser.CountvectorstmtContext) interface{} {
	id := ctx.ID().GetText()
	
	currentEnvironment := l.currentEnvironment
	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			return int64(len(vector.Valores))
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// isEmpty
func (l *Visitor) VisitIsemptyvectorstmt(ctx *parser.IsemptyvectorstmtContext) interface{} {
	id := ctx.ID().GetText()

	currentEnvironment := l.currentEnvironment
	for currentEnvironment != nil {
		if vector, ok := currentEnvironment.Vectores[id]; ok {
			if len(vector.Valores) == 0 {
				return true
			} else {
				return false
			}
		}
		currentEnvironment = currentEnvironment.parent
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}
