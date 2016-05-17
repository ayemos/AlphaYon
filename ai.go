package main

import (
	"fmt"
	"math"
)

type AI struct {
	Game        *Game
	MctsC       float64
	Tree        *Tree
	CurrentNode *Node
}

func (ai *AI) solve(timeLimit int) (int, int) {
	// TODO: handle timeLimit

	// use mcts to calculate next move
	coord := ai.CurrentNode.mcts()

	return coord.X, coord.Y
}

// TODO: time limited mcts
func (node *Node) mcts() (coord Coord) {
	fmt.Println("mcts")
	fmt.Println(node)
	node.expandChildren()
	fmt.Println(node)

	fmt.Println(node)

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
		winner = Judge(node.Game.Board)

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
	return BLACK
}

func (node *Node) expandChildren() {
	for range node.Game.Board.Frees {
		/*
			newBoard := node.Game.Board
		*/

		newBoard := NewBoard(node.Game.Board.Radius)
		CopyBoard(node.Game.Board, newBoard)

		newGame := &Game{
			Winner: node.Game.Winner,
			Turn:   node.Game.Turn.nextTurn(),
			Board:  newBoard,
		}

		newNode := NewNode(newGame)

		node.Children = append(node.Children, newNode)
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

	return nil
}

func (node *Node) tryRandomMove() {
	// node.move(rand.Intn(node.Radius), rand.Intn(node.Radius))
}

func (ai *AI) mctsFactor(node *Node, n int) float64 {
	// TODO: make it fast
	return float64(node.Wins/node.Trials) +
		float64(ai.MctsC)*math.Sqrt(math.Log(float64(n))/float64(node.Trials))
}

func NewAI(game *Game, mctsC float64) *AI {
	tree := NewTree(game)

	ai := &AI{
		Game:        game,
		MctsC:       mctsC,
		Tree:        tree,
		CurrentNode: tree.Root,
	}

	return ai
}
