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
		n, s := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, s)
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

const inf = 1 << 60

func solve(a []int, s int) int {
	// F = a1 * x2 + y2 * x3 + y3 * x4 ... + yn-1 * an
	// 要求 (xi - s) * (yi - s) >= 0
	//  1, xi or yi = 0
	//  2  xi > s and yi > s
	//  3 xi < s and yi < s
	// dp[i][0] 表示 x[i] = a[i],
	// dp[i][1] 表示 x[i] = 0
	n := len(a)
	// 如果 s >= a[i], 那么 x,y 的取值肯定是 (0, a[i])
	// 如果 s < a[i] 那么 x, y的取值肯定是 (s, a[i] - s)
	// dp[i][0] 表示 x取小的数 min(s, a[i] - s) or min(0, a[i]))
	// dp[i][1] 表示 x取大的数
	// dp[1][0] = a[0] * min(...)
	// dp[i][1] = a[0] * max(...)
	dp := make([]int, 2)

	getValue := func(i int) []int {
		if s >= a[i] {
			return []int{0, a[i]}
		}
		return []int{min(s, a[i]-s), max(s, a[i]-s)}
	}

	prev := getValue(1)
	dp[0] = a[0] * prev[0]
	dp[1] = a[0] * prev[1]

	for i := 1; i+2 < n; i++ {
		cur := getValue(i + 1)
		fp := []int{min(cur[0]*prev[0]+dp[1], cur[0]*prev[1]+dp[0]), min(cur[1]*prev[0]+dp[1], cur[1]*prev[1]+dp[0])}
		copy(dp, fp)
		prev = cur
	}

	return min(dp[0]+prev[1]*a[n-1], dp[1]+prev[0]*a[n-1])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
