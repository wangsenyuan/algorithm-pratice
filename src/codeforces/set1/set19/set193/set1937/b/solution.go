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
		readNum(reader)
		grid := make([]string, 2)
		grid[0] = readString(reader)
		grid[1] = readString(reader)
		res, cnt := solve(grid)
		buf.WriteString(fmt.Sprintf("%s\n%d\n", res, cnt))
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

func solve(grid []string) (string, int) {
	n := len(grid[0])

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	// 也就是只有到 s[0][i+1] == s[1][i]的时候，才能继续
	for i := 1; i < n; i++ {
		if dp[0][i-1] > 0 {
			if grid[0][i] == grid[1][i-1] {
				dp[0][i] = dp[0][i-1]
				dp[1][i-1] += dp[0][i-1]
			} else if grid[0][i] < grid[1][i-1] {
				dp[0][i] = dp[0][i-1]
				dp[1][i-1] = 0
			} else {
				dp[0][i] = 0
				dp[1][i-1] += dp[0][i-1]
			}
			dp[1][i] += dp[1][i-1]
		} else {
			dp[1][i] = dp[1][i-1]
		}
	}

	dp[1][n-1] += dp[0][n-1]

	buf := make([]byte, n+1)
	a, b := 1, n-1
	for i := n; i >= 0; i-- {
		buf[i] = grid[a][b]
		if b == 0 || a == 1 && dp[a-1][b] > 0 {
			a--
		} else {
			b--
		}
	}

	return string(buf), dp[1][n-1]
}
