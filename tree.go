package main

import (
	"fmt"
)

type MCTSRecord struct {
	Wins, Draws, Trials int
}

type Tree struct {
	Root *Node
}

type Node struct {
	Game     *Game
	Children []*Node
	Parent   *Node
	MCTSRecord
	Coord
	Played bool
}

func NewNode(game *Game) *Node {
	node := &Node{
		Game:       game,
		Children:   []*Node{},
		Parent:     nil,
		MCTSRecord: MCTSRecord{0, 0, 0},
		Coord:      Coord{-1, -1},
		Played:     false,
	}

	return node
}

func (node *Node) root() *Node {
	if node.Parent == nil {
		return node
	} else {
		return node.Parent.root()
	}
}

func (node *Node) appendChild(child *Node) error {
	child.Parent = node
	node.Children = append(node.Children, child)
	return nil
}

func NewTree(game *Game) *Tree {
	root := NewNode(game)

	tree := &Tree{
		Root: root,
	}

	return tree
}

func repNode(n Node, depth int) string {
	str := make([]byte, 0)
	str = append(str, "Node(\n"...)

	var yn string

	if n.Played {
		yn = "Yes"
	} else {
		yn = "No"
	}

	for i := 0; i < depth; i++ {
		str = append(str, fmt.Sprintf("\t")...)
	}

	str = append(str, fmt.Sprintf("\tPlayed: %s\n", yn)...)

	for i := 0; i < depth; i++ {
		str = append(str, fmt.Sprintf("\t")...)
	}

	str = append(str, fmt.Sprintf("\tWins: %d, Draws: %d, Loses: %d, Trials: %d\n",
		n.Wins, n.Draws, n.Trials-n.Wins, n.Trials)...)
	str = append(str, fmt.Sprintf("\tMCTS: %.2f\n", n.mctsFactor(n.root().Trials, 0.3))...)
	str = append(str, fmt.Sprintf(hDump(*n.Game.Board, depth+1))...)

	str = append(str, ")\n"...)
	return string(str)
}

func buildString(n Node, depth int) string {
	str := make([]byte, 0)

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
	str := make([]byte, 0)
	str = append(str, buildString(n, 0)...)
	return string(str)
}
