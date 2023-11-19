package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Info    *discordgo.ApplicationCommand
	Handler func(session *discordgo.Session, interaction *discordgo.Interaction)
}

type CommandManager struct {
	session     *discordgo.Session
	commandsMap map[string]Command
}

func NewCommandManager(session *discordgo.Session) *CommandManager {
	return &CommandManager{
		session:     session,
		commandsMap: make(map[string]Command),
	}
}

func (cm *CommandManager) RegisterCommand(command *Command) {
	cm.commandsMap[command.Info.Name] = *command
	cm.session.ApplicationCommandCreate(
		cm.session.State.User.ID,
		"",
		command.Info,
	)

	log.Printf(
		"Register command %v in-memory successfully.",
		command.Info.Name,
	)
}

func (cm *CommandManager) ExecuteCommand(
	session *discordgo.Session,
	interaction *discordgo.InteractionCreate,
) {
	if command, ok := cm.commandsMap[interaction.ApplicationCommandData().Name]; ok {
		command.Handler(session, interaction.Interaction)
	}
}
