package command

import (
	"fmt"

	"github.com/wafer-bw/disgoslash/discord"
)

func HelpCommand(request *discord.InteractionRequest) *discord.InteractionResponse {

	helpMessage := `候補から人数分ランダムに選ぶ:/select number:3 members:アーサー ランスロット ガウェイン モードレット
候補をランダムに並べる:/order members:アーサー ランスロット ガウェイン モードレット
ダイスを振る:/r parm:1d100`

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("<@%v>\n%v", request.Member.User.ID, helpMessage),
		},
	}
}
