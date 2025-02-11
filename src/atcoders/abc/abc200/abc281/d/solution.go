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
	n, k, d := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(k, d, a)
}

const inf = 1 << 60

func solve(k int, d int, a []int) int {
	n := len(a)

	dp := make([][]int, k+1)

	for i := 0; i <= k; i++ {
		dp[i] = make([]int, d)
		for j := 0; j < d; j++ {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for w := min(i, k-1); w >= 0; w-- {
			for s := 0; s < d; s++ {
				ns := (s + a[i]) % d
				dp[w+1][ns] = max(dp[w+1][ns], dp[w][s]+a[i])
			}
		}
	}

	return max(dp[k][0], -1)
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