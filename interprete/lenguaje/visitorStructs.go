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
	attri := l.Visit(ctx.Valor_struct_expr()).([]Dupla)

	structInstance := StructInstance{
		StructName: structName,
		Attributes: make(map[string]interface{}),
	}

	for _, attr := range attri {
		structInstance.Attributes[attr.AttributeName] = attr.AttributeValue
	}

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

		return attributes
	} else {
		return fmt.Sprintf("Error: el struct %s no está definido", structName)
	}
}

func (l *Visitor) VisitDuplastruct(ctx *parser.DuplastructContext) interface{} {
	duplas := []Dupla{}

	for i := 0; i < len(ctx.AllID()); i++ {
		attributeName := ctx.ID(i).GetText()
		attributeValue := l.Visit(ctx.Expr(i))
		fmt.Println("attributeValue: ", attributeValue)
		attriType := determineType(attributeValue)
		fmt.Println("attriType: ", attriType)
		
		if attriType == "Error" {
			fmt.Println("entro al error", attributeName)
			//verificar si es un struct
			if structInstance, ok := l.currentEnvironment.Instancias[attributeValue.(string)]; ok {
				attributeValue = structInstance
				fmt.Println("structInstance: ", structInstance)
				attriType = "struct"
			} else {
				return fmt.Sprintf("Error: el atributo %s no tiene un valor valido asignado", attributeName)
			}
		}

		// Agregar la dupla actual al arreglo de duplas
		duplas = append(duplas, Dupla{
			AttributeName:  attributeName,
			AttributeType:  "",
			AttributeValue: attributeValue,
		})
	}
	fmt.Println("DUPLAAAAAAAAAAAAAAAAS: ", duplas)
	return duplas
}

func (l *Visitor) VisitAccesoStruct(ctx *parser.AccesoStructContext) interface{} {
	structName := ctx.ID(0).GetText()
	attributeName := ctx.ID(1).GetText()

	if structInstance, ok := l.currentEnvironment.Instancias[structName]; ok {
		if attributeValue, ok := structInstance.Attributes[attributeName]; ok {
			return attributeValue
		}
		return fmt.Sprintf("Error: el atributo %s no existe en la instancia de %s", attributeName, structName)
	}
	return fmt.Sprintf("Error: la instancia de %s no existe", structName)
}