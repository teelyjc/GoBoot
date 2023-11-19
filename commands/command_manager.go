package commands

import (
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

}

func (cm *CommandManager) ExecuteCommand(
	session *discordgo.Session,
	interaction *discordgo.InteractionCreate,
) {

}
