package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	res := solve(n, k)
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, k int) int {
	dp := make([]int, 2*n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			dp[j] = add(dp[j-1], dp[2*j])
		}
		dp[0] = 0
	}

	return dp[k]
}

func solve2(n int, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2*n+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			dp[i][j] = add(dp[i-1][j-1], dp[i][2*j])
		}
	}

	return dp[n][k]
}

func solve1(n int, k int) int {
	dp := []int{1}
	for i := 1; i <= n; i++ {
		dp = append([]int{0}, dp...)
		for j := i / 2; j > 0; j-- {
			dp[j] = add(dp[j], dp[2*j])
		}
	}

	return dp[k]
}
