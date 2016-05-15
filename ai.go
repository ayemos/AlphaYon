package main

import (
	"math"
	"math/rand"
)

type AI struct {
	Game *Game
	MctsC float64
}

func (ai *AI) solve() (int, int) {
	// use mcts to calculate next move
	
	return rand.Intn(ai.Game.Radius), rand.Intn(ai.Game.Radius)
}

// TODO: time limited mcts
func (node *Node) mcts(n int) (coord Coord) {
	node.expandChildren()

	for {
		// Playout every initial nodes first
		for _, initialNode := range node.initialNodes() {
			initialNode.playout()
		}

		var maxNode Node
		// Choose the node has maximum mcts factor


		// playout
		maxNode.playout()

		// expand maxNode if threshold
	}

	return Coord{0, 0}
}

func (node *Node) playout() (winner Color) {
	for {
		winner = Judge(node.Board)

		// TODO: winner is evaluated twice
		if winner == WHITE || winner == BLACK {
			return winner
		}

		/*
		for _, f := range node.Frees {

		}
		*/
	}
}

func (node *Node) playoutInitialNode() (winner Color) {
	return BLACK;
}

func (node *Node) expandChildren() {
	for range node.Frees {
		/*
		newBoard := node.Game.Board
		*/
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

func (node *Node) initialNodes() (nodes []Node) {
	// TODO: 10
	nodes = make([]Node, 10)
	if !node.Played {
		nodes = append(nodes, *node)
	}

	for _, child := range node.Children {
		for _, initialChild := range child.initialNodes() {
			nodes = append(nodes, initialChild)
		}
	}

	return nil;
}

func (node *Node) tryRandomMove() {
	// node.move(rand.Intn(node.Radius), rand.Intn(node.Radius))
}

func (ai *AI) mctsFactor(node *Node, n int) float64 {
	// TODO: make it fast
	return float64(node.Wins/node.Trials) +
		float64(ai.MctsC) * math.Sqrt(math.Log(float64(n))/float64(node.Trials))
}

func NewAI(game *Game, mctsC float64) *AI {
	ai := &AI{
		Game: game,
		MctsC: mctsC,
	}

	return ai
}
