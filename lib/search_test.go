package alphaYon

import (
	"testing"
)

type searchTest struct {
	coordList []Coord
	depth     int
	status    ABNodeStatus
}

var searchToTest = []searchTest{
	searchTest{
		[]Coord{
			Coord{0, 0},
			Coord{1, 1},
			Coord{0, 1},
			Coord{2, 1},
			Coord{0, 2},
			Coord{2, 2},
			Coord{1, 0},
			Coord{1, 2},
			Coord{2, 0}}, 1, UNKNOWN},

	searchTest{
		[]Coord{
			Coord{0, 0},
			Coord{1, 1},
			Coord{0, 1},
			Coord{2, 1},
			Coord{0, 2},
			Coord{2, 2},
			Coord{1, 0},
			Coord{1, 2},
			Coord{2, 0}}, 3, WINNING},

	searchTest{
		[]Coord{
			Coord{3, 3},
			Coord{0, 0},
			Coord{1, 1},
			Coord{0, 1},
			Coord{2, 1},
			Coord{0, 2},
			Coord{2, 2},
			Coord{1, 0},
			Coord{1, 2}}, 3, LOSING},
}

func TestSearch(t *testing.T) {
	for _, test := range searchToTest {
		player := WHITE
		game := NewGame(player, 4)
		node := NewABNode(game)

		for _, coord := range test.coordList {
			err := game.Move(coord.X, coord.Y)

			if err != nil {
				t.Errorf("Setup error: %s.", err)
			}
		}

		status, _ := Search(node, player, test.depth)

		if *status != test.status {
			t.Errorf("Search(%s) = %s, want %s.", node.Game.Board, *status, test.status)
		}
	}
}
