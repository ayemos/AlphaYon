package main

type MCTSRecord struct {
	Wins, Trials int
}

type Tree struct {
    Root *Node
}

type Node struct {
	Game
	Children []Node
	Parent   *Node
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
