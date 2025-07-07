package tictactoe

import (
	"testing"
)

func boardgen(positions map[[2]int]rune, empty_board *[3][3]rune) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var pos_pair = [2]int{i, j}
			var value, ok = positions[pos_pair]
			if ok {
				empty_board[pos_pair[0]][pos_pair[1]] = value
			} else {
				empty_board[i][j] = '-'
			}
		}
	}
}

func TestT(t *testing.T) {
	var board [3][3]rune
	boardgen(
		map[[2]int]rune{
			{0, 0}: 'X',
			{1, 1}: 'O',
			{2, 0}: 'X',
		},
		&board,
	)
	render(&board);
	if !won(&board, 1, 0, 'X') {
		t.Error("X should have won")
	}
}
