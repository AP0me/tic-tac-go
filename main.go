package main
import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func render(board *[3][3]rune){
	fmt.Println("  1 2 3 x")
    for i, board_line := range board {
        fmt.Printf("%d ", i+1);
        for _, board_char := range board_line {
            fmt.Printf("%c ", board_char)
        }
        fmt.Println()
    }
	fmt.Println("y")
}

func won(board *[3][3]rune, x int , y int, player rune) bool {

	var all_same_x = true;
	for i := 0; i < 3; i++ {
		var x_i = (x+i) % 3;
		if(board[x_i][y] != player){
			all_same_x = false;
			break;
		}
	}

	var all_same_y = true;
	for i := 0; i < 3; i++ {
		var y_i = (y+i) % 3;
		if(board[x][y_i] != player){
			all_same_y = false;
			break;
		}
	}

	var all_same_d14 = true;
	for i := 0; i < 3; i++ {
		var x_i = (x+i) % 3;
		var y_i = (y+i) % 3;
		if(board[x_i][y_i] != player){
			all_same_d14 = false;
			break;
		}
	}

	var all_same_d32 = true;
	for i := 0; i < 3; i++ {
		var x_i = (x+i) % 3;
		var y_i = (3+y-i) % 3;
		if(board[x_i][y_i] != player){
			all_same_d32 = false;
			break;
		}
	}

	return all_same_x || all_same_y || all_same_d14 || all_same_d32;
}

func main(){
	var board = [3][3]rune{};
	for i, board_line := range board {
		for j := range board_line {
			board[i][j] = '-';
		}
	}

	var players = [2]rune{'X', 'O'};
	var game_ended = false;
	for !game_ended {
		for _, player := range players {
			render(&board);

			for {				
				fmt.Printf("Player %c's turn\n", player);
				fmt.Println("Enter an x y pair e.g: 3 1");
				var pos_input string;
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				pos_input = scanner.Text()
				var x_y = strings.Split(pos_input, " ");
				var y, errory = strconv.Atoi(x_y[0]); y--;
				var x, errorx = strconv.Atoi(x_y[1]); x--;
				if(errorx != nil) { fmt.Println("Invalid x"); continue; }
				if(errory != nil) { fmt.Println("Invalid y"); continue; }
				if(x < 0 || x > 3 || y < 0 || y > 3){
					fmt.Printf("x = %d, y = %d is out of board range\n", x+1, y+1);
					continue;
				}
				if(board[x][y] != '-'){
					fmt.Println("position is full");
					continue;
				}
				board[x][y] = player;
				if(won(&board, x, y, player)){ game_ended = true; }
				break;
			}
			if(game_ended){
				fmt.Printf("GAME OVER ! Player %c won\n", player);
				break;
			}
		}
	}
}