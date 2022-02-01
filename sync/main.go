package main

import (
	"github.com/fumiya007/tpDisSlash/api"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/disgoslash"
)

func main() {
	godotenv.Load()

	syncer := &disgoslash.Syncer{
		Creds:           api.GetCredentials(),
		GuildIDs:        api.GuildIDs,
		SlashCommandMap: api.SlashCommandMap,
	}
	syncer.Sync()
}
