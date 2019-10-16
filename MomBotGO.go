package MomBotGO

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&DiscordToken, "t", "", "Discord API Token")
	flag.Parse()
}

var DiscordToken string

func main() {
	if DiscordToken == "" {
		fmt.Println("==MomBotGO Error==\nNo API Token Provided")
		os.Exit(1)
	}
	mom, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		fmt.Println("==MomBotGO Error==\n" + err.Error())
		os.Exit(1)
	}

	mom.AddHandler(readyHandler)
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
	matchMom, err := regexp.MatchString("`\b(mom|mum|mommy|mummy|mother)\b`gmi", event.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
	matchQuestion, err := regexp.MatchString("`([a-z]\\?\\B|^\\?+$)`img", event.Content)
	if err != nil {
		fmt.Println(err)
		return
	}
	if matchQuestion == true {
		sendResponse(session, event, "Ask your father")
		return
	}
	if matchMom == true {
		sendResponse(session, event, "Not now sweetie")
		return
	}
}