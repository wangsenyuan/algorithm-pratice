package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)
		grid := make([][]int, m)
		for i := 0; i < m; i++ {
			grid[i] = readNNums(reader, n)
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(grid [][]int) int {
	// 假设从dp(i, j) alice 从 i, j 出发到达右下角的最大值是 x
	// fp(i, j) = becky 从i, j出发到达右下角的最大值
	// fp(i-1, j) = max(grid[i-1][j], dp[i][j])
	// fp(i, j-1) = max(grid[i][j-1], dp[i][j])
	// dp(i - 1, j) = max(grid[i-1][j], fp[i][j]) when grid[i][j] >= grid[i-1][j]
	// fp(i, j - 1) = max(grid[i][j-1], fp(i)[j]) when grid[i][j] >= grid[i+1][j]

	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	fp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = grid[i][j]
			fp[i][j] = grid[i][j]

			if i+1 < m {
				// can go down
				if j+1 == n || grid[i+1][j] >= grid[i][j+1] {
					dp[i][j] = max(dp[i][j], fp[i+1][j])
				}
			}
			if j+1 < n {
				// can go right
				if i+1 == m || grid[i+1][j] <= grid[i][j+1] {
					dp[i][j] = max(dp[i][j], fp[i][j+1])
				}
			}

			if i+1 == m && j+1 < n {
				// can only go right
				fp[i][j] = max(fp[i][j], dp[i][j+1])
			}
			if i+1 < m && j+1 == n {
				// can only go down
				fp[i][j] = max(fp[i][j], dp[i+1][j])
			}
			if i+1 < m && j+1 < n {
				fp[i][j] = max(fp[i][j], min(dp[i+1][j], dp[i][j+1]))
			}
		}
	}

	return fp[0][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
