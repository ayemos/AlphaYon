package main

type Node struct {
	Game
	Children []Node
}

const MCTS_C = 0.3

type MCTSRecrod struct {
	Wins, Trials int
}
