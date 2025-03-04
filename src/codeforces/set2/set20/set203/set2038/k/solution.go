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
	n, a, b := readThreeNums(reader)
	return solve(n, a, b)
}

const inf = 1 << 60

func solve(n int, a int, b int) int {
	// gcd(x, a) = 1

	get := func(x int) int {
		m := n
		for gcd(m, x) > 1 {
			m--
		}
		return m
	}
	x := get(a)
	y := get(b)
	// x行y列
	res := x - 1 + y - 1
	for i := 1; i <= x; i++ {
		res += gcd(i, a)
	}
	for i := 1; i <= y; i++ {
		res += gcd(i, b)
	}

	dp := make([][]int, n-x+1)
	for i := 0; i <= n-x; i++ {
		dp[i] = make([]int, n-y+1)
		for j := 0; j <= n-y; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := x; i <= n; i++ {
		for j := y; j <= n; j++ {
			if i == x && j == y {
				continue
			}
			tmp := gcd(i, a) + gcd(j, b)
			if i > x {
				dp[i-x][j-y] = min(dp[i-x][j-y], dp[i-x-1][j-y]+tmp)
			}
			if j > y {
				dp[i-x][j-y] = min(dp[i-x][j-y], dp[i-x][j-y-1]+tmp)
			}
		}
	}

	return dp[n-x][n-y] + res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
