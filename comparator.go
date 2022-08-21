package jsonpath

func comparator(leftVal interface{}, rightVal interface{}, op operator) bool {
	switch lVal := leftVal.(type) {
	case int, int64, float32, float64:
		switch rVal := rightVal.(type) {
		case int, int64, float32, float64:
			lValFloat, err := toFloat64(lVal)
			if err != nil {
				return false
			}
			rValFloat, err := toFloat64(rVal)
			if err != nil {
				return false
			}
			return compareFloat64(lValFloat, rValFloat, op)
		default:
			return false
		}
	case string:
		rVal, isString := rightVal.(string)
		if !isString {
			return false
		}
		return compareString(lVal, rVal, op)
	}
	return false
}

func compareString(leftVal, rightVal string, op operator) bool {
	switch op {
	case greater:
		return leftVal > rightVal
	case greaterEq:
		return leftVal >= rightVal
	case less:
		return leftVal < rightVal
	case lessEq:
		return leftVal <= rightVal
	case eq:
		return leftVal == rightVal
	case nEq:
		return leftVal != rightVal
	}
	return false
}

func compareFloat64(leftVal, rightVal float64, op operator) bool {
	switch op {
	case greater:
		return leftVal > rightVal
	case greaterEq:
		return leftVal >= rightVal
	case less:
		return leftVal < rightVal
	case lessEq:
		return leftVal <= rightVal
	case eq:
		return leftVal == rightVal
	case nEq:
		return leftVal != rightVal
	}
	return false
}

func logicalOperation(leftVal, rightVal interface{}, op operator) bool {
	var lBoolVal bool
	var rBoolVal bool
	var isBool bool

	lBoolVal, isBool = leftVal.(bool)
	if !isBool {
		lBoolVal = leftVal != nil
	}
	rBoolVal, isBool = rightVal.(bool)
	if !isBool {
		rBoolVal = rightVal != nil
	}
	switch op {
	case or:
		return lBoolVal || rBoolVal
	case and:
		return lBoolVal && rBoolVal
	}
	return false
}
