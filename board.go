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
	Radius int

	Pins        [][][]Color
	PinsHeights [][]int

	History []Coord

	/*
	 * TODO: Free positions are begin managed by game.moveFree
	 * and it's not the best way.
	 * Note that free positions would only be used by AI and its MCTS.
	 * (using more naive way is enough for displaying free posiions or
	 * other use cases.)
	 */
	Frees      []Coord
	FreesCount int

	Turns int
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

// Undo-ing and managing Free spaces isn't working together
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

func (b *Board) loadArray(arr []rune) (err error) {
	fmt.Printf("arr:len => %d:%d\n", len(arr), b.Radius * b.Radius)
	if len(arr) != b.Radius * b.Radius * b.Radius {
		return fmt.Errorf("Loaded array must be radius*radius in length.\n")
	}

	for i, rn := range arr {
		fmt.Println("hoge")
		x := i % b.Radius
		y := (i / b.Radius) % b.Radius
		z := (i / (b.Radius * b.Radius)) % b.Radius
		fmt.Printf("x:y:z => %d:%d:%d\n", x, y, z)

		switch rn {
		case 'b', 'B':
			b.Pins[x][y][z] = BLACK
		case 'w', 'W':
			b.Pins[x][y][z] = WHITE
		case '.', 'e', 'E':
			b.Pins[x][y][z] = EMPTY
		default:
			return fmt.Errorf("Invalid Color: %s.\n", rn)
		}
	}

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

	frees := make([]Coord, radius*radius)
	for x, _ := range frees {
		frees[x] = Coord{x % radius, x / radius}
	}

	board := &Board{
		Radius:      radius,
		Pins:        pins,
		PinsHeights: pinsHeights,
		History:     make([]Coord, radius*radius*radius),
		Turns:       0,
		Frees:       frees,
		FreesCount:  radius * radius,
	}

	return board
}

func (src *Board) CopyBoard (dst *Board) {
	r := src.Radius
	pins := make([][][]Color, r)
	copy(pins, src.Pins)

	pinsHeights := make([][]int, r)
	copy(pinsHeights, src.PinsHeights)

	history := make([]Coord, r*r*r)
	copy(history, src.History)

	frees := make([]Coord, r*r)
	copy(frees, src.Frees)

	dst = &Board {
		Radius:      r,
		Pins:        pins,
		PinsHeights: pinsHeights,
		History:     history,
		Turns:       src.Turns,
		Frees:       frees,
		FreesCount:  src.FreesCount,
	}
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

func (b Board) String() string {
	str := make([]byte, 0, 2*b.Radius*b.Radius*b.Radius+b.Radius+b.Radius*b.Radius)

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
