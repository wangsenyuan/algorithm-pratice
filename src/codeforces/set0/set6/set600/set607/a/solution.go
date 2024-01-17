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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			a[i], b[i] = readTwoNums(reader)
		}
		res := solve(a, b)

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

func solve(a []int, b []int) int {
	n := len(a)

	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{a[i], b[i]}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first < ps[j].first
	})

	var ans int

	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		j := search(i, func(j int) bool {
			return ps[i].first-ps[i].second <= ps[j].first
		})

		dp[i+1] = dp[j] + 1
		ans = max(ans, dp[i+1])
	}

	return n - ans
}

type Pair struct {
	first  int
	second int
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
