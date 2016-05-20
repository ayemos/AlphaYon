package main

import (
	//	"fmt"
	"math"
	"math/rand"
	"time"
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
	root.expandChildren()

	// use mcts to calculate next move
	mcts(root, player, ai.MctsC, 500, timeLimit)

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

// TODO: time limited mcts
func mcts(root *MctsNode, player Color, mctsC float64, mctsT int, timelimitSec int) {
	/*
		fmt.Println("mcts")
		fmt.Println(root.Game.Turn)
		fmt.Println(player)
	*/

	var maxNode, node *MctsNode
	var winner Color
	var win, draw int
	var diffSec float64

	startTime := time.Now()

	//for i := 0; i < 1000; i++ {
	for i := 0; ; i++ {
		if i%100 == 0 {
			diffSec = time.Now().Sub(startTime).Seconds()

			if diffSec >= float64(timelimitSec) {
				break
			}
		}

		maxNode = chooseMaxNode(root, mctsC)

		// playout
		winner = maxNode.playout()
		if winner == root.Game.Turn {
			win = 1
			draw = 0
		} else if winner == EMPTY {
			win = 0
			draw = 1
		} else {
			win = 0
			draw = 0
		}

		// backpropagate
		node = maxNode

		for {
			node.Wins += win
			node.Draws += draw
			node.Trials += 1

			if node.Parent == nil {
				break
			}

			node = node.Parent
		}

		if shouldExpand(maxNode) {
			maxNode.expandChildren()
		}
	}
}

func shouldExpand(n *MctsNode) bool {
	// Use Alpha-Beta to decide.

	return len(n.Children) == 0 &&
		n.Trials >= mctsT &&
		n.Game.FreesCount > 0 &&
		Judge(n.Game.Board) == EMPTY
}

func chooseMaxNode(root *MctsNode, mctsC float64) *MctsNode {
	if len(root.Children) == 0 {
		return root
	}

	maxNode := root
	maxScore := 0.0
	var score float64
	var childMax *MctsNode

	for _, child := range root.Children {
		childMax = chooseMaxNode(child, mctsC)

		// Playout every initial nodes first
		if !childMax.Played {
			return childMax
		}

		score = childMax.mctsFactor(root.Trials, mctsC)

		if score > maxScore {
			maxNode = childMax
			maxScore = score
		}
	}

	return maxNode
}

func (node *MctsNode) playout() (winner Color) {
	node.Played = true
	game := CopyGame(node.Game)

	//fmt.Println("playout on")
	//fmt.Println(hDump(*game.Board, 0))

	for {
		//fmt.Println("current board")
		//fmt.Println(hDump(*game.Board, 0))
		winner = Judge(game.Board)

		if winner != EMPTY {
			// fmt.Printf("%s won!\n", winner)

			return winner
		}

		if game.FreesCount <= 0 {
			return EMPTY
		}

		tryRandomMove(game)
	}
}

func (node *MctsNode) expandChildren() {
	for i := range node.Game.Board.Frees {
		newBoard := CopyBoard(node.Game.Board)

		newGame := &Game{
			Winner: node.Game.Winner,
			Turn:   node.Game.Turn,
			Board:  newBoard,
		}

		newGame.moveFree(i)

		child := NewMctsNode(newGame)

		node.appendChild(child)
	}
}

func tryRandomMove(game *Game) {
	game.moveFree(rand.Intn(game.Board.FreesCount))
}

func (node *MctsNode) mctsFactor(n int, mctsC float64) float64 {
	if node.Trials == 0 {
		return 0.0
	}

	return (float64(node.Wins) / float64(node.Trials)) +
		(mctsC * math.Sqrt(math.Log(float64(n))/float64(node.Trials)))
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
