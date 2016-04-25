package main

import (
	"bufio"
	"os"
	"strconv"
    "fmt"
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
    var err error
    var x, y int
	for {
        fmt.Println("Please input x and y")
		x = nextInt()
		y = nextInt()

        err = g.move(x, y)

        if err != nil {
            fmt.Printf("%s\n", err)
        }

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
    fmt.Println("Starting new Game")
	game := NewGame(WHITE, 4)
    game.playWithAI()
}
