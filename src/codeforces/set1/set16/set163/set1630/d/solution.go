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
		A := readNNums(reader, n)
		B := readNNums(reader, m)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const INF = 1 << 60
const N_INF = -INF

func solve(A []int, B []int) int64 {
	n := len(A)
	//m := len(B)
	// order of B doesn't matter, but sorted might help
	//sort.Ints(B)
	// 如果有存在1， 就可以把A中所有的负数翻转成正数
	// 假设对于（最小的）x，翻转从l开始，那么l和r必须同时翻转 （r - l + 1 == x)
	// 假设A[l] < 0, 但是A[r] > 0, 想要翻转A[l]， 但是不翻转 A[r], 要如何实现呢？
	// 那必须至少从l-1开始翻转？
	// 还是要看gcd， 比如 x = 2, y = 3, 可以先翻转2个，再翻转3个，就可以使得A[l]翻转，而A[r]不翻转
	// 假设 x = gcd(B)
	// 从后面往前开始
	// dp[i][0] = i没有翻转时的最大值
	// dp[i][1] = 从i开始翻转时的最大值
	// dp[i][0] = A[i] + max(dp[i+1][0], dp[i+1][1])
	// dp[i][1] = rev_sum(A[i..i+x]) + max(dp[i+x][0], dp[i+x][1])?
	// 好像还差点东西
	// 如果 从i+1开始翻转
	//  在 i 这里进行翻转 相当于
	// i+1到 i + x - 1退出了翻转,
	// 所以这个 rev_sum的要重新考虑
	// 所以 取最大值 dp[j][1] - (rev_sum(i + x) - rev_sum(j)) + (sum(i + x) - sum(j)) + rev_sum(j) - rev_sum(i)
	// dp[j][1] + rev_sum(j) + rev_sum(j) - sum(j) - rev_sum(i+x) + sum(i+x) - rev_sum(i)
	// 双端队列进行处理
	x := gcd_of_array(B)
	dp := make([][]int64, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int64, 2)
		dp[i][1] = N_INF
	}

	for i := 0; i < n; i++ {
		r := i % x
		v0 := max(dp[r][0]+int64(A[i]), dp[r][1]-int64(A[i]))
		v1 := max(dp[r][0]-int64(A[i]), dp[r][1]+int64(A[i]))
		dp[r][0] = v0
		dp[r][1] = v1
	}
	var s0, s1 int64
	for i := 0; i < x; i++ {
		s0 += dp[i][0]
		s1 += dp[i][1]
	}

	return max(s0, s1)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int64
	second int
}

func gcd_of_array(arr []int) int {
	g := arr[0]
	for i := 1; i < len(arr); i++ {
		g = gcd(g, arr[i])
	}
	return g
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
