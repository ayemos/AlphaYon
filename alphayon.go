package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func play(game *Game) {
	sc.Split(bufio.ScanWords)
	var err error

	for {
		x := nextInt()
		y := nextInt()

		err = game.move(x, y)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.pretty()
	}
}

func playWithAI(game *Game, numMcts int) Game {
	ai := NewAI(game, 0.3) // MCTS_C
	sc.Split(bufio.ScanWords)
	var err error

	for {
		for {
			fmt.Println("Please input x and y")
			x := nextInt()
			y := nextInt()

			err = game.move(x, y)

			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				break
			}
		}

		game.pretty()
		fmt.Printf("Free positions:\n");
		fmt.Printf("%s\n", game.Frees);
		fmt.Printf("%s\n", game.PinsHeights);

		fmt.Println("AI is thinking...");

		aiX, aiY := ai.solve()
		// AI turn
		err = game.move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err);
		}

		game.pretty()
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
	var withAI = flag.Bool("with-ai", true, "fight with ai.")
	var numMcts = flag.Int("num-mcts", 3000, "number of trial in mcts.")
	flag.Parse()

	fmt.Println("Starting new Game")
	game := NewGame(WHITE, 4)

	if *withAI {
		playWithAI(game, *numMcts)
	} else {
		play(game)
	}
}
