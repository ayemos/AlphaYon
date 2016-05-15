package main

import (
	"testing"
)


type boardTest struct {
	pins []rune
	winner Color
}

var boardsToTest = []boardTest{
	boardTest{
		[]rune{
			'b', '.', '.', '.',
			'b', '.', '.', '.',
			'b', '.', '.', '.',
			'b', '.', '.', '.',

			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.',

			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.',

			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.',
			'.', '.', '.', '.'}, BLACK},
}

func TestJudge(t *testing.T) {
	for _, test := range boardsToTest {
		b := NewBoard(4)
		err := b.loadArray(test.pins)

		if err != nil {
			t.Errorf("Setup error: %s.", err)
		}

		winner := Judge(b)
		if winner != test.winner {
			t.Errorf("Judge(%s) = %s, want %s.", b, winner, test.winner)
		}
	}
}
