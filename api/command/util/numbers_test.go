package util

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	type args struct {
		v1 int
		v2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test MaxInt1",
			args: args{
				v1: 1,
				v2: 2,
			},
			want: 2,
		},
		{
			name: "Test MaxInt2",
			args: args{
				v1: 2,
				v2: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		v1 int
		v2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test MinInt1",
			args: args{
				v1: 1,
				v2: 2,
			},
			want: 1,
		},
		{
			name: "Test MinInt2",
			args: args{
				v1: 2,
				v2: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
