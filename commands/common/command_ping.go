package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/teelyjc/GoBoot/commands"
)

func CommandPing() *commands.Command {
	return &commands.Command{
		Info: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "reply with pong !",
		},
		Handler: func(session *discordgo.Session, interaction *discordgo.Interaction) {
			user := interaction.Member.User.ID

			session.InteractionRespond(interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong ! <@" + user + ">",
				},
			})
		},
	}
}
