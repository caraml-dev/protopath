package protopath

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// FieldAccessOperation is one of jsonpath operation that access field given the object
type FieldAccessOperation struct {
	// fieldName is name of field that will be retrieved from the given object
	fieldName string
	// fieldGetter is way of accessing field, e.g accessing field from protobuf message
	fieldGetter FieldGetter
	// fromRoot indicate that access field from root object not the current object
	fromRoot bool
}

// Lookup evaluate operation and returning field value
func (f *FieldAccessOperation) Lookup(ctx context.Context, obj, rootObj any) (any, error) {
	selectedObj := obj
	if f.fromRoot {
		selectedObj = rootObj
	}
	return f.fieldGetter.GetField(f.fieldName, selectedObj)
}

// FieldGetter get value of field given the field name and object to evaluate
type FieldGetter interface {
	GetField(name string, obj any) (any, error)
}

// ProtoFieldGetter is way to access field for protobuf message object
type ProtoFieldGetter struct{}

// GetField retrieve the field value given the the field name and object
func (p *ProtoFieldGetter) GetField(fieldName string, obj any) (any, error) {
	objReflectVal := reflect.ValueOf(obj)
	switch objReflectVal.Kind() {
	case reflect.Slice:
		res := make([]any, 0, objReflectVal.Len())
		for i := 0; i < objReflectVal.Len(); i++ {
			val, err := p.getFieldFromMessage(fieldName, objReflectVal.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			res = append(res, val)
		}
		return res, nil
	default:
		return p.getFieldFromMessage(fieldName, obj)
	}
}

func (p *ProtoFieldGetter) getFieldFromMessage(fieldName string, obj any) (any, error) {
	protoObj, isProto := obj.(proto.Message)
	if !isProto {
		return nil, fmt.Errorf("object is not in proto.Message type")
	}
	messageReflect := protoObj.ProtoReflect()
	fieldDesciptor := messageReflect.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
	if fieldDesciptor == nil {
		return nil, nil
	}

	fieldValue := messageReflect.Get(fieldDesciptor).Interface()
	switch fieldT := fieldValue.(type) {
	case protoreflect.Message:
		if !fieldT.IsValid() {
			return nil, nil
		}
		return fieldT.Interface(), nil
	case protoreflect.List:
		result := make([]any, 0, fieldT.Len())
		for i := 0; i < fieldT.Len(); i++ {
			val := fieldT.Get(i).Interface()
			switch valT := val.(type) {
			case protoreflect.Message:
				result = append(result, valT.Interface())
			default:
				result = append(result, valT)
			}
		}
		return result, nil
	default:
		return fieldValue, nil
	}
}
