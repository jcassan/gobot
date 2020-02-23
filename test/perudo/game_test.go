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
	var game, nextPlayer = perudo.CreateGame(players)
	if len(game.Players) != 4 {
		t.Errorf("Player count invalid, expected: 4, got %d.", len(game.Players))
	}
	if nextPlayer.Name != "Tata1" {
		t.Errorf("Next payer invalid, expected: 'Tata1', got \"%s\".", nextPlayer.Name)
	}
}
