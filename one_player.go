package main

import (
	"bufio"
	"flag"
	"fmt"

	alphaYon "github.com/ayemos/AlphaYon/lib"
)

func onePlayerCmd() command {
	fs := flag.NewFlagSet("alphaYon one_players", flag.ExitOnError)
	opts := &onePlayerOpts{}

	fs.Float64Var(&opts.mctsC, "mctsC", alphaYon.DefaultMCTSC, "C factor of MCTS for AI.")
	fs.IntVar(&opts.timeLimit, "timeLimit", alphaYon.DefaultTimeLimit, "Time limit for AI")

	return command{fs, func(args []string) error {
		fs.Parse(args)
		return onePlayer(opts)
	}}
}

func onePlayer(opts *onePlayerOpts) error {
	game := alphaYon.NewGame(alphaYon.WHITE, 4)
	ai := alphaYon.NewAI(game, opts.mctsC) // MCTS_C
	sc.Split(bufio.ScanWords)

	var err error
	var winner alphaYon.Color

	for {
		for {
			fmt.Println("Please input x and y")
			x := nextInt()
			y := nextInt()

			err = game.Move(x, y)

			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				break
			}
		}

		game.Pretty()

		winner = alphaYon.Judge(game.Board)

		if winner != alphaYon.EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return nil
		}

		fmt.Println("AI is thinking...")

		aiX, aiY := ai.Solve(game.Turn, opts.timeLimit)

		// AI turn
		err = game.Move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.Pretty()

		winner = alphaYon.Judge(game.Board)

		if winner != alphaYon.EMPTY {
			fmt.Printf("%s Won!\n", winner)
			return nil
		}
	}
}

type onePlayerOpts struct {
	mctsC     float64
	timeLimit int
}
