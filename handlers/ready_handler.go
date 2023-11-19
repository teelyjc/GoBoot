package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(session *discordgo.Session, ready *discordgo.Ready) {
	log.Printf(
		"Logged in as %v#%v",
		session.State.User.Username,
		session.State.User.Discriminator,
	)
}
