package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	forest := make([]string, n)
	for i := 0; i < n; i++ {
		forest[i] = readString(reader)[:m]
	}
	res := solve(forest)
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

func solve(forest []string) int {
	n := len(forest)
	m := len(forest[0])
	// n, m <= 500
	// dp[i][j][u][v] = 从两端往中间移动，且始终是palindrome的方案数
	// 500 * 500 * 500 * 500 = 125 * 1e8 显然太大了
	// 但是最后一个v其实可以不要的, 因为 i + j = (n - 1 - u) + (m - 1 - v)
	// => v = n + m - 2 - (i + j + u)
	// 500 * 500 * 500 = 125 * 1e6 ~ 1e8
	if forest[0][0] != forest[n-1][m-1] {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+2)
	}
	dp[1][n] = 1
	for i := 1; i < (n+m)/2; i++ {
		for r1 := n; r1 > 0; r1-- {
			for r2 := 1; r2 <= n; r2++ {
				c1, c2 := i+2-r1, m+n-i-r2
				if c1 > 0 && c1 <= m && c2 > 0 && c2 <= m {
					if forest[r1-1][c1-1] == forest[r2-1][c2-1] {
						dp[r1][r2] = modAdd(dp[r1][r2], dp[r1][r2+1], dp[r1-1][r2+1], dp[r1-1][r2])
					} else {
						dp[r1][r2] = 0
					}
				}
			}
		}
	}

	var res int
	if (n+m)%2 == 1 {
		for i := 1; i <= n; i++ {
			res = modAdd(res, dp[i][i], dp[i][i+1])
		}
	} else {
		for i := 1; i <= n; i++ {
			res = modAdd(res, dp[i][i])
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const MOD = 1e9 + 7

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res += arr[i]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}
