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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
	n, m, x := readThreeNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a, x)
}

const inf = 1 << 60

func solve(a [][]int, x int) int {
	n := len(a)
	m := len(a[0])
	// dp[i][j] 表示到达dp[i][j]时的最小值
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = inf
	}
	// k是转动的次数
	for k := 0; k < m; k++ {
		var sum int
		for j := 0; j < m; j++ {
			nj := (j + k + m) % m
			sum += a[0][nj]
			dp[j] = min(dp[j], sum+k*x)
		}
	}
	fp := make([]int, m)

	gp := make([]int, m)

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			fp[j] = inf
		}

		for k := 0; k < m; k++ {
			for j := 0; j < m; j++ {
				nj := (j + k + m) % m
				gp[j] = dp[j] + a[i][nj]
				if j > 0 {
					gp[j] = min(gp[j], gp[j-1]+a[i][nj])
				}
			}
			for j := 0; j < m; j++ {
				gp[j] += k * x
				fp[j] = min(fp[j], gp[j])
			}
		}

		copy(dp, fp)
	}

	return dp[m-1]
}