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
		reader.ReadBytes('\n')
		n, k := readTwoNums(reader)
		blocked := make([][]int, k)
		for i := 0; i < k; i++ {
			blocked[i] = readNNums(reader, 2)
		}
		res := solve(n, blocked)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, blocks [][]int) []int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	for _, blk := range blocks {
		x, y := blk[0], blk[1]
		x--
		y--
		board[x][y] = -1
	}

	safe := func(r, c int) bool {
		// put a queue at position (r, c)
		// no need to check same column or column after c
		for j := c - 1; j >= 0; j-- {
			if board[r][j] < 0 {
				// blocked cell
				break
			}
			if board[r][j] == 1 {
				// meet another queen
				return false
			}
		}

		// top-left diagonal
		for j := 1; ; j++ {
			x := r - j
			y := c - j
			if x < 0 || y < 0 {
				break
			}
			if board[x][y] < 0 {
				// blocked cell
				break
			}
			if board[x][y] == 1 {
				// meet another queen
				return false
			}
		}
		// bottom-left
		for j := 1; ; j++ {
			x := r + j
			y := c - j
			if x == n || y < 0 {
				break
			}
			if board[x][y] < 0 {
				// blocked cell
				break
			}
			if board[x][y] == 1 {
				// meet another queen
				return false
			}
		}

		return true
	}

	var dfs func(c int) bool

	dfs = func(c int) bool {
		if c == n {
			return true
		}

		for r := 0; r < n; r++ {
			if board[r][c] == 0 && safe(r, c) {
				board[r][c] = 1
				if dfs(c + 1) {
					return true
				}
				board[r][c] = 0
			}
		}
		return false
	}

	if dfs(0) {
		res := make([]int, n)
		for c := 0; c < n; c++ {
			for r := 0; r < n; r++ {
				if board[r][c] == 1 {
					res[c] = r + 1
					break
				}
			}
		}
		return res
	}
	return nil
}
