package perudo

import (
	"github.com/stretchr/testify/require"
	"perubot/pkg/perudo"
	"testing"
)

// Test for CreateGame
func TestCreateGame(t *testing.T) {
	assert := require.New(t)

	var players = []perudo.Player{
		{Name: "Tata3"},
		{Name: "Toto1"},
		{Name: "Tata1"},
		{Name: "Tata2"},
	}
	var game, nextPlayer = perudo.CreateGame(players)

	assert.Len(game.Players, 4, "Player count invalid")
	assert.Condition(
		func() bool {
			return "Tata1" == players[0].Name && "Tata2" == players[1].Name && "Tata3" == players[2].Name && "Toto1" == players[3].Name
		},
		"Players array not sorted")
	assert.Equal("Tata1", nextPlayer.Name, "Next payer invalid")
	assert.Equal(0, game.LastBet.DiceOccurence, "Dice occurence of last bet invalid")
	assert.Equal(0, game.LastBet.DiceValue, "Dice value of last bet invalid")
}

// Test to increment DiceOccurence for CheckBet
func TestIncOccurenceCheckBet(t *testing.T) {
	assert := require.New(t)

	lastBet := perudo.Bet{
		DiceValue:     4,
		DiceOccurence: 2,
	}
	nextBet := perudo.Bet{
		DiceValue:     4,
		DiceOccurence: 3,
	}
	assert.Nil(perudo.CheckBet(lastBet, nextBet), "CheckBet must not return error")
}

// Test to increment DiceValue for CheckBet
func TestIncValueCheckBet(t *testing.T) {
	assert := require.New(t)

	lastBet := perudo.Bet{
		DiceValue:     4,
		DiceOccurence: 2,
	}
	nextBet := perudo.Bet{
		DiceValue:     5,
		DiceOccurence: 2,
	}
	assert.Nil(perudo.CheckBet(lastBet, nextBet), "CheckBet must not return error")
}

// Test for invalid value for CheckBet
func TestInvalidValueCheckBet(t *testing.T) {
	assert := require.New(t)

	lastBet := perudo.Bet{
		DiceValue:     4,
		DiceOccurence: 2,
	}
	nextBet := perudo.Bet{
		DiceValue:     7,
		DiceOccurence: 2,
	}
	assert.EqualError(perudo.CheckBet(lastBet, nextBet), "dice value lower than 0 or greater than 6")
	nextBet = perudo.Bet{
		DiceValue:     -1,
		DiceOccurence: 2,
	}
	assert.EqualError(perudo.CheckBet(lastBet, nextBet), "dice value lower than 0 or greater than 6")
	nextBet = perudo.Bet{
		DiceValue:     5,
		DiceOccurence: 3,
	}
	assert.EqualError(perudo.CheckBet(lastBet, nextBet), "incorrect bet")
}
