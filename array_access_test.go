package jsonpath

import (
	"context"
	"fmt"
	"testing"

	upiV1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/stretchr/testify/assert"
)

func TestArrayAccessOperation_Lookup(t *testing.T) {

	startIdx := 1
	endIdx := 2
	tests := []struct {
		name          string
		arrayAccessOp *ArrayAccessOperation
		obj           interface{}
		want          interface{}
		expectedErr   error
	}{
		{
			name: "using getByIndices method, single index",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByIndices{
					indices: []int{2},
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			want: &upiV1.PredictionRow{
				RowId: "id-3",
				ModelInputs: []*upiV1.NamedValue{
					{
						Name:        "rating",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(4.6),
					},
					{
						Name:        "acceptance_rate",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(3.8),
					},
				},
			},
		},
		{
			name: "using getByIndices method",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByIndices{
					indices: []int{0, 1},
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
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
							DoubleValue: float64(2.8),
						},
					},
				},
			},
		},
		{
			name: "using getByIndices method; out of range is skipped",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByIndices{
					indices: []int{0, 7},
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			expectedErr: fmt.Errorf("index out of bound"),
		},
		{
			name: "using getByRange method without specifying startIdx and endIdx",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
		},
		{
			name: "using getByRange method specifying startIdx and endIdx",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{
					startIdx: &startIdx,
					endIdx:   &endIdx,
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			want: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
			},
		},
		{
			name: "using getByRange method only specifying startIdx ",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{
					startIdx: &startIdx,
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			want: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
		},
		{
			name: "using getByRange method only specifying endIdx ",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{
					endIdx: &endIdx,
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
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
							DoubleValue: float64(2.8),
						},
					},
				},
			},
		},
		{
			name: "using getByRange method only specifying endIdx with negative number ",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{
					endIdx: toPointerInt(-1),
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
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
							DoubleValue: float64(2.8),
						},
					},
				},
			},
		},
		{
			name: "using getByRange method only specifying startIdx with negative number ",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByRange{
					startIdx: toPointerInt(-1),
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			want: []*upiV1.PredictionRow{
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
		},
		{
			name: "using getByBackwardIndex",
			arrayAccessOp: &ArrayAccessOperation{
				method: &getByBackwardIndex{
					index: 1,
				},
			},
			obj: []*upiV1.PredictionRow{
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
							DoubleValue: float64(2.8),
						},
					},
				},
				{
					RowId: "id-3",
					ModelInputs: []*upiV1.NamedValue{
						{
							Name:        "rating",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(4.6),
						},
						{
							Name:        "acceptance_rate",
							Type:        upiV1.NamedValue_TYPE_DOUBLE,
							DoubleValue: float64(3.8),
						},
					},
				},
			},
			want: &upiV1.PredictionRow{
				RowId: "id-3",
				ModelInputs: []*upiV1.NamedValue{
					{
						Name:        "rating",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(4.6),
					},
					{
						Name:        "acceptance_rate",
						Type:        upiV1.NamedValue_TYPE_DOUBLE,
						DoubleValue: float64(3.8),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.arrayAccessOp.Lookup(context.Background(), tt.obj, tt.obj)
			assert.Equal(t, tt.expectedErr, err)

			if tt.expectedErr != nil {
				return
			}
			assertValueEquality(t, tt.want, got)
		})
	}
}

func toPointerInt(val int) *int {
	return &val
}
