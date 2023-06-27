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
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
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

const MOD = 998244353

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func solve(n int, m int, k int) int {
	// n 块砖，m中颜色, 要求切好有k块砖和左边的不一致的情况
	// dp[i][j] 表示左边时颜色i时，且已有j块和左边颜色不一致的计数
	// dp[ii][j] += dp[i][j] 如果 i == ii, else dp[i][j-1] 如果 i != ii
	// dp[j] = dp[j] + dp[j-1]
	dp := make([]int, k+1)
	dp[0] = m
	for i := 1; i < n; i++ {
		for j := k; j > 0; j-- {
			dp[j] = add(dp[j], mul(m-1, dp[j-1]))
		}
		//dp[0] = mul(dp[0], m)
	}

	return dp[k]
}
