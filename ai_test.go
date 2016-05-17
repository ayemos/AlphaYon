package main

import (
	"testing"
)

func TestAI(t *testing.T) {
	game := NewGame(BLACK, 4)
	ai := NewAI(game, 0.3)
}
