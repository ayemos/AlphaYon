package main

import (
	"testing"
)

func TestJudge(t *testing.T) {
	b := NewBoard(4)

	winner := Judge(b)
	if winner != EMPTY {
		t.Errorf("Judge(%s) = %s, want %s.", b, winner, "EMPTY")
	}
}

