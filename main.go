package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	alphaYon "github.com/ayemos/AlphaYon/lib"
)

type command struct {
	flags *flag.FlagSet
	f     func(args []string) error
}

var sc = bufio.NewScanner(os.Stdin)

func play() {
	sc.Split(bufio.ScanWords)
	var err error
	game := alphaYon.NewGame(alphaYon.WHITE, 4)

	for {
		x := nextInt()
		y := nextInt()

		err = game.Move(x, y)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		game.Pretty()
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
	commands := map[string]command{
		"one_player":  onePlayerCmd(),
		"two_players": twoPlayersCmd(),
		"no_players":  noPlayersCmd(),
	}

	fs := flag.NewFlagSet("alphaYon", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println("Usage: alphaYon <command> [command flags]")

		fmt.Printf("\nglobal flags:\n")
		fs.PrintDefaults()
		for name, cmd := range commands {
			fmt.Printf("\n%s command:\n", name)
			cmd.flags.PrintDefaults()
		}
		fmt.Println(examples)
	}

	fs.Parse(os.Args[1:])

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	if cmd, ok := commands[args[0]]; !ok {
		log.Fatalf("Unknown command: %s", args[0])
	} else if err := cmd.f(args[1:]); err != nil {
		log.Fatal(err)
	}
}

const examples = `
examples:
  alphaYon one_player -mctsC=0.3 -timeLimit=5
  alphaYon no_players -mctsCA=0.3 -mctsCB=0.7 -timeLimitA=5 -timeLimitB=3
`
