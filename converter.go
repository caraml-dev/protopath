package jsonpath

import (
	"fmt"
	"strconv"
)

func toFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case *float64:
		return *v, nil
	case float64:
		return v, nil
	case *float32:
		return float64(*v), nil
	case float32:
		return float64(v), nil
	case *int:
		return float64(*v), nil
	case int:
		return float64(v), nil
	case *int8:
		return float64(*v), nil
	case int8:
		return float64(v), nil
	case *int16:
		return float64(*v), nil
	case int16:
		return float64(v), nil
	case *int32:
		return float64(*v), nil
	case int32:
		return float64(v), nil
	case *int64:
		return float64(*v), nil
	case int64:
		return float64(v), nil
	case *string:
		return strconv.ParseFloat(*v, 64)
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("unsupported conversion from %T to float64", v)
	}
}
