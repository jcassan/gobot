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
	ID           int
	Name         string
	Dices        []int
	DicesCount   int
	IsEliminated bool
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

func CreateGame(players []Player) (Game, Player) {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Name < players[j].Name
	})
	var game = Game{
		Players: players,
		LastBet: Bet{0, 0},
	}
	game.CurrentPlayer = game.Players[0]
	RollDices(game.Players)
	for i := 0; i < len(game.Players); i++ {
		game.Players[i].ID = i + 1
	}
	return game, game.Players[0]
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

func getPreviousPlayer(players []Player, currentPlayer Player) Player {
	var index = FindPlayerIndex(players, currentPlayer)
	var found bool = false
	var i = index - 1
	for found != true && i != index {
		if i < 0 {
			i = len(players)
		}
		if players[i].IsEliminated == false {
			found = true
		} else {
			i--
		}
	}
	return players[i]
}

func getNextPlayer(players []Player, currentPlayer Player) Player {
	var index = FindPlayerIndex(players, currentPlayer)
	var found bool = false
	var i = (index + 1) % len(players)
	for found != true && (i%len(players)) != index {
		if players[i].IsEliminated == false {
			found = true
		} else {
			i++
		}
	}
	return players[i]
}

//return the player who lost the round
func endRound(game Game) Player {
	//Count total number of DiceValue in the round
	var DiceCounter int = 0
	for i := 0; i < len(game.Players); i++ {
		if game.Players[i].IsEliminated == false {
			for j := 0; j < game.Players[i].DicesCount; j++ {
				if game.LastBet.DiceValue == game.Players[i].Dices[j] {
					DiceCounter++
				}
			}
		}
	}
	if game.LastBet.DiceOccurence > DiceCounter {
		return game.CurrentPlayer
	}
	return getPreviousPlayer(game.Players, game.CurrentPlayer)
}

/*
PlayRound
Return true and the player who lost if a player said stop,
Return false and the next Player if the bets continue
Return true and the player who bet if the bet was invalid
*/
func PlayRound(game Game, bet Bet) (bool, Player, error) {
	if bet.DiceOccurence == -1 && bet.DiceValue == -1 {
		return true, endRound(game), nil
	}
	err := CheckBet(game.LastBet, bet)
	if err != nil {
		game.LastBet = bet
		game.CurrentPlayer = getNextPlayer(game.Players, game.CurrentPlayer)
		return false, game.CurrentPlayer, nil
	} else {
		return true, game.CurrentPlayer, err
	}

}

func FindPlayerIndex(players []Player, p Player) int {
	for i, n := range players {
		if p.ID == n.ID {
			return i
		}
	}
	return -1
}
