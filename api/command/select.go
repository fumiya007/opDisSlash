package command

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fumiya007/tpDisSlash/api/command/util"
	"github.com/wafer-bw/disgoslash/discord"
)

// SelectCommand はSelectCommandを実行する
func SelectCommand(request *discord.InteractionRequest) *discord.InteractionResponse {
	loopCount, _ := request.Data.Options[0].IntValue()
	parmMembers, _ := request.Data.Options[1].StringValue()

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("<@%v> /select number:%v members:%v \n%v", request.Member.User.ID, loopCount, parmMembers, selectMembers(loopCount, parmMembers)),
		},
	}
}

func selectMembers(loopCount int, parmMembers string) string {
	if loopCount <= 0 {
		return fmt.Sprintf("自然数を指定してください。")
	}

	members := util.SplitMultiSep(parmMembers, ",", " ", "　")
	loopCount = util.MinInt(loopCount, len(members))
	ret := make([]string, loopCount, loopCount)
	for i := 0; i < loopCount; i++ {
		index := rand.Int() % len(members)
		ret[i] = members[index]
		members = append(members[:index], members[index+1:]...)
	}
	return strings.Join(ret, ",")
}
