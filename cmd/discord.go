package main

import (
	"flag"
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"perubot/pkg/perudo"
	"syscall"
)
var games map[string]DiscordGame 
// Variables used for command line parameters
var Token string

//Specific types for Discord
type DiscordPlayer struct {
	player perudo.Player
	PrivateChannel string
}

type DiscordGame struct {
	game perudo.Game
//	GameChannel string
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	fmt.Println(Token)
}

func convertMentions(users []*discordgo.User) []DiscordPlayer {
	var players []DiscordPlayer
	for _, u := range users {
		var player DiscordPlayer = DiscordPlayer{}
		player.player.ID = u.ID
		player.player.Name = u.Username
		player.PrivateChannel = ""
		players = append(players, player)
	}

	return players
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if strings.HasPrefix(m.Content, "start"){
		var users []*discordgo.User = m.Mentions
		games = make(map[string]DiscordGame)

		var discordPlayers []DiscordPlayer = convertMentions(users)
		var perudoPlayers []perudo.Player
		for _, p := range discordPlayers {
			perudoPlayers = append(perudoPlayers, p.player)
		}
		var game, firstPlayer = perudo.CreateGame(perudoPlayers)

		games[m.ChannelID] = DiscordGame{
			game: game,
		}
		
		s.ChannelMessageSend(m.ChannelID, firstPlayer.Name)
	}

}
