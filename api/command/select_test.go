package command

import (
	"strings"
	"testing"
)

func Test_selectMembers(t *testing.T) {
	type args struct {
		loopCount   int
		parmMembers string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "テスト1",
			args: args{
				loopCount:   2,
				parmMembers: "A,A",
			},
			want: "A,A",
		},
		{
			name: "テスト2",
			args: args{
				loopCount:   3,
				parmMembers: "A,A",
			},
			want: "A,A",
		},
		{
			name: "ランダムの結果が返却されるテスト1",
			args: args{
				loopCount:   3,
				parmMembers: "A,B",
			},
			want: "A,B",
		},
		{
			name: "ランダムの結果が返却されるテスト2",
			args: args{
				loopCount:   5,
				parmMembers: "A,B,C,D,E,F,G",
			},
			want: "A,B,C,D,E",
		},
		{
			name: "テスト3",
			args: args{
				loopCount:   -1,
				parmMembers: "A,B,C,D,E,F,G",
			},
			want: "自然数を指定してください。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectMembers(tt.args.loopCount, tt.args.parmMembers)
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
