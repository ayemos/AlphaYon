package main

import (
	"flag"
	"fmt"

	alphaYon "github.com/ayemos/AlphaYon/lib"
)

func noPlayersCmd() command {
	fs := flag.NewFlagSet("alphaYon no_players", flag.ExitOnError)
	opts := &noPlayersOpts{}

	// TODO: Use fs.DurationVar(...)
	fs.Float64Var(&opts.mctsCA, "mctsCA", alphaYon.DefaultMCTSC, "C factor of MCTS for AI A.")
	fs.Float64Var(&opts.mctsCB, "mctsCB", alphaYon.DefaultMCTSC, "C factor of MCTS for AI B.")
	fs.IntVar(&opts.timeLimitA, "timeLimitA", alphaYon.DefaultTimeLimit, "Time limit for AI A")
	fs.IntVar(&opts.timeLimitB, "timeLimitB", alphaYon.DefaultTimeLimit, "Time limit for AI B")
	fs.BoolVar(&opts.resultOnly, "resultOnly", false, "Print result only.")

	return command{fs, func(args []string) error {
		fs.Parse(args)
		return noPlayers(opts)
	}}
}

func noPlayers(opts *noPlayersOpts) error {
	game := alphaYon.NewGame(alphaYon.WHITE, 4)

	aiA := alphaYon.NewAI(game, opts.mctsCA)
	aiB := alphaYon.NewAI(game, opts.mctsCB)

	var err error
	var status alphaYon.GameStatus
	var aiX, aiY int

	for {
		if !opts.resultOnly {
			fmt.Println("AI A is thinking...")
		}
		aiX, aiY = aiA.Solve(game.Turn, opts.timeLimitA)

		err = game.Move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}

		if !opts.resultOnly {
			game.Pretty()
		}

		status = alphaYon.Judge(game.Board)

		if status != alphaYon.RUNNING {
			fmt.Println(status)

			return nil
		}

		if !opts.resultOnly {
			fmt.Println("AI B is thinking...")
		}

		aiX, aiY = aiB.Solve(game.Turn, opts.timeLimitB)

		err = game.Move(aiX, aiY)

		if err != nil {
			fmt.Printf("%s\n", err)
			return err
		}

		if !opts.resultOnly {
			game.Pretty()
		}

		status = alphaYon.Judge(game.Board)

		if status != alphaYon.RUNNING {
			fmt.Println(status)

			return nil
		}
	}

	return nil
}

type noPlayersOpts struct {
	mctsCA     float64
	mctsCB     float64
	timeLimitA int
	timeLimitB int
	resultOnly bool
}
