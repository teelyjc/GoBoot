package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/teelyjc/GoBoot/commands"
	command "github.com/teelyjc/GoBoot/commands/common"
)

func ReadyHandler(
	session *discordgo.Session,
	ready *discordgo.Ready,
	commandManager *commands.CommandManager,
) {
	log.Printf(
		"Logged in as %v#%v",
		session.State.User.Username,
		session.State.User.Discriminator,
	)

	commandManager.RegisterCommand(command.CommandPing())
}
