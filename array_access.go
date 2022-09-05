package protopath

import (
	"context"
	"fmt"
	"reflect"
)

// ArrayAccessMethod responsible to execute any method to access array field
type ArrayAccessMethod interface {
	Get(ctx context.Context, obj any) (any, error)
}

// ArrayAccessOperation is interface to access array fields
type ArrayAccessOperation struct {
	method ArrayAccessMethod
}

// Lookup find the desired value based on defined operation
func (a *ArrayAccessOperation) Lookup(ctx context.Context, obj, _ any) (any, error) {
	return a.method.Get(ctx, obj)
}

// getByIndices, array access method by specified indices
// e.g $.arr[1] or $.arr[1,2,5]
type getByIndices struct {
	indices []int
}

// Get find the desired value based on the defined method
func (gbi *getByIndices) Get(ctx context.Context, val any) (any, error) {
	reflectVal := reflect.ValueOf(val)
	if reflectVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("only work for slice")
	}
	result := make([]any, 0)
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

// getByRange, array access method by specify the range
// e.g
// 1. $.arr[1:5]
// 2. $.arr[1:]
// 3. $.arr[:3]
// 4. $.arr[:]
type getByRange struct {
	startIdx *int
	endIdx   *int
}

// Get find the desired value based on the defined method
func (i *getByRange) Get(ctx context.Context, val any) (any, error) {
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

	result := make([]any, 0)
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

// getByBackwardIndex, array access method to find the object in the specific index counted from the end of index
// e.g $.arr[@.length-4]
type getByBackwardIndex struct {
	index int
}

// Get find the desired value based on the defined method
func (i *getByBackwardIndex) Get(ctx context.Context, val any) (any, error) {
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

	result := make([]any, 0)
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
