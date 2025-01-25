package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, sum, col, move := process(reader)
	fmt.Println(sum)
	if sum >= 0 {
		fmt.Println(col)
		fmt.Println(move)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (board []string, sum int, col int, moves string) {
	n, m, k := readThreeNums(reader)
	board = make([]string, n)
	for i := range n {
		board[i] = readString(reader)[:m]
	}
	sum, col, moves = solve(board, k)
	return
}

const inf = 1 << 60

func solve(board []string, k int) (sum int, col int, moves string) {
	n := len(board)
	m := len(board[0])

	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, m)
		for j := range m {
			dp[i][j] = make([]int, k+1)
			for w := 0; w <= k; w++ {
				dp[i][j][w] = -inf
			}
		}
	}

	get := func(i int, j int) int {
		return int(board[i][j] - '0')
	}

	for j := 0; j < m; j++ {
		dp[n-1][j][get(n-1, j)%(k+1)] = get(n-1, j)
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < m; j++ {
			x := get(i, j)
			for y := 0; y <= k; y++ {
				if j-1 >= 0 {
					dp[i][j][(x+y)%(k+1)] = max(dp[i][j][(x+y)%(k+1)], dp[i+1][j-1][y]+x)
				}
				if j+1 < m {
					dp[i][j][(x+y)%(k+1)] = max(dp[i][j][(x+y)%(k+1)], dp[i+1][j+1][y]+x)
				}
			}
		}
	}
	sum = -1
	best := 0
	for j := 0; j < m; j++ {
		if sum < dp[0][j][0] {
			// valid
			sum = dp[0][j][0]
			best = 0
			col = j
		}
	}
	if sum < 0 {
		return
	}

	var buf []byte

	for r := 0; r+1 < n; r++ {
		val := get(r, col)
		for y := 0; y <= k; y++ {
			if col-1 >= 0 && (y+val)%(k+1) == best && dp[r+1][col-1][y]+val == dp[r][col][best] {
				col--
				best = y
				buf = append(buf, 'R')
				break
			}
			if col+1 < m && (y+val)%(k+1) == best && dp[r+1][col+1][y]+val == dp[r][col][best] {
				col++
				best = y
				buf = append(buf, 'L')
				break
			}
		}
	}

	reverse(buf)
	moves = string(buf)
	col++
	return
}

func reverse(buf []byte) {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}
