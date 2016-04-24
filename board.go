package main

import (
	"fmt"
)

type Color int

type Coord struct {
	X, Y int
}

const (
	EMPTY Color = iota
	WHITE
	BLACK
)

type Board struct {
	Radius      int
	Pins        [][][]Color
	PinsHeights [][]int
	History     []Coord
	Turns       int
}

func hoge() {
	fmt.Println("hoge!")
}

func (b *Board) push(x, y int, c Color) (err error) {
	if b.PinsHeights[x][y] == b.Radius {
		return fmt.Errorf("Invalid Move:(%d, %d)\n", x, y)
	}

	b.Pins[x][y][b.PinsHeights[x][y]] = c
	b.PinsHeights[x][y]++
	b.History[b.Turns] = Coord{x, y}
	b.Turns++

	return nil
}

func (b *Board) undo() (err error) {
	if b.Turns == 0 {
		return fmt.Errorf("History is Empty\n")
	}

	b.Turns--

	coord := b.History[b.Turns]
	x := coord.X
	y := coord.Y

	b.PinsHeights[x][y]--

	b.Pins[x][y][b.PinsHeights[x][y]] = EMPTY
	b.History[b.Turns] = Coord{0, 0}

	return nil
}

// コンストラクタ関数を定義
func NewBoard(radius int) *Board {
	pins := make([][][]Color, radius)

	for x, _ := range pins {
		pins[x] = make([][]Color, radius)

		for y, _ := range pins[x] {
			pins[x][y] = make([]Color, radius)
		}
	}

	pinsHeights := make([][]int, radius)
	for x, _ := range pinsHeights {
		pinsHeights[x] = make([]int, radius)
	}

	board := &Board{
		Radius:      radius,
		Pins:        pins,
		PinsHeights: pinsHeights,
		History:     make([]Coord, radius*radius*radius),
		Turns:       0,
	}

	return board
}

func (b Board) pretty() {
	fmt.Printf("+++++Pretty Board+++++\n")
	fmt.Printf("%s\n", b)
	fmt.Printf("++++++++++++++++++++++\n")
}

func (c Color) color2byte() byte {
	switch c {
	case WHITE:
		return 'w'
	case BLACK:
		return 'b'
	case EMPTY:
		return '.'
	}
	return '?'
}

// String
func (c Color) String() string {
	switch c {
	case WHITE:
		return "w"
	case BLACK:
		return "b"
	case EMPTY:
		return "."
	}
	return "?"
}

func (b Board) String() string {
	str := make([]byte, 1000, b.Radius*b.Radius*b.Radius+1000)

	for z := b.Radius - 1; z >= 0; z-- {
		str = append(str, '\n')

		for y := b.Radius - 1; y >= 0; y-- {
			str = append(str, '\n')

			for x := 0; x < b.Radius; x++ {
				str = append(str, b.Pins[x][y][z].color2byte())

				if x != b.Radius {
					str = append(str, ' ')
				}
			}
		}
	}

	return string(str)
}
