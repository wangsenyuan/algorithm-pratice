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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		cold := readNNums(reader, m)
		hot := readNNums(reader, m)
		res := solve(a, cold, hot)
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

func solve(a []int, cold []int, hot []int) int {
	n := len(a)
	k := len(cold)

	for i := 0; i < n; i++ {
		a[i]--
	}
	a = append(a, k)
	hot = append(hot, 0)
	cold = append(cold, 0)

	ans := cold[a[0]]
	for i := 1; i < n; i++ {
		if a[i-1] == a[i] {
			ans += hot[a[i]]
		} else {
			ans += cold[a[i]]
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = inf
	}
	from := make([]int, k+1)
	for i := 0; i <= k; i++ {
		from[i] = inf
	}
	var best int

	for i := 1; i <= n; i++ {
		dp[i] = min(best, from[a[i]]+hot[a[i]]-cold[a[i]])
		if a[i] == a[i-1] {
			dp[i] += cold[a[i]] - hot[a[i]]
		}
		best = min(best, dp[i])
		from[a[i-1]] = min(from[a[i-1]], dp[i])
	}

	ans += dp[n]

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
