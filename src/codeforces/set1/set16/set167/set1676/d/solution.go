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
		n, m := readTwoNums(reader)
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			board[i] = readNNums(reader, m)
		}
		res := solve(board)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(board [][]int) int {
	// dp[i][j] = board[i][j] + dp[i-1][j-1]
	n := len(board)
	m := len(board[0])
	dp := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, m)
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			dp[0][r][c] = board[r][c]
			if r > 0 && c > 0 {
				dp[0][r][c] += dp[0][r-1][c-1]
			}
			dp[1][r][c] = board[r][c]
			if r > 0 && c+1 < m {
				dp[1][r][c] += dp[1][r-1][c+1]
			}
		}
	}

	for r := n - 1; r >= 0; r-- {
		for c := 0; c < m; c++ {
			dp[2][r][c] = board[r][c]
			if r+1 < n && c > 0 {
				dp[2][r][c] += dp[2][r+1][c-1]
			}
			dp[3][r][c] = board[r][c]
			if r+1 < n && c+1 < m {
				dp[3][r][c] += dp[3][r+1][c+1]
			}
		}
	}
	var ans int
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			var tmp int
			for i := 0; i < 4; i++ {
				tmp += dp[i][r][c]
			}
			tmp -= 3 * board[r][c]
			if ans < tmp {
				ans = tmp
			}
		}
	}
	return ans
}
