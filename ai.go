package main

import (
	"fmt"
	"math"
	"math/rand"
)

type AI struct {
	Game        *Game
	MctsC       float64
	Tree        *Tree
	CurrentNode *Node
}

func (ai *AI) solve(player Color, timeLimit int) (int, int) {
	// TODO: handle timeLimit

	// create root node according to game state
	// search from the root node
	root := NewNode(ai.Game)
	root.expandChildren()

	// use mcts to calculate next move
	mcts(root, player, ai.MctsC)

	// choose child who has max score
	// for child := root.Children { ...

	return 0, 0
}

// TODO: time limited mcts
func mcts(root *Node, player Color, mctsC float64) {
	fmt.Println("mcts")
	fmt.Println(root)

	var maxNode, node *Node
	node = root

	for {
		maxNode = chooseMaxNode(node, mctsC)

		fmt.Println(node)
		// Choose the node has maximum mcts factor

		// playout
		winner = maxNode.playout()
		win = winner == root.Game.Turn

		// backpropagate
		for {
			maxNode
		}
	}
}

func chooseMaxNode(root *Node, mctsC float64) *Node {
	if len(root.Children) == 0 {
		return root
	}

	maxNode := root
	maxScore := 0.0

	for _, child := range root.Children {
		// Playout every initial nodes first
		if !child.Played {
			return child
		}

		score := mctsFactor(child, root.Trials, mctsC)

		if score > maxScore {
			maxNode = child
			maxScore = score
		}
	}

	return maxNode
}

func shouldExpand(node *Node) bool {
	return false
}

func (node *Node) playout() (winner Color) {
	board := CopyBoard(node.Game.Board)

	fmt.Println("playout on")
	fmt.Println(board)

	for {
		winner = Judge(node.Game.Board)

		if winner != EMPTY {
			fmt.Printf("%s won!\n", winner)

			return winner
		}

		tryRandomMove(node)
		fmt.Println(board)
	}
}

func (node *Node) expandChildren() {
	for i := range node.Game.Board.Frees {
		newBoard := CopyBoard(node.Game.Board)

		newGame := &Game{
			Winner: node.Game.Winner,
			Turn:   node.Game.Turn,
			Board:  newBoard,
		}

		newGame.moveFree(i)

		newNode := NewNode(newGame)

		node.Children = append(node.Children, newNode)
	}
}

func (node *Node) initialNodes() (nodes []*Node) {
	nodes = make([]*Node, 0)

	if !node.Played {
		nodes = append(nodes, node)
	}

	for _, child := range node.Children {
		for _, initialChild := range child.initialNodes() {
			nodes = append(nodes, initialChild)
		}
	}

	return nodes
}

func tryRandomMove(node *Node) {
	game := node.Game
	game.moveFree(rand.Intn(node.Game.Board.FreesCount))
}

func mctsFactor(node *Node, n int, mctsC float64) float64 {
	if node.Trials == 0 {
		return 0.0
	}

	// TODO: make it fast
	return float64(node.Wins/node.Trials) +
		mctsC*math.Sqrt(math.Log(float64(n))/float64(node.Trials))
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
