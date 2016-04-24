package main

import (
	"math/rand"
)

func (node *Node) playout() (winner Color) {
	for {
		winner = node.Board.judge()

		// TODO: winner is evaluated twice
		if winner == WHITE || winner == BLACK {
			return winner
		}

		node.tryRandomMove()
	}
}

func (node *Node) tryRandomMove() {
	node.move(rand.Intn(node.Radius), rand.Intn(node.Radius))
}

func (b *Board) judge() (winner Color) {
	return EMPTY
}
