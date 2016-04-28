package main

import (
	"math"
	"math/rand"
)

func (node *Node) mcts(n int) (coord Coord) {
	node.expandChildren()

	for {

	}
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

func (node *Node) playoutInitialNode() (winner Color) {

}

func (node *Node) expandChildren() {
	for range node.Frees {
		newGame := node.Game

		newNode := &Node{
			Game:       newGame,
			Children:   make([]Node, 1),
			MCTSRecord: MCTSRecord{0, 0},
			Played:     false,
		}
		node.Children = append(node.Children, *newNode)
	}
}

func (node *Node) tryRandomMove() {
	node.move(rand.Intn(node.Radius), rand.Intn(node.Radius))
}

func (b *Board) judge() (winner Color) {
	return EMPTY
}

func mctsFactor(node *Node, n int) float64 {
	// TODO: make it fast
	return float64(node.Wins/node.Trials) +
		MCTS_C*math.Sqrt(math.Log(float64(n))/float64(node.Trials))
}
