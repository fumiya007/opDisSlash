package command

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fumiya007/opDisSlash/api/command/util"
	"github.com/wafer-bw/disgoslash/discord"
)

// OrderCommand はOrderCommandを実行する
func OrderCommand(request *discord.InteractionRequest) *discord.InteractionResponse {
	parmMembers, _ := request.Data.Options[0].StringValue()

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("/order members:%v \n%v", parmMembers, order(parmMembers)),
		},
	}
}

func order(parm string) string {
	members := util.SplitMultiSep(parm, ",", " ", "　")
	loopCount := len(members)
	ret := make([]string, loopCount, loopCount)
	for i := 0; i < loopCount; i++ {
		index := rand.Int() % len(members)
		ret[i] = members[index]
		members = append(members[:index], members[index+1:]...)
	}
	return strings.Join(ret, ",")
}
