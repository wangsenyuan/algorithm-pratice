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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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
	n, m := readTwoNums(reader)
	mat := make([]string, n)
	for i := range n {
		mat[i] = readString(reader)[:m]
	}
	return solve(mat)
}

const mod = 1_000_000_007

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(mat []string) int {
	n := len(mat)
	m := len(mat[0])

	if mat[n-1][m-1] == 'R' {
		return 0
	}

	if n == 1 {
		for j := 0; j < m; j++ {
			if mat[0][j] == 'R' {
				return 0
			}
		}
		return 1
	}
	if m == 1 {
		for i := 0; i < n; i++ {
			if mat[i][0] == 'R' {
				return 0
			}
		}
		return 1
	}

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, m+1)
	}

	dp[n-1][m-1][0] = 1
	dp[n-1][m-1][1] = 1

	down := make([]int, m)
	col := make([]int, m)
	for j := 0; j < m; j++ {
		col[j] = n - 1
	}

	for i := n - 1; i >= 0; i-- {
		var right int
		r := m - 1
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				right = 1
				down[j] = 1
				continue
			}

			dp[i][j][1] = right
			dp[i][j][0] = down[j]

			right = add(right, dp[i][j][0])
			down[j] = add(down[j], dp[i][j][1])

			if mat[i][j] == 'R' {
				right = sub(right, dp[i][r][0])
				r--
				down[j] = sub(down[j], dp[col[j]][j][1])
				col[j]--
			}
		}
	}

	return add(dp[0][0][0], dp[0][0][1])
}
