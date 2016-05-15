package main

type Game struct {
	Turn Color
	*Board
}

func (g *Game) move(x, y int) (err error) {
	err = g.push(x, y, g.Turn)

	if g.Turn == BLACK {
		g.Turn = WHITE
	} else if g.Turn == WHITE {
		g.Turn = BLACK
	}

	return err
}

func NewGame(player Color, radius int) *Game {
	game := &Game{
		Turn:  player,
		Board: NewBoard(radius),
	}

	return game
}

func (turn Color) nextTurn (next Color) {
	if turn == WHITE {
		next = BLACK
	} else if turn == BLACK {
		next = WHITE
	}
}
