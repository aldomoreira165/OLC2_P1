package lenguaje

//para determinar el tipo de una expresion
func determineType(value interface{}) string {
	switch value.(type) {
	case int64:
		return "int"
	case float64:
		return "float"
	case bool:
		return "bool"
	case string:
		if len(value.(string)) == 1 {
			return "character"
		}
		return "String"
	case nil:
		return "nil"
	default:
		return "unknown"
	}
}

// Para determinar si una declaración tiene el mismo tipo de dato que el valor que se le asigna
func validateType(value interface{}, declType string) bool {
	switch declType {
	case "int":
		if value == nil {
			return true
		}
		_, isInt := value.(int64)
		return isInt
	case "float":
		if value == nil {
			return true
		}
		_, isFloat := value.(float64)
		return isFloat
	case "bool":
		if value == nil {
			return true
		}
		_, isBool := value.(bool)
		return isBool
	case "character":
		if value == nil {
			return false
		}
		_, isChar := value.(string)
		return isChar && len(value.(string)) == 1
	case "String":
		if value == nil {
			return false
		}
		_, isString := value.(string)
		return isString && len(value.(string)) > 1
	default:
		return false // Tipo desconocido o no admitido
	}
}

//para convertir una expresion en el tipo que se indica
func convertirExpresion(tipo string ,exp interface{}) interface{} {
	switch tipo {
	case "int":
		result, _ := exp.(int64)
		return result
	case "float":
		result, _ := exp.(float64)
		return result
	case "bool":
		result, _ := exp.(bool)
		return result
	case "character":
		result, _ := exp.(string)
		return result
	case "String":
		result, _ := exp.(string)
		return result
	default:
		return false // Tipo desconocido o no admitido
	}
}


