package main

import (
//	"fmt"
)

type AI struct {
	Game     *Game
	MctsC    float64
	MctsTree *MctsTree
}

func (ai *AI) solve(player Color, timeLimit int) (int, int) {
	// TODO: handle timeLimit

	// create root node according to game state
	// search from the root node
	root := NewMctsNode(ai.Game)

	// use mcts to calculate next move
	Mcts(root, player, ai.MctsC, 500, timeLimit)

	// choose child who has max score
	var maxChild *MctsNode
	// TODO: Consider other impl
	var maxTrials = 0

	for _, child := range root.Children {
		/*
			fmt.Println("Trials")
			fmt.Println(child.Trials)
		*/
		if child.Trials > maxTrials {
			maxTrials = child.Trials
			maxChild = child
		}
	}

	b := maxChild.Game.Board
	coord := b.History[b.Turns-1]
	/*
		fmt.Println("maxChild")
		fmt.Println(maxChild)
		fmt.Println(b.Turns)
		fmt.Println(b.History)
	*/

	return coord.X, coord.Y
}

func NewAI(game *Game, mctsC float64) *AI {
	tree := NewMctsTree(game)

	ai := &AI{
		Game:     game,
		MctsC:    mctsC,
		MctsTree: tree,
	}

	return ai
}
