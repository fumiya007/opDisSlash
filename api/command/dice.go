package command

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"strconv"
	"strings"

	"math/big"
	"math/rand"

	"github.com/wafer-bw/disgoslash/discord"
)

const mesDisError = "1d100のように指定してください。"
const mesManyDice = "20より多くのダイスはふれません。"

// DiceCommand はDiceコマンドを実行する
func DiceCommand(request *discord.InteractionRequest) *discord.InteractionResponse {
	parm, _ := request.Data.Options[0].StringValue()

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("/r parm:%v \n%v", parm, dice(parm)),
		},
	}
}

func dice(parm string) string {

	splitParm := strings.Split(parm, "d")

	if len(splitParm) != 2 {
		return mesDisError
	}

	loopCount, err := strconv.Atoi(splitParm[0])
	if err != nil {
		return mesDisError
	}
	max, err := strconv.Atoi(splitParm[1])
	if err != nil {
		return mesDisError
	}

	if loopCount > 20 {
		return mesManyDice
	}

	diceResult := 0
	midwayFormula := make([]string, loopCount)

	for i := 0; i < loopCount; i++ {
		seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		rand.Seed(seed.Int64())
		n := rand.Intn(max) + 1
		diceResult += n
		midwayFormula[i] = fmt.Sprintf("(%v)", n)
	}

	return fmt.Sprintf("`%v`=%v=%v", parm, strings.Join(midwayFormula, "+"), diceResult)
}
