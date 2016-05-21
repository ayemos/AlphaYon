package alphaYon

import (
	"fmt"
)

type MCTSRecord struct {
	Wins, Draws, Trials int
}

type MctsTree struct {
	Root *MctsNode
}

type MctsNode struct {
	Game     *Game
	Children []*MctsNode
	Parent   *MctsNode
	MCTSRecord
	Coord
	Played bool
}

func NewMctsNode(game *Game) *MctsNode {
	node := &MctsNode{
		Game:       game,
		Children:   []*MctsNode{},
		Parent:     nil,
		MCTSRecord: MCTSRecord{0, 0, 0},
		Coord:      Coord{-1, -1},
		Played:     false,
	}

	return node
}

func (node *MctsNode) root() *MctsNode {
	if node.Parent == nil {
		return node
	} else {
		return node.Parent.root()
	}
}

func (node *MctsNode) appendChild(child *MctsNode) error {
	child.Parent = node
	node.Children = append(node.Children, child)
	return nil
}

func NewMctsTree(game *Game) *MctsTree {
	root := NewMctsNode(game)

	tree := &MctsTree{
		Root: root,
	}

	return tree
}

func repMctsNode(n MctsNode, depth int) string {
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

	for i := 0; i < depth; i++ {
		str = append(str, fmt.Sprintf("\t")...)
	}

	str = append(str, fmt.Sprintf("\tMCTS: %.2f\n", n.mctsFactor(n.root().Trials, 0.3))...)
	str = append(str, fmt.Sprintf(hDump(*n.Game.Board, depth+1))...)

	str = append(str, ")\n"...)
	return string(str)
}

func (n MctsNode) buildString(depth int) string {
	str := make([]byte, 0)

	for i := 0; i < depth; i++ {
		str = append(str, '\t')
	}

	str = append(str, repMctsNode(n, depth)...)

	for _, child := range n.Children {
		str = append(str, child.buildString(depth+1)...)
	}

	return string(str)
}

func (n MctsNode) String() string {
	str := make([]byte, 0)
	str = append(str, n.buildString(0)...)
	return string(str)
}
