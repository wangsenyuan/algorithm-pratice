package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	chess := strings.Split(s, " ")
	res := solve(chess)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}
func solve(chess []string) string {
	parse := func(s string) []int {
		x := int(s[0] - 'a')
		y := int(s[1] - '1')
		return []int{x, y}
	}
	rock1 := parse(chess[0])
	rock2 := parse(chess[1])
	king1 := parse(chess[2])
	king2 := parse(chess[3])

	board := make([][]int, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]int, 8)
	}

	board[rock1[0]][rock1[1]] = 1
	board[rock2[0]][rock2[1]] = 1
	board[king1[0]][king1[1]] = 1

	checkRock := func(rock []int, x int, y int) bool {
		u, v := rock[0], rock[1]
		if board[u][v] != 1 || u != x && v != y {
			return false
		}

		if u == x {
			if v > y {
				v, y = y, v
			}
			for i := v + 1; i < y; i++ {
				if board[x][i] != 0 {
					return false
				}
			}
		} else {
			if u > x {
				u, x = x, u
			}
			for i := u + 1; i < x; i++ {
				if board[i][y] != 0 {
					return false
				}
			}
		}

		return true
	}

	checkKing := func(king []int, x int, y int) bool {
		u, v := king[0], king[1]
		if board[u][v] != 1 {
			return false
		}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				nu, nv := u+dx, v+dy
				if nu == x && nv == y {
					return true
				}
			}
		}
		return false
	}

	// let's move king2
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			x, y := king2[0]+dx, king2[1]+dy
			if x < 0 || x >= 8 || y < 0 || y >= 8 {
				continue
			}
			val := board[x][y]
			board[x][y] = 2
			if checkRock(rock1, x, y) || checkRock(rock2, x, y) || checkKing(king1, x, y) {
				board[x][y] = val
				continue
			}
			return "OTHER"
		}
	}

	return "CHECKMATE"
}
