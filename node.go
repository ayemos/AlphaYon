package main

type MCTSRecord struct {
	Wins, Trials int
}

type Node struct {
	Game
	Children []Node
	MCTSRecord
	Played bool
}

const MCTS_C = 0.3

// String

func (n Node) String() string {
	/*
	str := make([]byte)
	n.Board
	n.Children
	*/
	return ""
}
