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

//append
func (l *Visitor) VisitAppendvectorstmt(ctx *parser.AppendvectorstmtContext) interface{}  {
	id := ctx.ID().GetText()
	if vector, ok := l.currentEnvironment.Vectores[id]; ok {
		valor := l.Visit(ctx.Expr())

		//verificar que el valor sea del mismo tipo que el vector
		valorInsetar := determineType(valor)
		if valorInsetar != vector.Tipo {
			return fmt.Sprintf("Error: el valor a insertar no es del mismo tipo que el vector: %s", id)
		}else{
			//sustituir vector por el vector actualizado
			vector.Valores = append(vector.Valores, valor)
			l.currentEnvironment.Vectores[id] = vector
			return true
		}
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

// count
func (l *Visitor) VisitCountvectorstmt(ctx *parser.CountvectorstmtContext) interface{}  {
	id := ctx.ID().GetText()
	if vector, ok := l.currentEnvironment.Vectores[id]; ok {
		return int64(len(vector.Valores))
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}

//isEmpty
func (l *Visitor) VisitIsemptyvectorstmt(ctx *parser.IsemptyvectorstmtContext) interface{}  {
	id := ctx.ID().GetText()
	if vector, ok := l.currentEnvironment.Vectores[id]; ok {
		if len(vector.Valores) == 0 {
			return true
		}else{
			return false
		}
	}
	return fmt.Sprintf("Error vector no existente en el entorno actual: %s", id)
}