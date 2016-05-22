package alphaYon

import (
	"fmt"
)

type GameStatus int

type Game struct {
	Status GameStatus
	Turn   Color
	*Board
}

const (
	WHITE_WON GameStatus = iota
	BLACK_WON
	RUNNING
	DRAW
)

func (g *Game) Move(x, y int) (err error) {
	err = g.push(x, y, g.Turn)

	if g.Turn == BLACK {
		g.Turn = WHITE
	} else if g.Turn == WHITE {
		g.Turn = BLACK
	}

	if g.PinsHeights[x][y] == g.Radius {
		g.updateFrees(x, y)
	}

	return err
}

func (g *Game) updateFrees(x, y int) error {
	var target int

	for i, f := range g.Frees {
		if x == f.X && y == f.Y {
			target = i
			break
		}
	}

	g.FreesCount--

	if g.FreesCount >= 0 {
		g.Frees[target] = g.Frees[g.FreesCount]
	}

	return nil
}

func (g *Game) MoveFree(f int) (err error) {
	x := g.Frees[f].X
	y := g.Frees[f].Y

	if g.PinsHeights[x][y] == g.Radius {
		// position (x, y) was filled
		if f < 0 || g.FreesCount < 0 {
			fmt.Println(g.Board)
			fmt.Println(f)
			fmt.Println(g.FreesCount)
		}

		g.FreesCount--
		g.Frees[f] = g.Frees[g.FreesCount]
	}

	err = g.push(x, y, g.Turn)

	if g.Turn == BLACK {
		g.Turn = WHITE
	} else if g.Turn == WHITE {
		g.Turn = BLACK
	}

	return err
}

func Judge(b *Board) (status GameStatus) {
	status = RUNNING

	for x := 0; x < b.Radius; x++ {
		for y := 0; y < b.Radius; y++ {
			status = judgeFromPoint(b, x, y, 0)

			if status != RUNNING {
				return status
			}
		}
	}

	for x := 1; x < b.Radius; x++ {
		for z := 1; z < b.Radius; z++ {
			status = judgeFromPoint(b, x, 0, z)

			if status != RUNNING {
				return status
			}
		}
	}

	for y := 0; y < b.Radius; y++ {
		for z := 1; z < b.Radius; z++ {
			status = judgeFromPoint(b, 0, y, z)

			if status != RUNNING {
				return status
			}
		}
	}

	if b.FreesCount == 0 {
		return DRAW
	} else {
		return RUNNING
	}
}

func judgeFromPoint(b *Board, x, y, z int) (winner GameStatus) {
	var color Color
	var count int
	var tmpx, tmpy, tmpz int

	for dx := 0; dx <= 1; dx++ {
		for dy := 0; dy <= 1; dy++ {
			for dz := 0; dz <= 1; dz++ {
				tmpx = x
				tmpy = y
				tmpz = z

				if dx == 0 && dy == 0 && dz == 0 {
					continue
				} else if (tmpx*dx > 0) ||
					(tmpy*dy > 0) ||
					(tmpz*dz > 0) {
					// x != 0 && dx != 0
					// (no way to find line)
					continue
				}

				// pick first color
				color = b.Pins[x][y][z]

				if color == EMPTY {
					continue
				}

				tmpx += dx
				tmpy += dy
				tmpz += dz

				count = 1

				for tmpx >= 0 && tmpx < b.Radius &&
					tmpy >= 0 && tmpy < b.Radius &&
					tmpz >= 0 && tmpz < b.Radius {
					if color != b.Pins[tmpx][tmpy][tmpz] {
						break
					}

					tmpx += dx
					tmpy += dy
					tmpz += dz

					count += 1

					if count == b.Radius {
						if color == WHITE {
							return WHITE_WON
						} else {
							return BLACK_WON
						}
					}
				}

			}
		}
	}

	return RUNNING
}

func CopyGame(src *Game) *Game {
	return &Game{
		Status: src.Status,
		Turn:   src.Turn,
		Board:  CopyBoard(src.Board),
	}
}

func NewGame(player Color, radius int) *Game {
	game := &Game{
		Status: RUNNING,
		Turn:   player,
		Board:  NewBoard(radius),
	}

	return game
}

func (turn Color) nextTurn() Color {
	if turn == WHITE {
		return BLACK
	} else if turn == BLACK {
		return WHITE
	}
	return EMPTY
}
