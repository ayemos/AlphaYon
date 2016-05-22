package alphaYon

import (
//	"fmt"
)

func Search(root *ABNode, player Color, maxDepth int) (status *ABNodeStatus, err error) {
	res, err := searchSub(root, player, 0, maxDepth)
	return res, err
}

func searchSub(node *ABNode, player Color, depth int, maxDepth int) (*ABNodeStatus, error) {
	var res ABNodeStatus
	var err error

	status := Judge(node.Game.Board)

	if depth >= maxDepth || status != RUNNING {
		// TODO: なんとかする
		if (status == BLACK_WON && player == BLACK) ||
			(status == WHITE_WON && player == WHITE) {
			res = WINNING
		} else if (status == BLACK_WON && player == WHITE) ||
			(status == WHITE_WON && player == BLACK) {
			res = LOSING
		} else {
			res = UNKNOWN
		}

		return &res, nil
	} else {
		for i := 0; i < node.Game.Board.FreesCount; i++ {
			newBoard := CopyBoard(node.Game.Board)

			newGame := &Game{
				Status: node.Game.Status,
				Turn:   node.Game.Turn,
				Board:  newBoard,
			}

			err = newGame.MoveFree(i)
			if err != nil {
				return nil, err
			}

			child := NewABNode(newGame)

			node.appendChild(child)

			tmp, err := searchSub(child, player, depth+1, maxDepth)

			if err != nil {
				return nil, err
			}

			child.ABNodeStatus = *tmp
		}

		// summarize  children
		if player == node.Game.Turn {
			winAll := true

			for _, child := range node.Children {
				if child.ABNodeStatus != WINNING {
					winAll = false
					break
				}
			}

			if winAll {
				res = WINNING
				node.ABNodeStatus = res
				return &res, nil
			}

			// loseAny
			for _, child := range node.Children {
				if child.ABNodeStatus == LOSING {
					res = LOSING
					node.ABNodeStatus = res
					return &res, nil
				}
			}
		} else {
			loseAll := true

			for _, child := range node.Children {
				if child.ABNodeStatus != LOSING {
					loseAll = false
					break
				}
			}

			if loseAll {
				res = LOSING
				node.ABNodeStatus = res
				return &res, nil
			}

			// winAny
			for _, child := range node.Children {
				if child.ABNodeStatus == WINNING {
					res = WINNING
					node.ABNodeStatus = res
					return &res, nil
				}
			}
		}

		res = UNKNOWN
		node.ABNodeStatus = res
		return &res, nil
	}
}
