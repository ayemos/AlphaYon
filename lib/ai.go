package alphaYon

import (
	"fmt"
)

type AI struct {
	Game     *Game
	MctsC    float64
	MctsT    int
	MctsTree *MctsTree
}

const (
	// TODO: Comment
	DefaultTimeLimit   = 1
	DefaultMCTSC       = 0.5
	DefaultMCTST       = 500
	DefaultSearchDepth = 5
)

func NewAI(game *Game, mctsC float64, mctsT int) *AI {
	tree := NewMctsTree(game)

	ai := &AI{
		Game:     game,
		MctsC:    mctsC,
		MctsT:    mctsT,
		MctsTree: tree,
	}

	return ai
}

func (ai *AI) Solve(player Color, timeLimit int, searchDepth int) (int, int) {
	// TODO: handle timeLimit

	// create root node according to game state
	// search from the root node
	root := NewMctsNode(ai.Game)

	// use mcts to calculate next move
	Mcts(root, player, ai.MctsC, ai.MctsT, timeLimit, searchDepth)
	fmt.Println(root)

	// choose child who has max score
	var maxChild *MctsNode
	// TODO: Consider other impl
	var maxTrials = 0

	for _, child := range root.Children {
		if child.Trials >= maxTrials {
			maxTrials = child.Trials
			maxChild = child
		}
	}

	b := maxChild.Game.Board
	coord := b.History[b.Turns-1]

	return coord.X, coord.Y
}
