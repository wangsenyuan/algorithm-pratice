package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, x := readTwoNums(reader)
		a := readNNums(reader, n)
		c := readNNums(reader, n)
		res := solve(x, a, c)
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

func solve(m int, a []int, c []int) int {
	type pair struct {
		first  int
		second int
	}

	n := len(a)
	flowers := make([]pair, n)
	for i := 0; i < n; i++ {
		flowers[i] = pair{a[i], c[i]}
	}

	slices.SortFunc(flowers, func(a, b pair) int {
		return a.first - b.first
	})

	prev := pair{-1, 0}

	calc := func(a, b pair) int {
		x := a.first
		k1 := min(a.second, m/x)
		cnt := m - k1*x
		k2 := min(b.second, cnt/(x+1))
		cnt = m - k1*x - k2*(x+1)
		r := min(k1, b.second-k2, cnt)
		return (k1-r)*x + (k2+r)*(x+1)
	}

	var best int

	for i := 0; i < n; {
		j := i
		var sum int
		for i < n && flowers[i].first == flowers[j].first {
			sum += flowers[i].second
			i++
		}

		best = max(best, min(m/flowers[j].first, sum)*flowers[j].first)

		cur := pair{flowers[j].first, sum}

		if prev.first+1 == flowers[j].first {
			// 可以组成
			// let x = prev.first
			// a * x + b * (x + 1) <= m
			// a <= prev.second, b <= sum
			best = max(best, calc(prev, cur))
		}

		prev = cur
	}

	return best
}
