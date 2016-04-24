package main

import (
	"math"
	"math/rand"
)

func (node *Node) mcts(n) (coord Coord) {

}

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

func mctsFactor(node) float32 {
	(node.Wins / node.Trials) + MCTS_C * math.Sqrt(math.Log(n) / node.Trials)
}
