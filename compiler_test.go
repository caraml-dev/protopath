package jsonpath

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	upiV1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompiler_NewJsonPathCompiler(t *testing.T) {
	type args struct {
		jsonpath string
	}

	intOne := 1
	intThree := 3
	fieldGetter := &ProtoFieldGetter{}
	tests := []struct {
		name        string
		args        args
		want        *Compiler
		expectedErr error
	}{
		{
			name: "invalid field access",
			args: args{
				jsonpath: "$.12",
			},
			expectedErr: fmt.Errorf("got syntax error"),
		},
		{
			name: "field access",
			args: args{
				jsonpath: "$.field1",
			},
			want: &Compiler{
				Path: "$.field1",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "nested field access",
			args: args{
				jsonpath: "$.field1.field2",
			},
			want: &Compiler{
				Path: "$.field1.field2",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array all objects access",
			args: args{
				jsonpath: "$.field1[*]",
			},
			want: &Compiler{
				Path: "$.field1[*]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&ArrayAccessOperation{
						method: &getByRange{},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array all objects access with invalid symbol",
			args: args{
				jsonpath: "$.field1[**]",
			},
			expectedErr: fmt.Errorf("got syntax error"),
		},
		{
			name: "array access with indices",
			args: args{
				jsonpath: "$.field1.[1,2,3]",
			},
			want: &Compiler{
				Path: "$.field1.[1,2,3]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&ArrayAccessOperation{
						method: &getByIndices{
							indices: []int{1, 2, 3},
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access with indices but not int",
			args: args{
				jsonpath: "$.field1.[one, two, three]",
			},
			expectedErr: fmt.Errorf("got syntax error"),
		},
		{
			name: "array access with range",
			args: args{
				jsonpath: "$.field1.field2[1:3]",
			},
			want: &Compiler{
				Path: "$.field1.field2[1:3]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByRange{
							startIdx: &intOne,
							endIdx:   &intThree,
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access with invalid range",
			args: args{
				jsonpath: "$.field1.field2[startIdx:endIdx]",
			},
			expectedErr: fmt.Errorf("got syntax error"),
		},
		{
			name: "array access with range",
			args: args{
				jsonpath: "$.field1.field2[1:3]",
			},
			want: &Compiler{
				Path: "$.field1.field2[1:3]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByRange{
							startIdx: &intOne,
							endIdx:   &intThree,
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access with range; only set startIdx",
			args: args{
				jsonpath: "$.field1.field2[1:]",
			},
			want: &Compiler{
				Path: "$.field1.field2[1:]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByRange{
							startIdx: &intOne,
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access with range; only set endIdx",
			args: args{
				jsonpath: "$.field1.field2[:3]",
			},
			want: &Compiler{
				Path: "$.field1.field2[:3]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByRange{
							endIdx: &intThree,
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access with range; not setting startIdx and endIdx",
			args: args{
				jsonpath: "$.field1.field2[:]",
			},
			want: &Compiler{
				Path: "$.field1.field2[:]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByRange{},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "array access by index backward",
			args: args{
				jsonpath: "$.field1.field2[@.length-1]",
			},
			want: &Compiler{
				Path: "$.field1.field2[@.length-1]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&ArrayAccessOperation{
						method: &getByBackwardIndex{
							index: 1,
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "filter check existence",
			args: args{
				jsonpath: "$.field1.field2[?(@.val)]",
			},
			want: &Compiler{
				Path: "$.field1.field2[?(@.val)]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&QueryFilterOp{
						LeftOp: []Operation{
							&FieldAccessOperation{
								fieldName:   "val",
								fieldGetter: fieldGetter,
								fromRoot:    false,
							},
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "filter with comparator",
			args: args{
				jsonpath: "$.field1.field2[?(@.val > 2.2)]",
			},
			want: &Compiler{
				Path: "$.field1.field2[?(@.val > 2.2)]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&QueryFilterOp{
						LeftOp: []Operation{
							&FieldAccessOperation{
								fieldName:   "val",
								fieldGetter: fieldGetter,
								fromRoot:    false,
							},
						},
						Operator: greater,
						RightOp:  float64(2.2),
					},
				},
				fieldGetter: fieldGetter,
			},
		},
		{
			name: "filter with comparator incorrect right val",
			args: args{
				jsonpath: "$.field1.field2[?(@.val > randomVal)]",
			},
			expectedErr: fmt.Errorf("got syntax error"),
		},
		{
			name: "filter with logical operation",
			args: args{
				jsonpath: "$.field1.field2[?(@.val > 2.2 || $.field3 < $.field4)]",
			},
			want: &Compiler{
				Path: "$.field1.field2[?(@.val > 2.2 || $.field3 < $.field4)]",
				Operations: []Operation{
					&FieldAccessOperation{
						fieldName:   "field1",
						fieldGetter: fieldGetter,
						fromRoot:    true,
					},
					&FieldAccessOperation{
						fieldName:   "field2",
						fieldGetter: fieldGetter,
						fromRoot:    false,
					},
					&QueryFilterOp{
						LeftOp: &QueryFilterOp{
							LeftOp: []Operation{
								&FieldAccessOperation{
									fieldName:   "val",
									fieldGetter: fieldGetter,
									fromRoot:    false,
								},
							},
							Operator: greater,
							RightOp:  float64(2.2),
						},
						Operator: or,
						RightOp: &QueryFilterOp{
							LeftOp: []Operation{
								&FieldAccessOperation{
									fieldName:   "field3",
									fieldGetter: fieldGetter,
									fromRoot:    true,
								},
							},
							Operator: less,
							RightOp: []Operation{
								&FieldAccessOperation{
									fieldName:   "field4",
									fieldGetter: fieldGetter,
									fromRoot:    true,
								},
							},
						},
					},
				},
				fieldGetter: fieldGetter,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJsonPathCompiler(tt.args.jsonpath)
			require.Equal(t, tt.expectedErr, err)
			if tt.expectedErr != nil {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJsonPathCompiler() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestCompiler_Lookup(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		obj         interface{}
		want        interface{}
		expectedErr error
	}{
		{
			name: "simple field access",
			path: "$.metadata.prediction_id",
			obj: &upiV1.PredictValuesRequest{
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
			},
			want: "predictionID",
		},
		{
			name: "array access",
			path: "$.prediction_rows[*].model_inputs",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				[]*upiV1.NamedValue{
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

				[]*upiV1.NamedValue{
					{
						Name:        "rating",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(4.4),
					},
					{
						Name:        "acceptance_rate",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(0.9),
					},
				},
			},
		},
		{
			name: "array access; get all",
			path: "$.prediction_rows[*].model_inputs[*]",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.2),
				},
				&upiV1.NamedValue{
					Name:        "acceptance_rate",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(0.8),
				},

				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.4),
				},
				&upiV1.NamedValue{
					Name:        "acceptance_rate",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(0.9),
				},
			},
		},
		{
			name: "array access; get index",
			path: "$.prediction_rows[*].model_inputs[0]",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.2),
				},
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.4),
				},
			},
		},
		{
			name: "array access; get by range",
			path: "$.prediction_rows[*].model_inputs[1:2]",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				&upiV1.NamedValue{
					Name:        "acceptance_rate",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(0.8),
				},
				&upiV1.NamedValue{
					Name:        "acceptance_rate",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(0.9),
				},
			},
		},
		{
			name: "filter",
			path: "$.prediction_rows[*].model_inputs[?(@.name == 'rating')].double_value",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				float64(4.2),
				float64(4.4),
			},
		},
		{
			name: "complex filter; with and condition",
			path: "$.prediction_rows[*].model_inputs[?(@.name == 'rating' && $.target_name == 'target_name')]",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.2),
				},
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.4),
				},
			},
		},
		{
			name: "complex filter; with and condition",
			path: "$.prediction_rows[*].model_inputs[?(@.name == 'rating' || $.target_name == 'wrong')]",
			obj: &upiV1.PredictValuesRequest{
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
					{
						RowId: "id-2",
						ModelInputs: []*upiV1.NamedValue{
							{
								Name:        "rating",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(4.4),
							},
							{
								Name:        "acceptance_rate",
								Type:        upiV1.NamedValue_TYPE_DOUBLE,
								DoubleValue: float64(0.9),
							},
						},
					},
				},
			},
			want: []interface{}{
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.2),
				},
				&upiV1.NamedValue{
					Name:        "rating",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: float64(4.4),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compiler, err := NewJsonPathCompiler(tt.path)
			require.NoError(t, err)
			got, err := compiler.Lookup(context.Background(), tt.obj)
			assert.Equal(t, tt.expectedErr, err)
			if tt.expectedErr != nil {
				return
			}
			assertValueEquality(t, tt.want, got)
		})
	}
}
