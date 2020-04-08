package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "Discord API Token")
	flag.Parse()

	if DiscordToken == "" {
		flag.Usage()
		os.Exit(1)
	}
}

var DiscordToken string

func main() {
	mom, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		fmt.Println("==MomBotGO Error==\n" + err.Error())
		return
	}

	mom.AddHandler(readyHandler)
	mom.AddHandler(responder)

	err = mom.Open()
	if err != nil {
		fmt.Println("==MomBotGO Error==\n" + err.Error())
		return
	}

	fmt.Println("Bot is now running. Press CTRL-c to exit.")
	sc := make (chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	mom.Close()
}

func readyHandler(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateStatus(0, "Cooking Mama")
	return
}

func sendResponse(session *discordgo.Session, event *discordgo.MessageCreate, message string) {
	_, err := session.ChannelMessageSend(event.ChannelID, message)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func responder(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}
	matchMom, err := regexp.MatchString(`\bmom\b|\bmum\b|\bmommy\b|\bmummy\b|\bmother\b/gim`, event.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
	matchQuestion, err := regexp.MatchString(`[a-z ]\?\B|^\?+$/gim`, event.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(event.Content+" ")
	fmt.Println("matchQuestion: " + strconv.FormatBool(matchQuestion) + " matchMom: " + strconv.FormatBool(matchMom))
	if matchMom && matchQuestion {
		sendResponse(session, event, "Mhm..")
		return
	}
	if matchQuestion {
		sendResponse(session, event, "Ask your father")
		return
	}
	if matchMom {
		sendResponse(session, event, "Not now sweetie")
		return
	}
}
