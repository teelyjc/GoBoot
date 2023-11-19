package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/teelyjc/GoBoot/commands"
)

func InteractionHandler(
	session *discordgo.Session,
	interaction *discordgo.InteractionCreate,
	commandManager *commands.CommandManager,
) {
	commandManager.ExecuteCommand(session, interaction)
}
