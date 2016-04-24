package main

func playout() {
}


/*
    def playout(node)
      turn = node.turn
      c = 0
      while (winner = judge(node.board)).nil? &&
        !(available_positions = node.board.available_positions).empty?

        pos = available_positions.sample(1).first
        node.board.push(turn, pos[0], pos[1])
        turn = next_turn(turn)
        c += 1
      end

      c.times do
        node.board.undo
      end

      winner
    end
	*/
