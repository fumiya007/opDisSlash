package command

import (
	"regexp"
	"strings"
	"testing"
)

func Test_dice(t *testing.T) {
	type args struct {
		parm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ランダムの結果が返却されるテスト1",
			args: args{
				parm: "1d100",
			},
			want: "`1d100`=(.*)=[0-9*]",
		},
		{
			name: "ランダムの結果が返却されるテスト2",
			args: args{
				parm: "20d100",
			},
			want: "`20d100`=(.*)=[0-9*]",
		},
		{
			name: "フォーマットエラー1",
			args: args{
				parm: "100",
			},
			want: "1d100のように指定してください。",
		},
		{
			name: "フォーマットエラー2",
			args: args{
				parm: "1d1d1",
			},
			want: "1d100のように指定してください。",
		},
		{
			name: "ダイスの数が多すぎるエラー",
			args: args{
				parm: "21d100",
			},
			want: "20より多くのダイスはふれません。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dice(tt.args.parm)
			if !strings.HasPrefix(tt.name, "ランダム") {
				if got != tt.want {
					t.Errorf("diceCommand() = %v, want %v", got, tt.want)
				}
			} else {
				r := regexp.MustCompile(tt.want)
				if !r.MatchString(got) {
					t.Errorf("diceCommand() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
