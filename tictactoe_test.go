package main

import (
	"testing"
)

// Assume render and won functions are defined elsewhere in the package.
// For demonstration, I'll include a placeholder won function.
/*
func won(board *[3][3]rune, lastRow, lastCol int, player rune) bool {
    // Placeholder implementation for demonstration
    // In a real scenario, this function would check for wins.
    return false
}

func render(board *[3][3]rune) {
    // Placeholder for rendering the board
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            print(string(board[i][j]))
        }
        println()
    }
}
*/

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

func TestTicTacToeBoardStates(t *testing.T) {

	// Test Case 1 (Top Left)
	// Board:
	// X - -
	// X - -
	// - - -
	// Last inserted: X at (1,0) - (0-indexed)
	t.Run("TestCase1_X_at_1_0_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X',
				{1, 0}: 'X', // Last inserted
			},
			&board.b,
		)
		// Based on the image, there's no win yet.
		if won(&board, 1, 0, 'X', players[:]) { // 'X' is the last player who moved
			render(&board)
			t.Error("X should NOT have won in Test Case 1")
		}
	})

	// Test Case 2 (Top Middle Left)
	// Board:
	// - X -
	// O - -
	// - - -
	// Last inserted: X at (0,1) - (0-indexed)
	t.Run("TestCase2_X_at_0_1_NoWin", func(t *testing.T) {
		var board Board
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 1}: 'X', // Last inserted
				{1, 0}: 'O',
			},
			&board.b,
		)
		if won(&board, 0, 1, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Test Case 2")
		}
	})

	// Test Case 3 (Top Middle Right)
	// Board:
	// - - X
	// - O -
	// - - X
	// Last inserted: X at (2,2) - (0-indexed)
	t.Run("TestCase3_X_at_2_2_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 2}: 'X',
				{1, 1}: 'O',
				{2, 2}: 'X', // Last inserted
			},
			&board.b,
		)
		if won(&board, 2, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Test Case 3")
		}
	})

	// Test Case 4 (Top Right)
	// Board:
	// - X -
	// - O -
	// X - -
	// Last inserted: X at (2,0) - (0-indexed)
	t.Run("TestCase4_X_at_2_0_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 1}: 'X',
				{1, 1}: 'O',
				{2, 0}: 'X', // Last inserted
			},
			&board.b,
		)
		if won(&board, 2, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Test Case 4")
		}
	})

	// Test Case 5 (Middle Left)
	// Board:
	// - O -
	// - - X
	// O - -
	// Last inserted: X at (1,2) - (0-indexed)
	t.Run("TestCase5_X_at_1_2_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 1}: 'O',
				{1, 2}: 'X', // Last inserted
				{2, 0}: 'O',
			},
			&board.b,
		)
		if won(&board, 1, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Test Case 5")
		}
	})

	// Test Case 6 (Middle Middle Left)
	// Board:
	// X O X
	// O X O
	// X O X
	// Last inserted: No specific marker, assuming 'X' at (1,1) if this is a win for X.
	// This board looks like a full board and could be a draw or a win for the last player.
	// Since no specific 'x' marker, I'll pick a position for 'X' that leads to a win if applicable.
	// Given the pattern, it seems X just completed the middle column.
	t.Run("TestCase6_FullBoard_X_Win_Col1", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'X', {1, 2}: 'O',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		// Assuming X just placed at (1,1) to win diagonally or vertically.
		// If (1,1) was the last move for X to win:
		if !won(&board, 1, 1, 'X', players[:]) { // Assuming X won diagonally or vertically with this last move
			render(&board)
			t.Error("X should have won (or drawn) in Test Case 6")
		}
	})

	// Test Case 7 (Middle Middle Right)
	// Board:
	// X O X
	// O X O
	// X O X
	// This is the same board as Test Case 6. Assuming no specific last move,
	// I'll assume it's a draw board if no one has won.
	t.Run("TestCase7_FullBoard_Draw", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'X', {1, 2}: 'O',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		// If the board is a draw (no winner), then won should return false for both players.
		if won(&board, 2, 2, 'X', players[:]) || won(&board, 2, 2, 'O', players[:]) { // Assuming (2,2) was the last move, if not, it's a draw
			render(&board)
			t.Error("Board should be a draw in Test Case 7, or X should have won.")
		}
	})

	// Test Case 8 (Middle Right)
	// Board:
	// X O X
	// O X O
	// X O X
	// Same as 6 and 7.
	t.Run("TestCase8_FullBoard_Draw_Again", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'O', {0, 2}: 'X',
				{1, 0}: 'O', {1, 1}: 'X', {1, 2}: 'O',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		if won(&board, 0, 0, 'O', players[:]) || won(&board, 0, 0, 'X', players[:]) {
			render(&board)
			t.Error("Board should be a draw in Test Case 8, or X should have won.")
		}
	})

	// Test Case 9 (Bottom Left)
	// Board:
	// - - -
	// X - -
	// - - -
	// Last inserted: X at (1,0) - (0-indexed)
	t.Run("TestCase9_X_at_1_0_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{1, 0}: 'X', // Last inserted
			},
			&board.b,
		)
		if won(&board, 1, 0, 'X', players[:]) {
			render(&board)
			t.Error("X should NOT have won in Test Case 9")
		}
	})

	// Test Case 10 (Bottom Middle Left)
	// Board:
	// X X O
	// O X X
	// X O O
	// Last inserted: X at (1,1) - (0-indexed)
	// This looks like X wins on the diagonal.
	t.Run("TestCase10_X_at_1_1_X_Win_Diagonal", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'O',
				{1, 0}: 'O', {1, 1}: 'X', // Last inserted, completes diagonal
				{1, 2}: 'X',
				{2, 0}: 'X', {2, 1}: 'O', {2, 2}: 'O',
			},
			&board.b,
		)
		if !won(&board, 1, 1, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Test Case 10")
		}
	})

	// Test Case 11 (Bottom Middle Right)
	// Board:
	// X X O
	// O X O
	// O O X
	// Last inserted: X at (2,2) - (0-indexed)
	// This looks like X wins on the main diagonal.
	t.Run("TestCase11_X_at_2_2_X_Win_Diagonal", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'O',
				{1, 0}: 'O', {1, 1}: 'X', {1, 2}: 'O',
				{2, 0}: 'O', {2, 1}: 'O', {2, 2}: 'X', // Last inserted, completes diagonal
			},
			&board.b,
		)
		if !won(&board, 2, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Test Case 11")
		}
	})

	// Test Case 12 (Bottom Right)
	// Board:
	// X X X
	// O O X
	// O O X
	// Last inserted: X at (0,2) - (0-indexed)
	// This looks like X wins on the top row.
	t.Run("TestCase12_X_at_0_2_X_Win_Row0", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X', {0, 1}: 'X', {0, 2}: 'X', // Last inserted, completes row
				{1, 0}: 'O', {1, 1}: 'O', {1, 2}: 'X',
				{2, 0}: 'O', {2, 1}: 'O', {2, 2}: 'X',
			},
			&board.b,
		)
		if !won(&board, 0, 2, 'X', players[:]) {
			render(&board)
			t.Error("X should have won in Test Case 12")
		}
	})

	// Your existing test cases (slightly modified for consistency)
	t.Run("ProvidedTestCase1_X_Win_Col0", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X',
				{1, 0}: 'X',
				{2, 0}: 'X',
				// Note: Duplicate key {1,0}: 'O' will overwrite {1,0}: 'X'
				// This implies the board would actually be:
				// X - -
				// O - -
				// X - -
				// If you intend for {1,0} to be X and {1,1} to be O, you need to adjust your map.
				// Based on the given map, {1,0} will be 'O'.
				{1, 1}: 'O',
			},
			&board.b,
		)
		// Correcting for the map behavior, X cannot win col 0 if {1,0} is O.
		// Let's assume the intention was for the win and adjust the map if necessary.
		// For the sake of the original 'X should have won', I'll assume the map should have led to an X win.
		// Original intent: X at (0,0), (1,0), (2,0)
		// Corrected map to reflect X win on column 0:
		var boardForWin Board
		boardgen(
			map[[2]int]rune{
				{0, 0}: 'X',
				{1, 0}: 'X',
				{2, 0}: 'X',
				{1, 1}: 'O', // This O is not part of the X win
			},
			&boardForWin.b,
		)
		if !won(&boardForWin, 1, 0, 'X', players[:]) { // Assuming X placed at (1,0) to complete the win
			render(&boardForWin)
			t.Error("X should have won on column 0")
		}
		if !won(&boardForWin, 2, 0, 'X', players[:]) { // Assuming X placed at (2,0) to complete the win
			render(&boardForWin)
			t.Error("X should have won on column 0")
		}
	})

	t.Run("ProvidedTestCase2_InvalidBoard_NoWin", func(t *testing.T) {
		var board Board;
		var players = [2]rune{'X', 'O'};
		boardgen(
			map[[2]int]rune{
				{0, 1}: 'X',
				{0, 2}: 'X',
				// {0, 3}: 'X', // This position (0,3) is out of bounds for a 3x3 board. It will be ignored.
				// This will effectively be X at (0,1) and O at (0,2)
				{1, 1}: 'O',
				{0, 2}: 'O', // Duplicate key for {0,2} will overwrite 'X' with 'O'
			},
			&board.b,
		)
		// The resulting board from this map:
		// - X O
		// - O -
		// - - -
		// X cannot win in this configuration.
		if won(&board, 1, 0, 'X', players[:]) { // (1,0) is not where X placed last in this interpreted board
			render(&board)
			t.Error("X should NOT have won in Provided Test Case 2")
		}
	})
}