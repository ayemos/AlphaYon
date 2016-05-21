package main

import (
	"errors"
	"flag"
)

func twoPlayersCmd() command {
	fs := flag.NewFlagSet("alphaYon two_players", flag.ExitOnError)
	opts := &twoPlayersOpts{}

	return command{fs, func(args []string) error {
		fs.Parse(args)
		return twoPlayers(opts)
	}}
}

func twoPlayers(opts *twoPlayersOpts) error {
	return errors.New("Not implemented yet.")
}

type twoPlayersOpts struct {
}
