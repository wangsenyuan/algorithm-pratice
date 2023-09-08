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
		n, m := readTwoNums(reader)
		res := solve(n, m)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, m int) int64 {
	// 在给u, v 增加edge时，需要满足 w = gcd(u, v)
	// 对于w，一次需要添加w-1条，且如果u, v已经存在一条边，不能再加一条边
	// min cost
	// gcd(u, v) = 2，假设有x对，那么就可以消耗x条边
	// gcd(u, v) = 3 的话如果有x对，那么可以消耗 x / 2条边
	// gcd(u, v) = 4的话，x,消耗 x / 3...
	// 对于数字1，2，3.。。n
	// 需要直到 gcd(u, v) = g的个数
	// 这个好像可以用inclusion-exclusion进行处理
	// dp[x]表示x...n能够整除x的个数，显然，以x为gcd的数对，必然就是这些乘数
	// 不考虑其他情况的话，m * (m - 1) / 2
	// 但是需要排除掉 dp[2 * x], dp[3 * x]
	dp := make([]int, n+1)
	dp[n] = 0
	for i := n - 1; i > 1; i-- {
		// i, i + 1, i + 2, ... n
		// n = 4, i = 2
		// 2, 3, 4
		l := (n - i + 1 + i - 1) / i
		dp[i] = l * (l - 1) / 2
		for j := 2 * i; j <= n; j += i {
			dp[i] -= dp[j]
		}
	}
	// 假设以x条边,它的cost是x + 1, 相当于单条边的cost是 (x + 1) / x
	// x越大，单条边的cost越小
	var ans int64
	for i := n - 1; i > 1 && m > 0; i-- {
		x := min(dp[i], m)
		y := x / (i - 1)
		ans += int64(y) * int64(i)
		m -= y * (i - 1)
	}

	if m > 0 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
