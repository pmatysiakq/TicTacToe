package Game

import (
	"TicTacToe/Board"
	"TicTacToe/Player"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type game struct {
	player1  Player.Player
	player2  Player.Player
	board    Board.Board
	nextMove Player.Player
	winner   Player.Player
	isTie    bool
}

func newGame(p1_name, p2_name, p1_sign, p2_sign string) *game {
	return &game{
		player1: Player.Player{
			Name: p1_name,
			Sign: p1_sign,
		},
		player2: Player.Player{
			Name: p2_name,
			Sign: p2_sign,
		},
		board:    Board.Board{},
		nextMove: Player.Player{},
		winner:   Player.Player{},
		isTie:    false,
	}
}

func CreateNewGame() *game {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nPlease ENTER 1st Player name > ")
	scanner.Scan()
	p1_name := fmt.Sprintf("%s", scanner.Text())
	fmt.Print("\nPlease ENTER 2nd Player name > ")
	scanner.Scan()
	p2_name := fmt.Sprintf("%s", scanner.Text())

	value := rand.Float32()

	p1_sign, p2_sign := "O", "X"

	if value < 0.5 {
		p1_sign = "O"
		p2_sign = "X"
	} else {
		p1_sign = "X"
		p2_sign = "O"
	}

	return newGame(p1_name, p2_name, p1_sign, p2_sign)
}

func (g *game) makeMove() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var err error
		answer := -1
		for answer < 0 || answer > 9 {
			g.board.GetBoardString()
			fmt.Printf("\n%v Please ENTER number in range 1-9 > ", g.nextMove)
			scanner.Scan()
			answer, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("Value %v is not an int! Try again!", answer)
				continue
			} else if answer <= 0 && answer >= 9{
				fmt.Printf("Value %d is not in range 1-9. Try Again!", answer)
				continue
			}
		}
		if answer == 1 {
			err = g.board.UpdateBoard(0, 0, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 2 {
			err = g.board.UpdateBoard(0, 1, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 3 {
			err = g.board.UpdateBoard(0, 2, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 4 {
			err = g.board.UpdateBoard(1, 0, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 5 {
			err = g.board.UpdateBoard(1, 1, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 6 {
			err = g.board.UpdateBoard(1, 2, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 7 {
			err = g.board.UpdateBoard(2, 0, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 8 {
			err = g.board.UpdateBoard(2, 1, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else if answer == 9 {
			err = g.board.UpdateBoard(2, 2, g.nextMove.Sign)
			if err != nil {
				fmt.Printf("Error occured: %v. Try again!", err)
				continue
			}
		} else {
			g.board.GetBoardString()
			fmt.Printf("Please ENTER correct value!")
			continue
		}
		break
	}
}

func (g *game) isGameOver() bool{
	if g.winner == g.player1 || g.winner == g.player2 || g.isTie == true {
		return true
	} else {
		return false
	}
}

func (g *game) checkTie() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board.GetBoardArray()[i][j] == "" {
				g.isTie = false
				return
			}
		}
	}
	g.isTie = true
}

func (g *game) checkWinner() {
	b := g.board.GetBoardArray()
	if        (b[0][0] == b[0][1] && b[0][0] == b[0][2]) && (b[0][0] != ""){
		if g.player1.Sign == b[0][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[1][0] == b[1][1] && b[1][2] == b[1][2]) && (b[1][0] != ""){
		if g.player1.Sign == b[1][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[2][0] == b[2][1] && b[2][2] == b[2][2]) && (b[2][0] != ""){
		if g.player1.Sign == b[2][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[0][0] == b[1][0] && b[0][0] == b[2][0]) && (b[0][0] != ""){
		if g.player1.Sign == b[0][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[0][1] == b[1][1] && b[0][1] == b[2][1]) && (b[0][1] != ""){
		if g.player1.Sign == b[0][1] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[0][2] == b[1][2] && b[0][2] == b[2][2]) && (b[0][2] != ""){
		if g.player1.Sign == b[0][2] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[0][0] == b[1][1] && b[0][0] == b[2][2]) && (b[0][0] != ""){
		if g.player1.Sign == b[0][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	} else if (b[2][0] == b[1][1] && b[2][0] == b[0][2]) && (b[2][0] != ""){
		if g.player1.Sign == b[2][0] {
			g.winner = g.player1
		} else {
			g.winner = g.player2
		}
	}
}

func (g *game) setNextMove() {
	if g.nextMove == g.player1 {
		g.nextMove = g.player2
	} else if g.nextMove== g.player2 {
		g.nextMove = g.player1
	} else {
		if g.player1.Sign == "0" {
			g.nextMove = g.player1
		} else {
			g.nextMove = g.player2
		}
	}
}

func (g *game) showInstruction() {
	fmt.Println("Hello. Welcome to the Tic Tac Toe game. There are two players. You will see board with numbers. If you choose field, where you want to put sign, press proper number.")
}

func (g *game) StartGame() {
	g.showInstruction()

	for !g.isGameOver() {
		g.setNextMove()
		g.makeMove()
		g.checkWinner()
		g.checkTie()
	}
	if g.isTie {
		fmt.Printf("It's a tie!")
	} else {
		fmt.Printf("The winner is %v!", g.winner)
	}
}