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
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(reader, m)
		}
		res := solve(grid)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])
	if (n+m-1)%2 == 1 {
		return false
	}
	// 1 和 -1的个数要相等
	// dp[i][j] & 1 表示=0
	//            2 for 1
	//            4 for -1
	// dp[x][j] 表示运行x步后，sum = j 的情况
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -(n + m)
			dp[i][j][1] = n + m
		}
	}
	dp[0][0][0] = grid[0][0]
	dp[0][0][1] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0]+grid[i][j])
				dp[i][j][1] = min(dp[i][j][1], dp[i-1][j][1]+grid[i][j])
			}
			if j > 0 {
				dp[i][j][0] = max(dp[i][j][0], dp[i][j-1][0]+grid[i][j])
				dp[i][j][1] = min(dp[i][j][1], dp[i][j-1][1]+grid[i][j])
			}
		}
	}

	return dp[n-1][m-1][0] >= 0 && dp[n-1][m-1][1] <= 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
