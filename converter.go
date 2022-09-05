package protopath

import (
	"fmt"
	"reflect"
	"strconv"
)

func toFloat64(v any) (float64, error) {
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

func toBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case *bool:
		return *v, nil
	case bool:
		return v, nil
	case int, int8, int16, int32, int64:
		i := reflect.ValueOf(v).Int()
		if i == 1 {
			return true, nil
		} else if i == 0 {
			return false, nil
		}
		return false, fmt.Errorf("error parsing %v to bool", v)
	case float32, float64:
		i := reflect.ValueOf(v).Float()
		if i == float64(1) {
			return true, nil
		} else if i == float64(0) {
			return false, nil
		}
		return false, fmt.Errorf("error parsing %v to bool", v)
	case *string:
		return strconv.ParseBool(*v)
	case string:
		return strconv.ParseBool(v)
	default:
		return false, fmt.Errorf("unsupported conversion from %T to bool", v)
	}
}
