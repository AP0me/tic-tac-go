package main

import (
	"testing"
)

// boardgen populates a board from a map of positions.
func boardgen(positions map[[2]int]rune, empty_board *[3][3]rune) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var pos_pair = [2]int{i, j}
			var value, ok = positions[pos_pair]
			if ok {
				empty_board[pos_pair[0]][pos_pair[1]] = value
			} else {
				empty_board[i][j] = ' ' // Use space for empty cells for clarity
			}
		}
	}
}

func TestTicTacToeBoardStates(t *testing.T) {
	var players = [2]rune{'X', 'O'}

	// === Test Cases from Image Analysis ===
	// Red is 'X', Blue/Cyan is 'O'. Black 'x' is the last move.

	// Image Board 1
	// Board:
	// X O -
	// X - -
	// X - -  (Last move at (2,0)) -> X wins
	t.Run("ImageBoard1_X_Win_Col0", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O',
				{1, 0}: 'X',
				{2, 0}: 'X', // Last move
			},
			&board.b,
		)
		if !won(&board, 2, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Image Board 1")
		}
	})

	// Image Board 2
	// Board:
	// X X X (Last move at (0,1)) -> X wins
	// O - -
	// O - -
	t.Run("ImageBoard2_X_Win_Row0", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'X', // Last move at (0,1)
				{1, 0}: 'O',
				{2, 0}: 'O',
			},
			&board.b,
		)
		if !won(&board, 0, 1, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Image Board 2")
		}
	})

	// Image Board 3
	// Board:
	// X - -
	// - O - (Last move at (1,1)) -> No win
	// - X X
	t.Run("ImageBoard3_O_at_1_1_NoWin", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X',
				{1, 1}: 'O', // Last move
				{2, 1}: 'X', {2, 2}: 'X',
			},
			&board.b,
		)
		if won(&board, 1, 1, 'O', players[:]) {
			render(&board)
			t.Error("O should NOT have won in Image Board 3")
		}
	})

	// Image Board 4
	// Board:
	// - - X
	// - - X
	// X - - (Last move at (2,0)) -> No win
	t.Run("ImageBoard4_X_at_2_0_NoWin", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 2}: 'X',
				{1, 2}: 'X',
				{2, 0}: 'X', // Last move
			},
			&board.b,
		)
		if won(&board, 2, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Image Board 4")
		}
	})

	// Image Board 5
	// Board:
	// O O -
	// X O X
	// X X - (Last move at (2,0)) -> No win
	t.Run("ImageBoard5_X_at_2_0_NoWin", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'O', {0, 1}: 'O',
				{1, 0}: 'X', {1, 1}: 'O', {1, 2}: 'X',
				{2, 0}: 'X', {2, 1}: 'X', // Last move at (2,0)
			},
			&board.b,
		)
		if won(&board, 2, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Image Board 5")
		}
	})

	// Image Board 6, 7, 8 (All are the same board state)
	// Board:
	// X O X
	// O O X
	// X O X
	// No last move marker. This is a win for X in column 2.
	// We'll assume the last move was X at (2,2) to complete the column.
	t.Run("ImageBoard6_7_8_X_Win_Col2", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'O', {1, 2}: 'X',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		// Check for win with the last move that would cause it.
		if !won(&board, 2, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Image Board 6/7/8")
		}
	})

	// Image Board 9
	// Board:
	// X - - (Last move at (0,0))
	// - - -
	// - - -
	// The marker is on an empty cell. Assuming it's the first move.
	t.Run("ImageBoard9_X_at_0_0_NoWin", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', // Last move
			},
			&board.b,
		)
		if won(&board, 0, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Image Board 9")
		}
	})

	// Image Board 10
	// Board:
	// X X X -> X has already won on row 0
	// O O O -> O has also won on row 1
	// X O X
	// Last move marked on O at (1,1). This is an invalid board state.
	// The test will check if O's last move resulted in a win for O.
	t.Run("ImageBoard10_O_Win_Row1", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'O', {1, 2}: 'O', // Last move at (1,1)
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		if !won(&board, 1, 1, 'O', players[:]) {
			render(&board)
			t.Error("O should have won in Image Board 10")
		}
	})

	// Image Board 11
	// Board:
	// X O X
	// O O O (Last move at (1,2)) -> O wins
	// X O X
	t.Run("ImageBoard11_O_Win_Row1", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'O', {1, 2}: 'O', // Last move
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		if !won(&board, 1, 2, 'O', players[:]) {
			render(&board)
			t.Error("O should have won in Image Board 11")
		}
	})

	// Image Board 12
	// Board:
	// X X X -> X has already won on row 0
	// X O X
	// X O X (Last move at (2,2))
	// Test checks if X's last move resulted in a win.
	t.Run("ImageBoard12_X_Win_Row0", func(t *testing.T) {
		var board Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'X',
				{1, 0}: 'X', {1, 1}: 'O', {1, 2}: 'X',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X', // Last move
			},
			&board.b,
		)
		// The win is on row 0, but the last move was (2,2).
		// A robust `won` function would detect the win regardless of the last move.
		// Our simple one requires the last move to be part of the winning line.
		// So we test the winning line directly.
		if !won(&board, 0, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Image Board 12 on row 0")
		}
	})
}
