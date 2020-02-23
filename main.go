package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"flag"
	"sort"
	"math/rand"
)

// Variables used for command line parameters
var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	fmt.Println(Token)
}
// Generic types for Perudooooooooo 
type Player struct {
	gorm.Model
	ID string
	Name string
	Dices []int
	DicesCount int

}

type Bet struct {
	DiceValue int
	DiceOccurence int	
}

type Game struct {
	gorm.Model
	Players []Player
	CurrentPlayer Player
	LastBet Bet	
}

//Specific types for Discord
type DiscordPlayer struct {
	Player
	PrivateChannel string
}

type DiscordGame struct {
	Game
	GameChannel string
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.

func rollDices(players []Player){
	for i:=0; i < len(players); i++ {
		players[i].Dices=make([]int, players[i].DicesCount)
		for j:=0; j < players[i].DicesCount; j++ {
			players[i].Dices[j]=rand.Intn(5)+1
		}
	}
}

func createGame(players []Player) Game{
	sort.Slice(players, func(i, j int) bool {
		return players[i].Name < players[j].Name
	})
	var game = Game{
		Players: players,
		LastBet: Bet{0, 0},		
	}
	game.CurrentPlayer=game.Players[0]
	rollDices(game.Players)
	return game
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
}
