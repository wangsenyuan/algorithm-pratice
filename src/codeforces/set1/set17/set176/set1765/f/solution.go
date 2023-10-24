package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	x := make([]int, n)
	w := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		x[i], w[i], c[i] = readThreeNums(reader)
	}
	res := solve(k, x, w, c)

	fmt.Printf("%.8f\n", res)
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

func solve(k int, x []int, w []int, c []int) float64 {
	n := len(x)
	a := make([]Base, n)
	for i := 0; i < n; i++ {
		a[i] = Base{x[i], w[i], c[i]}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].x > a[j].x
	})

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -inf
	}
	var ans int
	for i := 0; i < n; i++ {
		dp[i] = max(dp[i], -a[i].w*200)
		for j := i + 1; j < n; j++ {
			dp[j] = max(dp[j], dp[i]+a[i].Area(a[j])*k-a[j].w*200)
		}
		ans = max(ans, dp[i])
	}

	return float64(ans) / 200
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Base struct {
	x int
	w int
	c int
}

func (a Base) Area(b Base) int {
	return (a.c + b.c) * abs(a.x-b.x)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
