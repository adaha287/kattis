package main

import (
	"fmt"
)

func main() {
	var boardLength, index, magic, hops int
	_, _ = fmt.Scanln(&boardLength, &index, &magic)

	board, err := read(boardLength)
	if err != nil {
		fmt.Println("ERROR!")
	}
	if magic == board[index-1] {
		fmt.Println("magic\n0")
		return
	}
	visited := make(map[int]string)
	visited[index] = ""
	for true {
		index = index + board[index-1]
		hops += 1

		if index < 1 {
			fmt.Printf("left\n%d", hops)
			return
		} else if index > len(board) {
			fmt.Printf("right\n%d", hops)
			return
		} else if board[index-1] == magic {
			fmt.Printf("magic\n%d", hops)
			return
		} else if _, ok := visited[index]; ok {
			fmt.Printf("cycle\n%d", hops)
			return
		}
		visited[index] = ""
	}
}

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return nil, err
		}
	}
	return in, nil
}
