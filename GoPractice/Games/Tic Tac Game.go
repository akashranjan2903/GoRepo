package Games

import (
	"errors"
	"fmt"
)

var (
	player1    string = "X"
	player2    string = "O"
	emptySpace string = "0"
)

var board [3][3]string

// newBoard -  creates a new board
func newBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = emptySpace
		}
	}
}

// printBoard - prints the board
func printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" ")
			if (board[i][j]) == "X" || board[i][j] == "O" {
				fmt.Printf("%s", board[i][j])

			} else if board[i][j] == "0" {
				fmt.Printf(" ")
			}
			if j < 2 {
				fmt.Printf(" |")
			}

		}
		fmt.Println()
		fmt.Println("-----------")
	}
}
func checkBoardInput(row, col int) error {

	if row < 1 || row > 3 || col < 1 || col > 3 {
		return errors.New("entered Postion is out of Bound")
	}
	if board[row-1][col-1] != emptySpace {
		return errors.New("position You Entered is Already Filled")
	}
	return nil
}
func takeinput() (int, int) {
	var row, col int
	for {
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}
		e := checkBoardInput(row, col)
		if e != nil {
			fmt.Println(e)
		} else {
			return row, col
		}
		fmt.Println("Please Try Again!!")
	}
}
func remainingpos() {
	println("Remaining Postion to be filled are:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "0" {
				fmt.Printf("(%d,%d)", i+1, j+1)
			}
		}

	}
	fmt.Println("")
}
func Playgame() {
	newBoard()
	fmt.Println("Welcome to Tic Tac Toe")
	fmt.Println("\nCurrent Board :")
	currentPLayer := player1
	for {
		printBoard()
		remainingpos()
		fmt.Printf("\nHello Player %s, Please enter your move in row and column format (e.g. 1 2): ", currentPLayer)
		// Taking Input Each Time
		var row, col = takeinput()
		updateBoard(row, col, currentPLayer)

		// Check if someone won the game
		player := checkWon()
		if player == "X" || player == "O" {
			// updateBoard(row, col, currentPLayer)
			printBoard()
			fmt.Println("Congrats Player ", player, ", You won the game")
			break
		}

		// Check if the game is draw
		isDraw := checkDraw()
		if isDraw {
			// updateBoard(row, col, currentPLayer)
			printBoard()
			fmt.Println("Game is draw")
			break
		}

		// Alternate the players in game
		if currentPLayer == player1 {
			currentPLayer = player2
		} else {
			currentPLayer = player1
		}
	}
}

// updateBoard - updates the board with the move
func updateBoard(row, col int, player string) {
	board[row-1][col-1] = player
}

func checkWon() string {
	// Check row
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] &&
			board[i][0] != emptySpace {
			return board[i][0]
		}
	}

	// Check column
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] &&
			board[1][i] == board[2][i] &&
			board[0][i] != emptySpace {
			// fmt.Println("Player ", board[0][i], " won the game")
			return board[0][i]
		}
	}

	// Check diagonal
	if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] &&
		board[0][0] != emptySpace {
		// fmt.Println("Player ", board[0][0], " won the game")
		return board[0][0]
	}

	return ""
}

// Check Draw - check if the game is draw
func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == emptySpace {
				return false
			}
		}
	}
	return true
}
