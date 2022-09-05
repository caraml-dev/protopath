package protopath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "from float64",
			args: args{
				v: float64(11.111),
			},
			want:    float64(11.111),
			wantErr: false,
		},
		{
			name: "from float32",
			args: args{
				v: float32(11.111),
			},
			want:    float64(11.111),
			wantErr: false,
		},
		{
			name: "from int",
			args: args{
				v: int(1111),
			},
			want:    float64(1111),
			wantErr: false,
		},
		{
			name: "from int8",
			args: args{
				v: int8(111),
			},
			want:    float64(111),
			wantErr: false,
		},
		{
			name: "from int16",
			args: args{
				v: int16(11111),
			},
			want:    float64(11111),
			wantErr: false,
		},
		{
			name: "from int32",
			args: args{
				v: int32(11111),
			},
			want:    float64(11111),
			wantErr: false,
		},
		{
			name: "from int64",
			args: args{
				v: int64(11111),
			},
			want:    float64(11111),
			wantErr: false,
		},
		{
			name: "from string",
			args: args{
				v: "1111.1111",
			},
			want:    float64(1111.1111),
			wantErr: false,
		},
		{
			name: "from bool",
			args: args{
				v: false,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toFloat64(tt.args.v)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.InEpsilon(t, tt.want, got, 0.0001)
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "from bool",
			args: args{
				v: true,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "from string",
			args: args{
				v: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "from int",
			args: args{
				v: 0,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "from int",
			args: args{
				v: 1,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "from int",
			args: args{
				v: 3,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "from float",
			args: args{
				v: 1.1,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "from float",
			args: args{
				v: float64(1.0),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "from float",
			args: args{
				v: float64(0.0),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBool(tt.args.v)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
