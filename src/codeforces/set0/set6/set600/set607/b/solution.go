package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

func solve(c []int) int {
	n := len(c)
	// dp[i][j] 表示把i...j都删除所需要的最时间
	// dp[i][j] = min(dp[i][k] + dp[k+1][j]) 这是一种显然的策略
	// 还有一种策略是，删除中间的[l..r] 然后，将剩余的部分组成一个回文后，删除
	//  如果不是这种，就会退回到上面那种策略
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j - i + 1
		}
		dp[i][i] = 1
	}

	for ln := 2; ln <= n; ln++ {
		for j := ln - 1; j < n; j++ {
			i := j - ln + 1
			dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
			if c[i] == c[i+1] {
				dp[i][j] = min(dp[i][j], 1+dp[i+2][j])
			}
			for k := i + 2; k <= j; k++ {
				if c[i] == c[k] {
					dp[i][j] = min(dp[i][j], dp[i+1][k-1]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
