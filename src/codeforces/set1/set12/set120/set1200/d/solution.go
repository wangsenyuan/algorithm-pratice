package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	return solve(grid, k)
}

func solve(grid []string, k int) int {
	n := len(grid)

	// row[i][0]是这一行内部，最小的black的区域
	// row[i][1]是这一行内部，最大的black的区域
	row := make([][]int, n)
	col := make([][]int, n)
	dp := make([][]int, n)
	for i := range n {
		row[i] = []int{n, 0}
		col[i] = []int{n, 0}
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'B' {
				row[i][0] = min(row[i][0], j)
				row[i][1] = j
				col[j][0] = min(col[j][0], i)
				col[j][1] = i
			}
		}
	}

	pr := make([]int, n)
	pc := make([]int, n)
	for i := 0; i < n; i++ {
		if row[i][0] > row[i][1] {
			pr[i] = 1
		}
		if col[i][0] > col[i][1] {
			pc[i] = 1
		}
		if i > 0 {
			pr[i] += pr[i-1]
			pc[i] += pc[i-1]
		}
	}

	best := pr[n-1] + pc[n-1]

	getCol := func(j int, i int) int {
		if col[j][0] > col[j][1] || col[j][0] >= i && col[j][1] < i+k {
			return 1
		}
		return 0
	}

	for i := 0; i+k <= n; i++ {
		// 确定了上边界的情况下
		var cnt int
		for j := 0; j < n; j++ {
			if j-k >= 0 {
				cnt -= getCol(j-k, i)
			}

			cnt += getCol(j, i)

			if j-k+1 >= 0 {
				dp[i][j-k+1] = cnt + pc[n-1] - pc[j] + pr[n-1] - pr[i+k-1]
				if i > 0 {
					dp[i][j-k+1] += pr[i-1]
				}
				if j-k+1 > 0 {
					dp[i][j-k+1] += pc[j-k]
				}
			}
		}
	}

	getRow := func(i int, j int) int {
		if row[i][0] > row[i][1] || row[i][0] >= j && row[i][1] < j+k {
			return 1
		}
		return 0
	}

	for j := 0; j+k <= n; j++ {
		var cnt int
		for i := 0; i < n; i++ {
			if i-k >= 0 {
				cnt -= getRow(i-k, j)
			}
			cnt += getRow(i, j)

			if i-k+1 >= 0 {
				dp[i-k+1][j] += cnt
			}
		}
	}

	for i := 0; i+k <= n; i++ {
		for j := 0; j+k <= n; j++ {
			best = max(best, dp[i][j])
		}
	}

	return best
}
