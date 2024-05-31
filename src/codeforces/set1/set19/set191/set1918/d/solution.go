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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)

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

func solve(a []int) int {
	n := len(a)

	sum := make([]int, n+1)
	for i, num := range a {
		sum[i+1] = sum[i] + num
	}

	que := make([]Pair, n+1)

	// dp[i] = 前i个元素在符合条件的情况下，最小的blocked的sum
	dp := make([]int, n+1)

	check := func(x int) bool {
		for i := 0; i <= n; i++ {
			dp[i] = inf
		}
		dp[0] = 0
		var front, tail int
		// que[front] = Pair{0, 0}

		front = 1

		for i := 1; i <= n; i++ {
			// 如果选择block i
			dp[i] = dp[i-1] + a[i-1]
			// 或者不block i
			for tail < front && sum[i]-sum[que[tail].second] > x {
				tail++
			}
			if tail < front {
				dp[i] = min(dp[i], que[tail].first)
			}
			val := dp[i-1] + a[i-1]
			for tail < front && que[front-1].first >= val {
				front--
			}
			que[front] = Pair{val, i}
			front++
		}
		return dp[n] <= x
	}

	l, r := 0, sum[n]

	for l < r {
		mid := (l + r) / 2

		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

type Pair struct {
	first  int
	second int
}
