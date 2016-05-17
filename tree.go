package main

import (
	"fmt"
)

type MCTSRecord struct {
	Wins, Trials int
}

type Tree struct {
	Root *Node
}

type Node struct {
	Game     *Game
	Children []*Node
	Parent   *Node
	MCTSRecord
	Played bool
}

func NewNode(game *Game) *Node {
	node := &Node{
		Game:       game,
		Children:   []*Node{},
		Parent:     nil,
		MCTSRecord: MCTSRecord{0, 0},
		Played:     false,
	}

	return node
}

func NewTree(game *Game) *Tree {
	root := NewNode(game)

	tree := &Tree{
		Root: root,
	}

	return tree
}

func repNode(n Node, depth int) string {
	str := make([]byte, 1)
	str = append(str, "Node(\n"...)

	var yn string

	if n.Played {
		yn = "Yes"
	} else {
		yn = "No"
	}

	str = append(str, fmt.Sprintf("\tPlayed: %s\n", yn)...)
	str = append(str, fmt.Sprintf("\tWins: %d, Trials: %d, Loses: %d\n",
		n.Wins, n.Trials, n.Trials-n.Wins)...)
	str = append(str, fmt.Sprintf(hDump(*n.Game.Board, depth))...)

	str = append(str, ")\n"...)
	return string(str)
}

func buildString(n Node, depth int) string {
	str := make([]byte, 1)

	for i := 0; i < depth; i++ {
		str = append(str, '\t')
	}

	str = append(str, repNode(n, depth)...)

	for _, child := range n.Children {
		str = append(str, buildString(*child, depth+1)...)
	}

	return string(str)
}

func (n Node) String() string {
	str := make([]byte, 1)
	str = append(str, buildString(n, 0)...)
	return string(str)
}
