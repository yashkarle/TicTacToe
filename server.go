package main

import "net"
import "fmt"
import "bufio"
import "strconv"
import "encoding/json"

func trimString(s1 string) (s string) {
	last := len(s1) - 1
	s = s1[:last]
	return s
}

func echo(conn net.Conn) (row string, col string) {
	r := bufio.NewReader(conn)

	line1, err := r.ReadBytes(byte('\n'))
	line2, err := r.ReadBytes(byte('\n'))

	if err != nil {
		fmt.Println("ERROR", err)
	}

	conn.Write(line1)
	conn.Write(line2)

	row1 := string(line1)
	col1 := string(line2)

	row = trimString(row1)
	col = trimString(col1)

	return row, col
}

type TicTacToe struct {
	board [][]string
	turn  string
}

func (t *TicTacToe) initBoard() {
	for i, _ := range t.board {
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

func (t *TicTacToe) markSpot(row int, col int) {
	if t.board[row][col] != "-" {
		return
	}

	t.board[row][col] = t.turn
	//t.printBoard()

	if t.isWinner() {
		fmt.Println(t.turn + " wins!")
		return
	}

	if t.turn == "X" {
		t.turn = "O"
	} else {
		t.turn = "X"
	}
}

func (t *TicTacToe) makeMove() {
	var row, col int

	fmt.Println("Which row would you like to mark?")
	fmt.Scanf("%d", &row)
	fmt.Println("Which column would you like to mark?")
	fmt.Scanf("%d", &col)

	t.markSpot(row, col)
}

func main() {

	fmt.Println("Waiting for other player to join...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port
	conn, _ := ln.Accept()

	gameBoard := TicTacToe{make([][]string, 3), "X"}
	gameBoard.initBoard()
	gameBoard.printBoard()

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			if gameBoard.isWinner() {
				break
			}
			gameBoard.makeMove()

			conn1, _ := net.Dial("tcp", "localhost:8082")
			json.NewEncoder(conn1).Encode(gameBoard.board)

		} else {
			fmt.Println("After O's turn: \n")
			row1, col1 := echo(conn)
			row, _ := strconv.Atoi(row1)
			col, _ := strconv.Atoi(col1)
			gameBoard.markSpot(row, col)
			gameBoard.printBoard()
		}
	}
	// run loop forever (or until ctrl-c)
}
