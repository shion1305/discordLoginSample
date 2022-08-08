package discordLoginSample

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("discord_token"))
	err = dg.Open()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}
	dg.AddHandler(onMessageCreate)
	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot
}
