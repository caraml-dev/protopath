package jsonpath

import (
	"context"
	"testing"

	upiV1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/stretchr/testify/assert"
)

func TestQueryFilterOp_Lookup(t *testing.T) {
	tests := []struct {
		name        string
		queryOp     *QueryFilterOp
		obj         interface{}
		rootObj     interface{}
		want        interface{}
		expectedErr error
	}{
		{
			name: "filter using greater comparator for double",
			queryOp: &QueryFilterOp{
				LeftOp: []Operation{
					&FieldAccessOperation{
						fieldName:   "double_value",
						fieldGetter: &ProtoFieldGetter{},
						fromRoot:    false,
					},
				},
				Operator: greater,
				RightOp:  int(2),
			},
			obj: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_60d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
			want: []*upiV1.NamedValue{
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
		},
		{
			name: "filter using equality comparator",
			queryOp: &QueryFilterOp{
				LeftOp: []Operation{
					&FieldAccessOperation{
						fieldName:   "name",
						fieldGetter: &ProtoFieldGetter{},
						fromRoot:    false,
					},
				},
				Operator: eq,
				RightOp:  "avg_30d_total_orders",
			},
			obj: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
			want: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
			},
		},
		{
			name: "filter using equality comparator",
			queryOp: &QueryFilterOp{
				LeftOp: []Operation{
					&FieldAccessOperation{
						fieldName:   "name",
						fieldGetter: &ProtoFieldGetter{},
						fromRoot:    false,
					},
				},
				Operator: eq,
				RightOp:  "avg_30d_total_orders",
			},
			obj: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
			want: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
			},
		},
		{
			name: "filter using logical operator; and operator",
			queryOp: &QueryFilterOp{
				LeftOp: &QueryFilterOp{
					LeftOp: []Operation{
						&FieldAccessOperation{
							fieldName:   "name",
							fieldGetter: &ProtoFieldGetter{},
							fromRoot:    false,
						},
					},
					Operator: eq,
					RightOp:  "avg_30d_total_orders",
				},
				Operator: and,
				RightOp: &QueryFilterOp{
					LeftOp: []Operation{
						&FieldAccessOperation{
							fieldName:   "double_value",
							fieldGetter: &ProtoFieldGetter{},
							fromRoot:    false,
						},
					},
					Operator: greater,
					RightOp:  int(2),
				},
			},
			obj: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
			want: []*upiV1.NamedValue{},
		},
		{
			name: "filter using logical operator; or operator",
			queryOp: &QueryFilterOp{
				LeftOp: &QueryFilterOp{
					LeftOp: []Operation{
						&FieldAccessOperation{
							fieldName:   "name",
							fieldGetter: &ProtoFieldGetter{},
							fromRoot:    false,
						},
					},
					Operator: eq,
					RightOp:  "avg_30d_total_orders",
				},
				Operator: or,
				RightOp: &QueryFilterOp{
					LeftOp: []Operation{
						&FieldAccessOperation{
							fieldName:   "double_value",
							fieldGetter: &ProtoFieldGetter{},
							fromRoot:    false,
						},
					},
					Operator: greater,
					RightOp:  int(2),
				},
			},
			obj: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
			want: []*upiV1.NamedValue{
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 1.9,
				},
				{
					Name:        "avg_30d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 2.0,
				},
				{
					Name:        "avg_90d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 3.2,
				},
				{
					Name:        "avg_120d_total_orders",
					Type:        upiV1.NamedValue_TYPE_DOUBLE,
					DoubleValue: 4.0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.queryOp.Lookup(context.Background(), tt.obj, tt.rootObj)
			assert.Equal(t, tt.expectedErr, err)
			if tt.expectedErr != nil {
				return
			}
			assertValueEquality(t, tt.want, got)
		})
	}
}
