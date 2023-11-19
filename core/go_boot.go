package core

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/teelyjc/GoBoot/commands"
	"github.com/teelyjc/GoBoot/handlers"
)

type goBoot struct {
	session *discordgo.Session
}

func NewGoBoot(TOKEN string) *goBoot {
	session, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatalln("Error creating discord session, ", err)
	}

	return &goBoot{
		session: session,
	}
}

func (gb *goBoot) Start() {
	commandManager := commands.NewCommandManager(gb.session)

	gb.session.AddHandler(handlers.ReadyHandler)
	gb.session.AddHandler(
		func(
			session *discordgo.Session,
			interaction *discordgo.InteractionCreate,
		) {
			commandManager.ExecuteCommand(session, interaction)
		},
	)

	if err := gb.session.Open(); err != nil {
		log.Fatalln("Error opening connection, ", err)
	}
}

func (gb *goBoot) Stop() {
	if err := gb.session.Close(); err != nil {
		log.Fatalln("Error closing connection, ", err)
	} else {
		log.Println("Successfully shutdown.")
	}
}
