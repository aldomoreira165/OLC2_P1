package lenguaje

type Struct struct {
    Name       string
    Attributes []Attribute
}

type Attribute struct {
    Name       string
    IsMutable  bool
    Type       string
    Expression interface{}
}

type Dupla struct {
    AttributeName string
    AttributeType string
    AttributeValue interface{}
}

type StructInstance struct {
    StructName string
    Attributes map[string]interface{}
}