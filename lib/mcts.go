package alphaYon

import (
	//	"fmt"
	"math"
	"math/rand"
	"time"
)

func Mcts(root *MctsNode, player Color, mctsC float64, mctsT int, timeLimitSec int, searchDepth int) {
	root.expandChildren(player, searchDepth)

	/*
		fmt.Println("mcts")
		fmt.Println(root.Game.Turn)
		fmt.Println(player)
	*/

	var maxNode, node *MctsNode
	var status GameStatus
	var win, draw int
	var diffSec float64

	startTime := time.Now()

	//for i := 0; i < 1000; i++ {
	for i := 0; ; i++ {
		if i%100 == 0 {
			diffSec = time.Now().Sub(startTime).Seconds()

			if diffSec >= float64(timeLimitSec) {
				break
			}
		}

		maxNode = chooseMaxNode(root, mctsC)

		// TODO: なんとかする
		// playout
		status = maxNode.playout()

		if (root.Game.Turn == BLACK && status == BLACK_WON) ||
			(root.Game.Turn == WHITE && status == WHITE_WON) {
			win = 1
			draw = 0
		} else if status == DRAW {
			win = 0
			draw = 1
		} else {
			// error ?
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

		if shouldExpand(maxNode, player, mctsT) {
			maxNode.expandChildren(player, searchDepth)
		}
	}
}

func chooseMaxNode(root *MctsNode, mctsC float64) *MctsNode {
	if len(root.Children) == 0 {
		return root
	}

	maxNode := root
	maxScore := 0.0
	var score float64
	var child, childMax *MctsNode
	var j int
	l := len(root.Children)
	c := rand.Intn(l)

	for i := 0; i < l; i++ {
		j = (i + c) % l
		child = root.Children[j]

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

func shouldExpand(n *MctsNode, player Color, mctsT int) bool {
	return n.Trials >= mctsT &&
		len(n.Children) == 0 &&
		n.Game.FreesCount > 0 &&
		Judge(n.Game.Board) == RUNNING
}

func (node *MctsNode) playout() (status GameStatus) {
	node.Played = true
	game := CopyGame(node.Game)

	//fmt.Println("playout on")
	//fmt.Println(hDump(*game.Board, 0))

	for {
		//fmt.Println("current board")
		//fmt.Println(hDump(*game.Board, 0))
		status = Judge(game.Board)

		if status != RUNNING {
			// fmt.Printf("%s won!\n", winner)

			return status
		}

		if game.FreesCount <= 0 {
			return RUNNING
		}

		tryRandomMove(game)
	}
}

func (node *MctsNode) expandChildren(player Color, searchDepth int) error {
	l := node.Game.Board.FreesCount
	c := rand.Intn(l)
	var j int

	for i := 0; i < l; i++ {
		j = (i + c) % l // randomize i

		newBoard := CopyBoard(node.Game.Board)

		newGame := &Game{
			Status: node.Game.Status,
			Turn:   node.Game.Turn,
			Board:  newBoard,
		}

		newGame.MoveFree(j)

		child := NewMctsNode(newGame)

		//		fmt.Println(child)
		//		now := time.Now()
		searchResult, err := Search(child.convToABNode(), player, searchDepth)
		//		fmt.Println("TIME", ":", time.Now().Sub(now))

		if err != nil {
			return err
		}

		/*
			if *searchResult == LOSING {
				fmt.Println("LOSING!")
				fmt.Println(node)
			}

			if *searchResult == WINNING {
				fmt.Println("WINNING!")
				fmt.Println(node)
			}
		*/

		if *searchResult != LOSING {
			node.appendChild(child)
		}
	}

	return nil
}

func tryRandomMove(game *Game) {
	game.MoveFree(rand.Intn(game.Board.FreesCount))
}

func (node *MctsNode) mctsFactor(n int, mctsC float64) float64 {
	if node.Trials == 0 {
		return 0.0
	}

	return (float64(node.Wins) / float64(node.Trials)) +
		(mctsC * math.Sqrt(math.Log(float64(n))/float64(node.Trials)))
}
