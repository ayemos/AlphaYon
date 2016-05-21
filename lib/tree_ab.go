package alphaYon

import (
	"fmt"
)

type ABTree struct {
	Root *ABNode
}

type ABNode struct {
	Game     *Game
	Children []*ABNode
	Parent   *ABNode
	Coord
}

func NewABNode(game *Game) *ABNode {
	node := &ABNode{
		Game:     game,
		Children: []*ABNode{},
		Parent:   nil,
		Coord:    Coord{-1, -1},
	}

	return node
}

func (node *ABNode) root() *ABNode {
	if node.Parent == nil {
		return node
	} else {
		return node.Parent.root()
	}
}

func (node *ABNode) appendChild(child *ABNode) error {
	child.Parent = node
	node.Children = append(node.Children, child)
	return nil
}

func NewABTree(game *Game) *ABTree {
	root := NewABNode(game)

	tree := &ABTree{
		Root: root,
	}

	return tree
}

func repABNode(n ABNode, depth int) string {
	str := make([]byte, 0)
	str = append(str, "Node(\n"...)

	for i := 0; i < depth; i++ {
		str = append(str, fmt.Sprintf("\t")...)
	}

	str = append(str, fmt.Sprintf(hDump(*n.Game.Board, depth+1))...)

	str = append(str, ")\n"...)
	return string(str)
}

func (n ABNode) buildString(depth int) string {
	str := make([]byte, 0)

	for i := 0; i < depth; i++ {
		str = append(str, '\t')
	}

	str = append(str, repABNode(n, depth)...)

	for _, child := range n.Children {
		str = append(str, child.buildString(depth+1)...)
	}

	return string(str)
}

func (n ABNode) String() string {
	str := make([]byte, 0)
	str = append(str, n.buildString(0)...)
	return string(str)
}
