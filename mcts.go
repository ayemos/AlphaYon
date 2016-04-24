package main

import (
	"math"
	"math/rand"
)

func (node *Node) mcts(n int) (coord Coord) {
	node.expandChildren()

}

func (node *Node) playout() (winner Color) {
	for {
		winner = node.Board.judge()

		// TODO: winner is evaluated twice
		if winner == WHITE || winner == BLACK {
			return winner
		}


		for f := range node.Frees {

		}
	}
}

func (node *Node) expandChildren() {
	for f := range node.Frees {
		var newGame Game
		copy(newGame, node.Game, Game)
		newNode := &Node{
			Game: newGame,
			Children: make([]Node),
			MCTSRecrod: &MCTSRecord{0, 0},
			Played: false
		}
		node.Children = append(node.Children,
	}
}

func (node *Node) tryRandomMove() {
	node.move(rand.Intn(node.Radius), rand.Intn(node.Radius))
}

func (b *Board) judge() (winner Color) {
	return EMPTY
}

func mctsFactor(node *Node) float64 {
	// TODO: fix
	return float64(node.Wins/node.Trials) +
		MCTS_C*math.Sqrt(math.Log(10)/float64(node.Trials))
}
