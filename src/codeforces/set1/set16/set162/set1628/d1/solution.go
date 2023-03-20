package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const MOD = 1e9 + 7

func solve(n int, m int, k int) int {
	// bob 必须选择进行m次addition
	// 如果bob已经进行了 n - m次sub, 那么这时候，那么alice只需要提供k即可
	// 但是从bob的角度看，
	//  1. 如果已经进行了m次add，那么接下来所有的操作都可以进行sub， 这时候可以认为已经结束了；
	//  2. 如果alice选择了一个数字num，那么bob选择add/sub？
	// bob肯定要进行 n - m次sub， 和 m次 add， 不管alice怎么选择数字
	// n =  6, m =  3, k = 10  => 75 / 8
	// 如果alice选择x, min(-x + dp[1][m], x + dp[1][m-1])
	// 对于alice来说，应该是选择使得 -x + dp[1][m] 接近 x + dp[1][m-1]的结果
	// x = (dp[1][m] - dp[1][m-1]) / 2

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = modMul(k, i)
	}

	inv2 := inverse(2)

	for i := 2; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			dp[j] = modMul(modAdd(dp[j], dp[j-1]), inv2)
		}
	}

	return dp[m]
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}
