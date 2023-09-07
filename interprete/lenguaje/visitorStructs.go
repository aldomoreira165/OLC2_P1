package lenguaje

import (
	"fmt"
	"interprete/Parser"
)

func (l *Visitor) VisitDefStruct(ctx *parser.DefStructContext) interface{} {
	structName := ctx.ID().GetText()
	structAttributes := []Attribute{}

	for _, attrCtx := range ctx.AllLista_atributos() {
		attribute := l.Visit(attrCtx).(Attribute)
		structAttributes = append(structAttributes, attribute)
	}
	l.currentEnvironment.Structs[structName] = Struct{
		Name:       structName,
		Attributes: structAttributes,
	}
	fmt.Println("lista de structs: ", l.currentEnvironment.Structs)
	return nil
}

func (l *Visitor) VisitAtributoStruct(ctx *parser.AtributoStructContext) interface{} {
	isMutable := ctx.VAR() != nil
	attributeName := ctx.ID(0).GetText()
	var attributeType string
	var attributeExpr interface{}

	//verificar si es un dato primitivo u otro struct
	if ctx.Tipo() != nil {
		attributeType = ctx.Tipo().GetText()
	} else if ctx.ID(1) != nil {
		attributeType = ctx.ID(1).GetText()
	}

	if ctx.Expr() != nil {
		attributeExpr = l.Visit(ctx.Expr())
	} else {
		attributeExpr = nil
	}

	// Crear una nueva instancia de Attribute
	attribute := Attribute{
		Name:       attributeName,
		IsMutable:  isMutable,
		Type:       attributeType,
		Expression: attributeExpr,
	}
	return attribute
}

func (l *Visitor) VisitStructExpr(ctx *parser.StructExprContext) interface{} {
	instanceName := ctx.ID(0).GetText()
	structInstance := l.Visit(ctx.Valor_struct_expr()).(StructInstance)

	//verificar que no hayan atributos nil
	for key, value := range structInstance.Attributes {
		if value.Data == nil {
			l.errores.InsertarError("Error: el atributo "+key+" no puede ser nil", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error: el atributo %s no puede ser nil", key)
		}
	}

	//verificar que no se haya asignado valor a una propiedad let o inmutable que ya tenia un valor definido en el struct
	if structDef, ok := l.currentEnvironment.Structs[structInstance.StructName]; ok {
		for _, attr := range structDef.Attributes {
			if !attr.IsMutable {
				if attr.Expression != nil {
					fmt.Println("exp", attr.Expression)
					if instance, ok := structInstance.Attributes[attr.Name]; ok {
						if instance.Asignado {
							l.errores.InsertarError("Error: el atributo "+attr.Name+" del struct "+structInstance.StructName+" es inmutable y ya tiene un valor definido", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
							return fmt.Sprintf("Error: el atributo %s del struct %s es inmutable y ya tiene un valor definido", attr.Name, structInstance.StructName)
						}
					}
				}
			}
		}
	}

	structInstance.StructName = instanceName
	l.currentEnvironment.Instancias[instanceName] = structInstance
	return nil
}

func (l *Visitor) VisitValorStructExpr(ctx *parser.ValorStructExprContext) interface{} {
	structName := ctx.ID().GetText()
	attributes := l.Visit(ctx.L_dupla()).([]Dupla)
	if structDef, ok := l.currentEnvironment.Structs[structName]; ok {

		//verificar que attributes tenga los mismos atributos que structDef
		if len(attributes) != len(structDef.Attributes) {
			//en caso no agregar el atributo faltante de structDef y el valor que tenga por defecto
			for i := 0; i < len(structDef.Attributes); i++ {
				//buscar si el atributo ya existe en attributes
				encontrado := false
				for _, attr := range attributes {
					if attr.AttributeName == structDef.Attributes[i].Name {
						encontrado = true
						break
					}
				}
				if !encontrado {
					attributes = append(attributes, Dupla{
						AttributeName:     structDef.Attributes[i].Name,
						AttributeAsignado: false,
						AttributeType:     structDef.Attributes[i].Type,
						AttributeValue:    structDef.Attributes[i].Expression,
					})
				}
			}
		}

		//retornando data con los atributos
		dataInstance := StructInstance{
			StructName: structName,
			Attributes: make(map[string]infoDataInstance),
		}

		for _, attr := range attributes {

			infoDataInstance := infoDataInstance{
				Asignado: attr.AttributeAsignado,
				Data:     attr.AttributeValue,
			}
			dataInstance.Attributes[attr.AttributeName] = infoDataInstance
		}

		return dataInstance
	} else {
		l.errores.InsertarError("Error: el struct "+structName+" no está definido", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: el struct %s no está definido", structName)
	}
}

// #StructExprID
func (l *Visitor) VisitStructExprID(ctx *parser.StructExprIDContext) interface{} {
	structName := ctx.ID(0).GetText()
	id := ctx.ID(1).GetText()

	//buscar el id en las instancias de structs y asignarle el valor
	if structInstance, ok := l.currentEnvironment.Instancias[id]; ok {
		//crear una nueva instancia de StructInstance
		newStructInstance := StructInstance{
			StructName: structName,
			Attributes: make(map[string]infoDataInstance),
		}

		//agregar los atributos de structInstance a newStructInstance
		for key, value := range structInstance.Attributes {
			newStructInstance.Attributes[key] = value
		}

		l.currentEnvironment.Instancias[structName] = newStructInstance
	} else {
		l.errores.InsertarError("Error: el struct "+structName+" no está definido", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: el struct %s no está definido", structName)
	}
	return nil
}

func (l *Visitor) VisitDuplastruct(ctx *parser.DuplastructContext) interface{} {
	duplas := []Dupla{}

	for i := 0; i < len(ctx.AllID()); i++ {
		attributeName := ctx.ID(i).GetText()
		attributeValue := l.Visit(ctx.Expr(i))
		attriType := determineType(attributeValue)
		fmt.Println("attributeName: ", attributeName, "attributeValue: ", attributeValue, "attriType: ", attriType)

		// Agregar la dupla actual al arreglo de duplas
		duplas = append(duplas, Dupla{
			AttributeName:     attributeName,
			AttributeAsignado: true,
			AttributeType:     "",
			AttributeValue:    attributeValue,
		})
	}
	return duplas
}

func (l *Visitor) VisitAccesoStruct(ctx *parser.AccesoStructContext) interface{} {
	structID := ctx.ID(0).GetText()
	attributeNames := ctx.AllID()[1:]

	// Verificar si el struct existe en el entorno actual
	if structInstance, ok := l.currentEnvironment.Instancias[structID]; ok {
		var currentValue interface{} = structInstance

		for _, attributeNameCtx := range attributeNames {
			attributeName := attributeNameCtx.GetText()

			// Verificar si el valor actual es una instancia de un struct
			if currentStructInstance, ok := currentValue.(StructInstance); ok {
				// Verificar si el struct tiene el atributo
				if attributeValue, ok := currentStructInstance.Attributes[attributeName]; ok {
					currentValue = attributeValue.Data
				} else {
					l.errores.InsertarError("Error: el atributo "+attributeName+" no existe en el struct "+currentStructInstance.StructName, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
					return fmt.Sprintf("Error: el atributo %s no existe en el struct %s", attributeName, currentStructInstance.StructName)
				}
			} else {
				l.errores.InsertarError("Error: el valor "+fmt.Sprintf("%v", currentValue)+" no es una instancia de un struct", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return fmt.Sprintf("Error: el valor %v no es una instancia de un struct", currentValue)
			}
		}

		return currentValue
	} else {
		l.errores.InsertarError("Error: la instancia del struct con id "+structID+" no existe", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: la instancia del struct con id %s no existe", structID)
	}
}

/*asignstructstmt
: ID (PUNTO ID)+ IG expr #AsignStruct
;*/

func (l *Visitor) VisitAsignStruct(ctx *parser.AsignStructContext) interface{} {
	structID := ctx.ID(0).GetText()
	attributeNames := ctx.AllID()[1:]
	valor := l.Visit(ctx.Expr())

	// Verificar si el struct existe en el entorno actual
	if structInstance, ok := l.currentEnvironment.Instancias[structID]; ok {
		var currentValue interface{} = structInstance

		for _, attributeNameCtx := range attributeNames {
			attributeName := attributeNameCtx.GetText()

			// Verificar si el valor actual es una instancia de un struct
			if currentStructInstance, ok := currentValue.(StructInstance); ok {
				// Verificar si el struct tiene el atributo
				if _, ok := currentStructInstance.Attributes[attributeName]; ok {
					//actualizar la instanci
					data := infoDataInstance{
						Asignado: true,
						Data:     valor,
					}
					currentStructInstance.Attributes[attributeName] = data
				} else {
					l.errores.InsertarError("Error: el atributo "+attributeName+" no existe en el struct "+currentStructInstance.StructName, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
					return fmt.Sprintf("Error: el atributo %s no existe en el struct %s", attributeName, currentStructInstance.StructName)
				}
			} else {
				l.errores.InsertarError("Error: el valor "+fmt.Sprintf("%v", currentValue)+" no es una instancia de un struct", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				return fmt.Sprintf("Error: el valor %v no es una instancia de un struct", currentValue)
			}
		}
	} else {
		l.errores.InsertarError("Error: la instancia del struct con id "+structID+" no existe", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: la instancia del struct con id %s no existe", structID)
	}
	return nil
}

// vector_struct_stmt
func (l *Visitor) VisitVectorStruct(ctx *parser.VectorStructContext) interface{} {
	idVector := ctx.ID(0).GetText()
	idStruct := ctx.ID(1).GetText()

	//verificar que el struct exista
	if _, ok := l.currentEnvironment.Structs[idStruct]; !ok {
		l.errores.InsertarError("Error: el struct "+idStruct+" no está definido", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: el struct %s no está definido", idStruct)
	}

	//verificar que no exista un vector con el mismo id
	if _, ok := l.currentEnvironment.VectoresStruct[idVector]; ok {
		l.errores.InsertarError("Error: el vector "+idVector+" ya existe", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: el vector %s ya existe", idVector)
	}

	//crear el vector
	vectorStruct := VectorStruct{
		Id:      idVector,
		Tipo:    idStruct,
		Valores: []StructInstance{},
	}

	//agregar el vector al entorno
	l.currentEnvironment.VectoresStruct[idVector] = vectorStruct
	return nil
}

// acceso_vector_struct_stmt
func (l *Visitor) VisitAccesoVectorStruct(ctx *parser.AccesoVectorStructContext) interface{} {
	idvVector := ctx.ID(0).GetText()
	atributoInstancia := ctx.ID(1).GetText()

	//verificar que el vector exista
	if vector, ok := l.currentEnvironment.VectoresStruct[idvVector]; ok {
		//verificar que el indice sea un entero
		indice := l.Visit(ctx.Expr()).(int64)

		if indice >= int64(len(vector.Valores)) {
			l.errores.InsertarError("Error: el indice "+fmt.Sprintf("%d", indice)+" está fuera de los límites del vector "+idvVector, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error: el indice %d está fuera de los límites del vector %s", indice, idvVector)
		}

		//acceder al atributo de la instancia
		if instancia, ok := vector.Valores[indice].Attributes[atributoInstancia]; ok {
			return instancia.Data
		} else {
			l.errores.InsertarError("Error: el atributo "+atributoInstancia+" no existe en el struct "+vector.Tipo, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error: el atributo %s no existe en el struct %s", atributoInstancia, vector.Tipo)
		}
	}
	l.errores.InsertarError("Error: el vector "+idvVector+" no existe", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return fmt.Sprintf("Error: el vector %s no existe", idvVector)
}

func (l *Visitor) VisitAppendVectorStruct(ctx *parser.AppendVectorStructContext) interface{} {
	vectorId := ctx.ID().GetText()
	instancia := l.Visit(ctx.Valor_struct_expr()).(StructInstance)

	//buscar el vector
	if vector, ok := l.currentEnvironment.VectoresStruct[vectorId]; ok {
		//verificar que la instancia sea del mismo tipo que el vector
		if instancia.StructName != vector.Tipo {
			l.errores.InsertarError("Error: la instancia "+instancia.StructName+" no es del mismo tipo que el vector "+vectorId, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			return fmt.Sprintf("Error: la instancia %s no es del mismo tipo que el vector %s", instancia.StructName, vectorId)
		}

		//agregar la instancia al vector
		vector.Valores = append(vector.Valores, instancia)
		l.currentEnvironment.VectoresStruct[vectorId] = vector
	} else {
		l.errores.InsertarError("Error: el vector "+vectorId+" no existe", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return fmt.Sprintf("Error: el vector %s no existe", vectorId)
	}
	return nil
}
