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
	n, x := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(x, a)
}

const inf = 1 << 60

func solve(x int, a []int) int {
	n := len(a)
	fp := make([]int, n+1)
	var res int
	for i := 0; i < n; i++ {
		fp[i] = a[i]
		if i > 0 {
			fp[i] += max(0, fp[i-1])
		}
		fp[i] = max(fp[i], 0)
		res = max(res, fp[i])
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(max(dp[i+1], 0)+a[i], 0)
	}

	// fp[l]˝ + x * (sum[r] - sum[l]) + dp[r]
	best := 0
	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
		tmp := dp[i+1] + x*sum + best
		res = max(res, tmp)
		best = max(best, fp[i]-x*sum)
	}

	return res
}

func solve1(x int, a []int) int {
	n := len(a)

	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				dp[i][j][k] = -inf
			}
		}
	}

	f := func(j int, i int) int {
		if j == 1 {
			return a[i]
		}
		return 0
	}

	g := func(k int) int {
		if k == 1 {
			return x
		}
		return 1
	}

	dp[0][0][0] = 0
	for i := 0; i <= n; i++ {
		for j := range 3 {
			for k := range 3 {
				if k < 2 {
					dp[i][j][k+1] = max(dp[i][j][k+1], dp[i][j][k])
				}
				if j < 2 {
					dp[i][j+1][k] = max(dp[i][j+1][k], dp[i][j][k])
				}
				if i < n {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j][k]+f(j, i)*g(k))
				}
			}
		}
	}

	return dp[n][2][2]
}
