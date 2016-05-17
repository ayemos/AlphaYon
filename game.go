package main

type Game struct {
	Winner Color
	Turn   Color
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
		Winner: EMPTY,
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
