package util

import (
	"reflect"
	"testing"
)

func TestSplitMultiSep(t *testing.T) {
	type args struct {
		s    string
		seps []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "半角スペースでスプリットできること",
			args: args{
				s:    "A B C",
				seps: []string{" "},
			},
			want: []string{"A", "B", "C"},
		},
		{
			name: "複数のセパレータを指定できること",
			args: args{
				s:    "A B　C,D",
				seps: []string{" ", "　", ","},
			},
			want: []string{"A", "B", "C", "D"},
		},
		{
			name: "セパレータがないときも正常に動作すること",
			args: args{
				s:    "A B C",
				seps: []string{},
			},
			want: []string{"A B C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitMultiSep(tt.args.s, tt.args.seps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitMultiSep() = %v, want %v", got, tt.want)
			}
		})
	}
}
