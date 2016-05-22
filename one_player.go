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
	fs.IntVar(&opts.mctsT, "mctsT", alphaYon.DefaultMCTST, "T factor of MCTS for AI.")
	fs.IntVar(&opts.timeLimit, "searchDepth", alphaYon.DefaultSearchDepth, "Search depth for AI")
	fs.IntVar(&opts.timeLimit, "timeLimit", alphaYon.DefaultTimeLimit, "Time limit for AI")

	return command{fs, func(args []string) error {
		fs.Parse(args)
		return onePlayer(opts)
	}}
}

func onePlayer(opts *onePlayerOpts) error {
	game := alphaYon.NewGame(alphaYon.WHITE, 4)
	ai := alphaYon.NewAI(game, opts.mctsC, opts.mctsT)
	sc.Split(bufio.ScanWords)

	var err error
	var status alphaYon.GameStatus

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

		status = alphaYon.Judge(game.Board)

		if status != alphaYon.RUNNING {
			fmt.Println(status)
			return nil
		}

		fmt.Println("AI is thinking...")

		aiX, aiY := ai.Solve(game.Turn, opts.timeLimit, opts.searchDepth)

		// AI turn
		err = game.Move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.Pretty()

		status = alphaYon.Judge(game.Board)

		if status != alphaYon.RUNNING {
			fmt.Println(status)
			return nil
		}
	}
}

type onePlayerOpts struct {
	mctsC       float64
	mctsT       int
	timeLimit   int
	searchDepth int
}
