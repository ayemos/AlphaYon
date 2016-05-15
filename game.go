package main

import (
	"fmt"
)

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

func (g *Game) moveFree(f int) (err error) {
	x := g.Frees[f].X
	y := g.Frees[f].Y

	if g.PinsHeights[x][y] == g.Radius {
		// position (x, y) was filled
		g.FreesCount--
		g.Frees[f] = g.Frees[g.FreesCount]
	}

	err = g.move(x, y)

	return err
}

func Judge(b *Board) (winner Color) {
	winner = EMPTY
	for x := 0; x < b.Radius; x++ {
		for y := 0; y < b.Radius; y++ {
			winner = judgeFromPoint(b, x, y, 0)

			if winner != EMPTY {
				return winner
			}
		}
	}

	for y := 0; y < b.Radius; y++ {
		for z := 1; z < b.Radius; z++ {
			winner = judgeFromPoint(b, 0, y, z)

			if winner != EMPTY {
				return winner
			}
		}
	}

	return EMPTY
}

func judgeFromPoint(b *Board, x, y, z int) (winner Color) {
	fmt.Printf("X:Y:Z => %d:%d:%d\n", x, y, z)

	var color Color
	var count int

	for dx := 0; dx <= 1; dx++ {
		for dy := 0; dy <= 1; dy++ {
			for dz := 0; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				} else if (x + dx > 0) ||
				(y + dy > 0) ||
				(z + dx > 0) {
					// x != 0 && dx != 0
					// (no way to find line)
					continue
				}

				// pick first color
				color = b.Pins[x][y][z]

				if color == EMPTY {
					return EMPTY
				}

				x += dx
				y += dy
				z += dz
				
				count = 1

				for ; x >= 0 && x < b.Radius &&
				y >= 0 && y < b.Radius &&
				z >= 0 && z < b.Radius; {
					if color != b.Pins[x][y][z] {
						break
					}

					x += dx
					y += dy
					z += dz

					count += 1

					if count == b.Radius {
						return color
					}
				}

			}
		}
	}

	return EMPTY
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
