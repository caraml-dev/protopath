package jsonpath

import (
	"context"
	"reflect"
	"testing"

	upiV1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestFieldAccessOperation_Lookup(t *testing.T) {

	obj := &upiV1.PredictValuesRequest{
		TargetName: "target_name",
		Metadata: &upiV1.RequestMetadata{
			PredictionId: "predictionID",
		},
		PredictionRows: []*upiV1.PredictionRow{
			{
				RowId: "id-1",
				ModelInputs: []*upiV1.NamedValue{
					{
						Name:        "rating",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(4.2),
					},
					{
						Name:        "acceptance_rate",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(0.8),
					},
				},
			},
		},
	}

	tests := []struct {
		name          string
		obj           interface{}
		rootObj       interface{}
		fieldAccessOp *FieldAccessOperation
		want          interface{}
		expectedErr   error
	}{
		{
			name:    "simple field access",
			obj:     obj,
			rootObj: obj,
			fieldAccessOp: &FieldAccessOperation{
				fieldName:   "target_name",
				fieldGetter: &ProtoFieldGetter{},
				fromRoot:    true,
			},
			want: "target_name",
		},
		{
			name:    "simple field access returning object",
			obj:     obj,
			rootObj: obj,
			fieldAccessOp: &FieldAccessOperation{
				fieldName:   "metadata",
				fieldGetter: &ProtoFieldGetter{},
				fromRoot:    true,
			},
			want: &upiV1.RequestMetadata{
				PredictionId: "predictionID",
			},
		},
		{
			name:    "simple field access returning array of object",
			obj:     obj,
			rootObj: obj,
			fieldAccessOp: &FieldAccessOperation{
				fieldName:   "prediction_rows",
				fieldGetter: &ProtoFieldGetter{},
				fromRoot:    true,
			},
			want: []*upiV1.PredictionRow{
				{
					RowId: "id-1",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.2),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(0.8),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fieldAccessOp.Lookup(context.Background(), tt.obj, tt.rootObj)
			assert.Equal(t, tt.expectedErr, err)
			if tt.expectedErr != nil {
				return
			}

			assertValueEquality(t, tt.want, got)
		})
	}
}

func assertValueEquality(t *testing.T, expected, actual interface{}) {
	wantReflectVal := reflect.ValueOf(expected)

	switch wantReflectVal.Kind() {
	case reflect.Slice:
		gotReflectVal := reflect.ValueOf(actual)
		if gotReflectVal.Len() != wantReflectVal.Len() {
			t.Errorf("FieldAccessOperation.Lookup() having different length between actual: %d and expected: %d", gotReflectVal.Len(), wantReflectVal.Len())
			return
		}

		for i := 0; i < wantReflectVal.Len(); i++ {
			wVal := wantReflectVal.Index(i).Interface()
			assertValueEquality(t, wVal, gotReflectVal.Index(i).Interface())
		}
	default:
		if expectedVal, isProto := expected.(proto.Message); isProto {
			if !proto.Equal(expectedVal, actual.(proto.Message)) {
				t.Errorf("FieldAccessOperation.Lookup() = %v, want %v", actual, expected)
			}
			return
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("FieldAccessOperation.Lookup() = %v, want %v", actual, expected)
		}
	}
}
