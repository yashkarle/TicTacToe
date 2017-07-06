package main

import (
	"fmt"
)

type TicTacToe struct {
	board [][]string
	turn string
}

func (t *TicTacToe) initBoard() {
	for i,_ := range t.board {
		t.board[i] = make([]string, 3)
		for j, _ := range t.board[i] {
      			t.board[i][j] = "-"
    		}
	}
}

func (t *TicTacToe) printBoard() {
	for i, _ := range t.board {
    		fmt.Println(t.board[i])
  	}	
}

func (t *TicTacToe) isWinner() bool {
	return (t.board[0][0] == t.turn && t.board[0][1] == t.turn && t.board[0][2] == t.turn) ||
               (t.board[1][0] == t.turn && t.board[1][1] == t.turn && t.board[1][2] == t.turn) ||
               (t.board[2][0] == t.turn && t.board[2][1] == t.turn && t.board[2][2] == t.turn) ||
               (t.board[0][0] == t.turn && t.board[1][0] == t.turn && t.board[2][0] == t.turn) ||
               (t.board[0][1] == t.turn && t.board[1][1] == t.turn && t.board[2][1] == t.turn) ||
               (t.board[0][2] == t.turn && t.board[1][2] == t.turn && t.board[2][2] == t.turn) ||
               (t.board[0][0] == t.turn && t.board[1][1] == t.turn && t.board[2][2] == t.turn) ||
               (t.board[0][2] == t.turn && t.board[1][1] == t.turn && t.board[2][0] == t.turn)
}

func (t *TicTacToe) makeMove() {
	var row,col int

	fmt.Println("Which row would you like to mark?")
	fmt.Scanf("%d",&row);
	fmt.Println("Which row would you like to mark?")
	fmt.Scanf("%d",&col);

	t.markSpot(row,col);
}

func main() {
	gameBoard := TicTacToe{make([][]string,3),"X"}
	gameBoard.initBoard();
	gameBoard.printBoard();

	for i := 0; i < 9; i++ {
    		if gameBoard.isWinner() {
      			break;
    		}
    		gameBoard.makeMove()
  	}

	fmt.Println(gameBoard);
}
