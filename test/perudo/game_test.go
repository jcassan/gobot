package perudo

import (
	"perubot/pkg/perudo"
	"testing"
)

func TestCreateGame(t *testing.T) {
	var players = []perudo.Player{
		{Name: "Tata3"},
		{Name: "Toto1"},
		{Name: "Tata1"},
		{Name: "Tata2"},
	}
	var game, player = perudo.CreateGame(players)
	if len(game.Players) != 4 {
		t.Errorf("Player count invalid. Actual value : %d, Expected value : 4.", len(game.Players))
	}
}
