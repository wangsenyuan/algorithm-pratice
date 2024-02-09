package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	res := solve(mat, k)
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

func solve(mat [][]int, k int) int {
	n := len(mat)
	m := len(mat[0])

	dp := make([]int, m+2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for x := min(m, i+k-1); x >= max(i, k); x-- {
				dp[x+1] = max(dp[x+1], dp[x]+mat[j-1][i-1])
			}
		}
	}

	var res int
	for i := 1; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func solve2(mat [][]int, k int) int {
	n := len(mat)
	m := len(mat[0])
	mx := m - k + 1

	dp := make([][][]int, mx+1)
	for i := 0; i <= mx; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, k+1)
			for x := 0; x <= k; x++ {
				dp[i][j][x] = -1
			}
		}
	}

	var f func(j int, x int, y int) int

	f = func(j int, x int, y int) int {
		if j > mx {
			return 0
		}
		if x > n {
			x = 1
			y++
		}
		if y < 0 {
			x = 1
			y = 0
		}
		if y >= k || j+y > m {
			return -(1 << 30)
		}
		if dp[j][x][y] < 0 {
			dp[j][x][y] = max(f(j, x+1, y), mat[x-1][j+y-1]+f(j+1, x+1, y-1))
		}

		return dp[j][x][y]
	}

	return f(1, 1, 0)
}
func solve1(mat [][]int, k int) int {
	n := len(mat)
	m := len(mat[0])

	arr := make([]int, n*m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[(j-1)*n+i] = mat[i-1][j-1]
		}
	}

	dp := make([]int, n*k+1)

	for i := 1; i <= m-k+1; i++ {
		var x int
		for j := 1; j <= n*k; j++ {
			x = max(x, dp[min(j+n-1, n*k)]+arr[(i-1)*n+j])
			dp[j] = x
		}
	}

	return dp[n*k]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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
