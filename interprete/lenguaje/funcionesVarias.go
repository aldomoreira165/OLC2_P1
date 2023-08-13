package lenguaje

//para determinar el tipo de una expresion
func determineType(value interface{}) string {
	switch value.(type) {
	case int64:
		return "int"
	case float64:
		return "float"
	case bool:
		//ARREGLARRRR
		return "bool"
	case string:
		if len(value.(string)) == 1 {
			return "character"
		}
		return "String"
	default:
		return "unknown"
	}
}

//para determinar si una declaracion tiene el mismo tipo del dato que se le asigna
func validateType(value interface{}, declType string) bool {
	switch declType {
	case "int":
		_, isInt := value.(int64)
		return isInt
	case "float":
		_, isFloat := value.(float64)
		return isFloat
	case "bool":
		//ARREGLARRRRRRRR
		_, isBool := value.(bool)
		return isBool
	case "character":
		_, isChar := value.(string)
		return isChar && len(value.(string)) == 1
	case "String":
		_, isString := value.(string)
		return isString && len(value.(string)) > 1
	default:
		return false // Tipo desconocido o no admitido
	}
}