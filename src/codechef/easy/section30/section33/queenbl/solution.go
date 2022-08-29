package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		king := readNNums(reader, 2)
		res := solve(king)

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(king []int) [][]int {
	x, y := king[0], king[1]
	x--
	y--
	board := make([][]int, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]int, 8)
	}

	board[x][y] = 1

	var cnt int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i >= 0 && x+i < 8 && y+j >= 0 && y+j < 8 {
				cnt++
			}
		}
	}

	placeQueen := func(u, v int) map[int]int {
		res := make(map[int]int)
		for d := 1; u-d >= 0; d++ {
			i := u - d
			j := v
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}
		for d := 1; u+d < 8; d++ {
			i := u + d
			j := v
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}

		for d := 1; v-d >= 0; d++ {
			i := u
			j := v - d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}

		for d := 1; v+d < 8; d++ {
			i := u
			j := v + d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}

		for d := 1; u-d >= 0 && v-d >= 0; d++ {
			i := u - d
			j := v - d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}
		for d := 1; u+d < 8 && v+d < 8; d++ {
			i := u + d
			j := v + d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}

		for d := 1; u-d >= 0 && v+d < 8; d++ {
			i := u - d
			j := v + d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}
		for d := 1; u+d < 8 && v-d >= 0; d++ {
			i := u + d
			j := v - d
			if abs(i-x) <= 1 && abs(j-y) <= 1 {
				res[i*8+j]++
			}
		}
		return res
	}

	expect := 1
	if x > 0 && x < 7 || y > 0 && y < 7 {
		expect++
	}

	for a := 0; a < 64; a++ {
		x1, y1 := a/8, a%8
		if attack(x1, y1, x, y) {
			continue
		}

		block := placeQueen(x1, y1)

		if len(block) == cnt {
			board[x1][y1] = 2
			return board
		}
		if expect > 1 {

			for b := 0; b < 64; b++ {
				x2, y2 := b/8, b%8
				if attack(x2, y2, x, y) {
					continue
				}
				clock := placeQueen(x2, y2)
				for k := range block {
					clock[k]++
				}
				if len(clock) == cnt {
					board[x1][y1] = 2
					board[x2][y2] = 2
					return board
				}
			}
		}
	}

	return board
}

func attack(x1, y1, x2, y2 int) bool {
	if x1 == x2 || y1 == y2 {
		return true
	}
	dx := x1 - x2
	dy := y1 - y2
	return abs(dx) == abs(dy)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
