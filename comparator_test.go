package protopath

import "testing"

func Test_comparator(t *testing.T) {
	type args struct {
		leftVal  interface{}
		rightVal interface{}
		op       operator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "compare float64 greater",
			args: args{
				leftVal:  2.0,
				rightVal: 3,
				op:       greater,
			},
			want: false,
		},
		{
			name: "compare float64 greaterEq",
			args: args{
				leftVal:  2.0,
				rightVal: 2.0,
				op:       greaterEq,
			},
			want: true,
		},
		{
			name: "compare float64 less",
			args: args{
				leftVal:  2.0,
				rightVal: 3,
				op:       less,
			},
			want: true,
		},
		{
			name: "compare float64 with bool",
			args: args{
				leftVal:  2.0,
				rightVal: true,
				op:       less,
			},
			want: false,
		},
		{
			name: "compare bool with float64",
			args: args{
				leftVal:  true,
				rightVal: 2.0,
				op:       less,
			},
			want: false,
		},
		{
			name: "compare float64 eq",
			args: args{
				leftVal:  2.0,
				rightVal: 2.0,
				op:       eq,
			},
			want: true,
		},
		{
			name: "compare float64 nEq",
			args: args{
				leftVal:  2.0,
				rightVal: 2.0,
				op:       nEq,
			},
			want: false,
		},
		{
			name: "compare string greater",
			args: args{
				leftVal:  "ab",
				rightVal: "aa",
				op:       greater,
			},
			want: true,
		},
		{
			name: "compare string greaterEq",
			args: args{
				leftVal:  "ab",
				rightVal: "ab",
				op:       greaterEq,
			},
			want: true,
		},
		{
			name: "compare string less",
			args: args{
				leftVal:  "ab",
				rightVal: "aa",
				op:       less,
			},
			want: false,
		},
		{
			name: "compare string eq",
			args: args{
				leftVal:  "aa",
				rightVal: "cc",
				op:       eq,
			},
			want: false,
		},
		{
			name: "compare string nEq",
			args: args{
				leftVal:  "aa",
				rightVal: "bb",
				op:       nEq,
			},
			want: true,
		},
		{
			name: "compare string with bool",
			args: args{
				leftVal:  "aa",
				rightVal: false,
				op:       nEq,
			},
			want: false,
		},
		{
			name: "compare bool with string",
			args: args{
				leftVal:  false,
				rightVal: "aa",
				op:       nEq,
			},
			want: false,
		},
		{
			name: "compare bool with bool; nEq",
			args: args{
				leftVal:  false,
				rightVal: true,
				op:       nEq,
			},
			want: true,
		},
		{
			name: "compare bool with bool; eq",
			args: args{
				leftVal:  false,
				rightVal: true,
				op:       eq,
			},
			want: false,
		},
		{
			name: "compare bool with bool; greater",
			args: args{
				leftVal:  true,
				rightVal: true,
				op:       greater,
			},
			want: false,
		},
		{
			name: "compare bool with bool; greaterEq",
			args: args{
				leftVal:  true,
				rightVal: true,
				op:       greaterEq,
			},
			want: false,
		},
		{
			name: "compare bool with bool; less",
			args: args{
				leftVal:  true,
				rightVal: true,
				op:       less,
			},
			want: false,
		},
		{
			name: "compare bool with bool; lessEq",
			args: args{
				leftVal:  true,
				rightVal: true,
				op:       lessEq,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comparator(tt.args.leftVal, tt.args.rightVal, tt.args.op); got != tt.want {
				t.Errorf("comparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
