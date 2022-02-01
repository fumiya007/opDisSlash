// Package api provides a Vercel Serverless Function that uses disgoslash
// to serve Discord Slash Command Interactions.
//
// https://vercel.com/docs/serverless-functions/supported-languages#go
//
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
package api

import (
	"net/http"
	"os"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"

	"github.com/fumiya007/tpDisSlash/api/command"
)

// GuildIDs holds the list of Guild (server) IDs you would like to register
// a slash command to.
var GuildIDs = []string{}

// Global indicates whether or not a slash command should be registered globally
// across all Guilds the bot has access to.
var Global = true

var credentials *discord.Credentials

func GetCredentials() *discord.Credentials {
	if credentials == nil {
		credentials = &discord.Credentials{
			PublicKey: os.Getenv("PUBLIC_KEY"), // Your Discord Application's Public Key
			ClientID:  os.Getenv("CLIENT_ID"),  // Your Discord Application's Client ID
			Token:     os.Getenv("TOKEN"),      // Your Discord Application's Bot's Token
		}
	}

	return credentials
}

var helpAC = &discord.ApplicationCommand{
	Name:              "help-tp",
	Description:       "opボットのヘルプ",
	Options:           []*discord.ApplicationCommandOption{},
	DefaultPermission: true,
}

var selectAC = &discord.ApplicationCommand{
	Name:        "select",
	Description: "候補から人数分ランダムに選ぶ",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeInteger,
			Name:        "number",
			Description: "選抜数を入力してください",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "members",
			Description: "候補一覧を「,」、「 」または「　」区切りで入力してください。例:アーサー,ランスロット,ガウェイン,モードレット",
			Required:    true,
		},
	},
	DefaultPermission: true,
}

var orderAC = &discord.ApplicationCommand{
	Name:        "order",
	Description: "ランダムに選びかえる",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "members",
			Description: "並び変えたいものを「,」、「 」または「　」区切りで入力してください。例:アーサー,ランスロット,ガウェイン,モードレット",
			Required:    true,
		},
	},
	DefaultPermission: true,
}

var diceAC = &discord.ApplicationCommand{
	Name:        "r",
	Description: "ダイスを振る",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "parm",
			Description: "1d100のように指定してください。例:「1d100」「2d6」",
			Required:    true,
		},
	},
	DefaultPermission: true,
}

var helpSlashCommand = disgoslash.NewSlashCommand(helpAC, command.HelpCommand, Global, GuildIDs)
var selectSlashCommand = disgoslash.NewSlashCommand(selectAC, command.SelectCommand, Global, GuildIDs)
var orderSlashCommand = disgoslash.NewSlashCommand(orderAC, command.OrderCommand, Global, GuildIDs)
var diceSlashCommand = disgoslash.NewSlashCommand(diceAC, command.DiceCommand, Global, GuildIDs)

// SlashCommandsMap is exported for use with the sync package which will
// register the slash command on Discord.
var SlashCommandMap = disgoslash.NewSlashCommandMap(helpSlashCommand, selectSlashCommand, orderSlashCommand, diceSlashCommand)

// Handler acts as the entrypoint for slash command requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	handler := &disgoslash.Handler{SlashCommandMap: SlashCommandMap, Creds: GetCredentials()}
	handler.Handle(w, r)
}
