package jsonpath

import (
	"context"
	"reflect"
)

type operator uint

const (
	greater operator = iota
	greaterEq
	less
	lessEq
	eq
	nEq
	or
	and
)

// QueryFilterOp is operation to filter given array of object
// LeftOp must be present
// RightOp and Operator is optional
// LeftOp and RightOp possible values:
// 1. List of operation. Needs to evaluate all the operations to get the value
// 2. QueryFilterOp. Operation contains comparation or logical operation
// 3. Exact value, e.g string or float64
type QueryFilterOp struct {
	LeftOp   interface{}
	RightOp  interface{}
	Operator operator
}

var _ Operation = (*QueryFilterOp)(nil)

// Lookup find the filtered value given object
func (q *QueryFilterOp) Lookup(ctx context.Context, obj, rootObj interface{}) (interface{}, error) {
	reflectVal := reflect.ValueOf(obj)
	switch reflectVal.Kind() {
	case reflect.Slice:
		res := make([]interface{}, 0)
		for i := 0; i < reflectVal.Len(); i++ {
			currVal := reflectVal.Index(i).Interface()
			currReflectVal := reflect.ValueOf(currVal)
			// if each of entry is also a slice, we need to flatten it
			if currReflectVal.Kind() == reflect.Slice {
				for j := 0; j < currReflectVal.Len(); j++ {
					innerVal := currReflectVal.Index(j).Interface()
					meetCriteria, err := q.eval(ctx, innerVal, rootObj)
					if err != nil {
						return nil, err
					}
					if meetCriteria {
						res = append(res, innerVal)
					}
				}
				continue
			}
			meetCriteria, err := q.eval(ctx, currVal, rootObj)
			if err != nil {
				return nil, err
			}
			if meetCriteria {
				res = append(res, currVal)
			}
		}
		return res, nil
	default:
		meetCriteria, err := q.eval(ctx, obj, rootObj)
		if err != nil {
			return nil, err
		}
		if !meetCriteria {
			return nil, nil
		}
		return obj, nil
	}
}

func (q *QueryFilterOp) eval(ctx context.Context, obj, rootObj interface{}) (bool, error) {
	leftVal, err := q.getValFromAnyOperations(ctx, q.LeftOp, obj, rootObj)
	if err != nil {
		return false, err
	}

	// return early when left value is boolean with conditions
	// 1. operator is 'or'
	// 2. there is no right operation
	if boolVal, isBoolVal := leftVal.(bool); isBoolVal {
		if q.Operator == or && boolVal == true {
			return true, nil
		}
		if q.RightOp == nil {
			return boolVal, nil
		}
	}

	// if there is no right operation but the left value is not in bool type
	// it check whether the value is nil
	if q.RightOp == nil {
		return leftVal != nil, nil
	}

	rightVal, err := q.getValFromAnyOperations(ctx, q.RightOp, obj, rootObj)
	if err != nil {
		return false, err
	}

	switch q.Operator {
	case greater, greaterEq, less, lessEq, eq, nEq:
		return comparator(leftVal, rightVal, q.Operator), nil
	case or, and:
		return logicalOperation(leftVal, rightVal, q.Operator), nil
	}
	return false, nil
}

func (q *QueryFilterOp) getValFromAnyOperations(ctx context.Context, ops, obj, rootObj interface{}) (interface{}, error) {
	switch opsT := ops.(type) {
	case []Operation:
		val := obj
		var err error
		for _, op := range opsT {
			val, err = op.Lookup(ctx, val, rootObj)
			if err != nil {
				return nil, err
			}
		}
		return val, nil
	case *QueryFilterOp:
		return opsT.eval(ctx, obj, rootObj)
	default:
		return opsT, nil
	}
}
