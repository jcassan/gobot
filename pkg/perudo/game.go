package perudo

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"sort"
)

// Generic types for Perudooooooooo 
type Player struct {
	gorm.Model
	ID         string
	Name       string
	Dices      []int
	DicesCount int
}

type Bet struct {
	DiceValue     int
	DiceOccurence int
}

type Game struct {
	gorm.Model
	Players       []Player
	CurrentPlayer Player
	LastBet       Bet
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.

func RollDices(players []Player) {
	for i := 0; i < len(players); i++ {
		players[i].Dices = make([]int, players[i].DicesCount)
		for j := 0; j < players[i].DicesCount; j++ {
			players[i].Dices[j] = rand.Intn(5) + 1
		}
	}
}

func CreateGame(players []Player) Game {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Name < players[j].Name
	})
	var game = Game{
		Players: players,
		LastBet: Bet{0, 0},
	}
	game.CurrentPlayer = game.Players[0]
	RollDices(game.Players)
	return game
}

func CheckBet(lastBet Bet, newBet Bet) error {
	if newBet.DiceValue > 6 || newBet.DiceValue < 0 {
		return errors.New("Dice value lower than 0 or greater than 6")
	}
	if newBet.DiceValue > lastBet.DiceValue && newBet.DiceOccurence == lastBet.DiceOccurence {
		return nil
	}
	if newBet.DiceValue == lastBet.DiceValue && newBet.DiceOccurence > lastBet.DiceOccurence {
		return nil
	}
	return errors.New("Incorrect Bet")
}

func PlayRound() {
	var playerBet Bet
	for playerBet != {
		-1, -1
	}
	{

	}
}
