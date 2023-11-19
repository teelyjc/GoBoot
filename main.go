package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/teelyjc/GoBoot/core"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")
	if TOKEN == "" {
		log.Fatalln("No token found")
	}

	goBoot := core.NewGoBoot(TOKEN)
	goBoot.Start()

	log.Println("Bot is now running, Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	goBoot.Stop()
}
