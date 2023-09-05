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
	structName := ctx.ID(0).GetText()
	structInstance := l.Visit(ctx.Valor_struct_expr()).(StructInstance)
	structInstance.StructName = structName
	l.currentEnvironment.Instancias[structName] = structInstance
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
						AttributeName:  structDef.Attributes[i].Name,
						AttributeType:  structDef.Attributes[i].Type,
						AttributeValue: structDef.Attributes[i].Expression,
					})
				}
			}
		}

		//verificar que todos los struct tengan un valor para cada atributo
		for i := 0; i < len(structDef.Attributes); i++ {
			if attributes[i].AttributeValue == nil {
				return fmt.Sprintf("Error: el atributo %s no tiene un valor asignado", structDef.Attributes[i].Name)
			}
		}

		//retornando data con los atributos
		dataInstance := StructInstance{
			StructName: "",
			Attributes: make(map[string]interface{}),
		}

		for _, attr := range attributes {
			dataInstance.Attributes[attr.AttributeName] = attr.AttributeValue
		}

		return dataInstance
	} else {
		return fmt.Sprintf("Error: el struct %s no está definido", structName)
	}
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
			AttributeName:  attributeName,
			AttributeType:  "",
			AttributeValue: attributeValue,
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
					currentValue = attributeValue
				} else {
					return fmt.Sprintf("Error: el atributo %s no existe en el struct %s", attributeName, currentStructInstance.StructName)
				}
			} else {
				return fmt.Sprintf("Error: el valor %v no es una instancia de un struct", currentValue)
			}
		}

		return currentValue
	} else {
		return fmt.Sprintf("Error: la instancia del struct con id %s no existe", structID)
	}
}
