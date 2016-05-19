package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func stat() {
	m := 5
	n := 10
	var winner Color
	wins := make([]int, n)
	draws := make([]int, n)

	for i := m; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			winner = aiMatch(i, j)

			if winner == WHITE {
				fmt.Printf("%d defeated %d\n", i, j)
				wins[i-1]++
			} else if winner == BLACK {
				fmt.Printf("%d defeated %d\n", j, i)
				wins[j-1]++
			} else {
				fmt.Printf("%d draw %d\n", i, j)
				draws[i-1]++
				draws[j-1]++
			}

			winner = aiMatch(j, i)

			if winner == BLACK {
				fmt.Printf("%d defeated %d\n", i, j)
				wins[i-1]++
			} else if winner == WHITE {
				fmt.Printf("%d defeated %d\n", j, i)
				wins[j-1]++
			} else {
				fmt.Printf("%d draw %d\n", i, j)
				draws[i-1]++
				draws[j-1]++
			}
		}
	}

	fmt.Println(wins)
	fmt.Println(draws)
}

func play() {
	sc.Split(bufio.ScanWords)
	var err error
	game := NewGame(WHITE, 4)

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

func playWithAI(timeLimit int) {
	game := NewGame(WHITE, 4)
	ai := NewAI(game, 0.3) // MCTS_C
	sc.Split(bufio.ScanWords)

	var err error
	var winner Color

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

		winner = Judge(game.Board)

		if winner != EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return
		}

		fmt.Println("AI is thinking...")

		aiX, aiY := ai.solve(game.Turn, timeLimit)

		// AI turn
		err = game.move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.pretty()

		winner = Judge(game.Board)

		if winner != EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return
		}
	}
}

func aiMatch(timeLimitA int, timeLimitB int) Color {
	game := NewGame(WHITE, 4)

	aiA := NewAI(game, 0.3) // MCTS_C
	aiB := NewAI(game, 0.3) // MCTS_C

	var err error
	var winner Color
	var aiX, aiY int

	for {
		fmt.Println("AI A is thinking...")
		aiX, aiY = aiA.solve(game.Turn, timeLimitA)

		err = game.move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.pretty()

		winner = Judge(game.Board)

		if winner != EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return winner
		}

		fmt.Println("AI B is thinking...")
		aiX, aiY = aiB.solve(game.Turn, timeLimitB)

		err = game.move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.pretty()

		winner = Judge(game.Board)

		if winner != EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return winner
		}
	}

	return EMPTY
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
	var withAI = flag.Bool("with-ai", false, "fight with ai.")
	var onlyAI = flag.Bool("only-ai", false, "ai v.s. ai.")
	var statAI = flag.Bool("stat-ai", false, "run statistic on ai.")
	flag.Parse()

	fmt.Println(*withAI)
	fmt.Println(*onlyAI)

	fmt.Println("Starting new Game")

	if *withAI {
		fmt.Println("'Play with AI' mode")
		playWithAI(3)
	} else if *onlyAI {
		fmt.Println("'AI V.S. AI' mode")
		aiMatch(1, 3)
	} else if *statAI {
		stat()
	}
}
