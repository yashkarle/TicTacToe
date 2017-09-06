package main

import "os"
import "net"
import "fmt"
import "bufio"
import "encoding/json"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:8081")

	ln, _ := net.Listen("tcp", ":8082")

	for {
		conn1, _ := ln.Accept()

		var board [][]string
		json.NewDecoder(conn1).Decode(&board)

		fmt.Println("After X's turn: \n")
		for i, _ := range board {
			fmt.Println(board[i])
		}

		userInput := bufio.NewReader(os.Stdin)
		// read in input from stdin
		fmt.Println("Which row would you like to mark?")
		row, err := userInput.ReadBytes(byte('\n'))
		fmt.Println("Which column would you like to mark?")
		col, err := userInput.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("ERROR", err)
		}
		// send to socket
		conn.Write(row)
		conn.Write(col)
		//fmt.Fprintf(conn, col + "\n")
		// listen for reply
	}
}
