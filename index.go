package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	session *discordgo.Session
)

func main() {
	log.Info("Starting UltraHooks Bot...")

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Error("Error creating Discord session ", err.Error())
		panic("Error creating Discord session " + err.Error())
	}

	// Open a websocket connection to Discord and begin listening.
	log.Info("Launching Shard-", dg.ShardID)
	err = dg.Open()
	if err != nil {
		log.Error("Error opening connection ", err.Error())
		panic("Error opening connection " + err.Error())
	}

	// Launch web server
	go server()

	session = dg

	log.Info("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info("Stopping bot...")
	dg.Close()
}
