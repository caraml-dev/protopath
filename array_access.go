package jsonpath

import (
	"context"
	"fmt"
	"reflect"
)

type ArrayAccessMethod interface {
	Get(ctx context.Context, obj interface{}) (interface{}, error)
}

type ArrayAccessOperation struct {
	method ArrayAccessMethod
}

func (a *ArrayAccessOperation) Lookup(ctx context.Context, obj, _ interface{}) (interface{}, error) {
	return a.method.Get(ctx, obj)
}

type getByIndices struct {
	indices []int
}

func (gbi *getByIndices) Get(ctx context.Context, val interface{}) (interface{}, error) {
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("only work for slice")
	}
	result := make([]interface{}, 0)
	lenOfArray := reflectVal.Len()

	isNestedArray := reflect.TypeOf(reflectVal.Index(0).Interface()).Kind() == reflect.Slice

	if len(gbi.indices) == 1 && !isNestedArray {
		idx := gbi.indices[0]
		if idx >= lenOfArray {
			return nil, fmt.Errorf("index out of bound")
		}
		return reflectVal.Index(idx).Interface(), nil
	}

	if isNestedArray {
		for i := 0; i < lenOfArray; i++ {
			currVal := reflectVal.Index(i).Interface()
			currValReflection := reflect.ValueOf(currVal)
			for _, idxVal := range gbi.indices {
				if idxVal >= currValReflection.Len() || idxVal < 0 {
					return nil, fmt.Errorf("index out of bound")
				}
				result = append(result, currValReflection.Index(idxVal).Interface())
			}
		}
		return result, nil
	}
	for _, idxVal := range gbi.indices {
		if idxVal >= lenOfArray || idxVal < 0 {
			return nil, fmt.Errorf("index out of bound")
		}
		result = append(result, reflectVal.Index(idxVal).Interface())
	}
	return result, nil
}

type getByRange struct {
	startIdx *int
	endIdx   *int
}

func (i *getByRange) Get(ctx context.Context, val interface{}) (interface{}, error) {
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("only work for slice")
	}
	lenOfArray := reflectVal.Len()
	startIdx := 0
	if i.startIdx != nil {
		startIdx = *i.startIdx
	}
	if startIdx < 0 {
		startIdx = startIdx + lenOfArray
	}
	endIdx := lenOfArray
	if i.endIdx != nil {
		endIdx = *i.endIdx
	}
	if endIdx < 0 {
		endIdx = endIdx + lenOfArray
	}

	if startIdx > endIdx {
		return nil, fmt.Errorf("end index should be greater than start index")
	}
	if endIdx > lenOfArray {
		return nil, fmt.Errorf("end index should be less or equal length of rows")
	}

	result := make([]interface{}, 0)
	if lenOfArray == 0 {
		return result, nil
	}

	isNestedArray := reflect.TypeOf(reflectVal.Index(0).Interface()).Kind() == reflect.Slice
	for i := 0; i < lenOfArray; i++ {
		currVal := reflectVal.Index(i).Interface()
		if isNestedArray {
			currValReflection := reflect.ValueOf(currVal)
			for j := 0; j < currValReflection.Len(); j++ {
				if j < startIdx || j >= endIdx {
					continue
				}
				result = append(result, currValReflection.Index(j).Interface())
			}
			continue
		}

		if i < startIdx || i >= endIdx {
			continue
		}
		result = append(result, currVal)
	}
	return result, nil
}

type getByBackwardIndex struct {
	index int
}

func (i *getByBackwardIndex) Get(ctx context.Context, val interface{}) (interface{}, error) {
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("only work for slice")
	}
	lenOfArray := reflectVal.Len()
	idx := lenOfArray - i.index
	if idx < 0 {
		return nil, fmt.Errorf("index should be greater equal 0")
	}

	if lenOfArray == 0 {
		return nil, nil
	}

	isNestedArray := reflect.TypeOf(reflectVal.Index(0).Interface()).Kind() == reflect.Slice
	if !isNestedArray {
		return reflectVal.Index(idx).Interface(), nil
	}

	result := make([]interface{}, 0)
	for i := 0; i < lenOfArray; i++ {
		currVal := reflectVal.Index(i).Interface()
		currValReflection := reflect.ValueOf(currVal)
		if idx < currValReflection.Len() {
			return nil, fmt.Errorf("index out of bound")
		}
		result = append(result, currValReflection.Index(idx).Interface())
	}
	return result, nil
}
