package command

import (
	"strings"
	"testing"
)

func Test_orderMembers(t *testing.T) {
	type args struct {
		parm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "テスト1",
			args: args{parm: "b"},
			want: "b",
		},
		{
			name: "ランダムの結果が返却されるテスト1",
			args: args{parm: "5 A,B,C,D,E,F,G"},
			want: "5 A,B,C,D,E,F,G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := order(tt.args.parm)
			if !strings.HasPrefix(tt.name, "ランダム") {
				if got != tt.want {
					t.Errorf("OrderCommand() = %v, want %v", got, tt.want)
				}
			} else {
				if len(got) != len(tt.want) {
					t.Errorf("len(OrderCommand()) = %v, want %v", len(got), len(tt.want))
				}
			}
		})
	}
}
