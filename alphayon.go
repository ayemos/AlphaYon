package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func (g *Game) play() {
	sc.Split(bufio.ScanWords)
	for {
		x := nextInt()
		y := nextInt()

		g.move(x, y)
		g.pretty()
	}
}

func (g *Game) playWithAI() Game {
	sc.Split(bufio.ScanWords)
	for {
		x := nextInt()
		y := nextInt()

		g.move(x, y)

		g.pretty()
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())

	if e != nil {
		panic(e)
	}

	return i
}

func main() {
	game := NewGame(WHITE, 4)
}
