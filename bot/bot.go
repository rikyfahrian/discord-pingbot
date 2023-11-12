package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	goBot, err := discordgo.New("Bot " + os.Getenv("TOKEN_BOT"))
	if err != nil {
		log.Fatal("error creating discord session ", err.Error())
	}

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal("error opening connection ", err.Error())
	}
	defer goBot.Close()

	log.Println("Bot is running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
