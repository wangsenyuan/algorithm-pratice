package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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

const inf = 1 << 50

func solve(a []int) int {
	n := len(a)
	// (k + 1) * k / 2 <= n
	k := sort.Search(n, func(k int) bool {
		return (k+1)*k/2 > n
	})
	k--
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + a[i]
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[n][0] = inf

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i+1][j]
			if i+j <= n && sum[i]-sum[i+j] < dp[i+j][j-1] {
				dp[i][j] = max(dp[i][j], sum[i]-sum[i+j])
			}
		}
		dp[i][0] = inf
	}

	for j := k; j > 0; j-- {
		if dp[0][j] > 0 {
			return j
		}
	}
	return 1
}
