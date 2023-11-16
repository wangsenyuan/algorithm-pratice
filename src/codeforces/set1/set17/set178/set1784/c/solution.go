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
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

func solve(a []int) []int {
	n := len(a)

	nodes := make([]Pair, 4*n)
	lazy := make([]int, 4*n)

	pushValue := func(i int, v int) {
		nodes[i] = nodes[i].AddFirst(v)
		lazy[i] += v
	}

	pull := func(i int) {
		nodes[i] = min_pair(nodes[2*i+1], nodes[2*i+2]).AddFirst(lazy[i])
	}

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			nodes[i] = Pair{l, l}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		pull(i)
	}

	var modify func(i int, l int, r int, L int, R int, v int)

	modify = func(i int, l int, r int, L int, R int, v int) {
		if L <= l && r <= R {
			pushValue(i, v)
			return
		}
		mid := (l + r) / 2
		if L <= mid {
			modify(2*i+1, l, mid, L, R, v)
		}
		if mid < R {
			modify(2*i+2, mid+1, r, L, R, v)
		}
		pull(i)
	}

	build(0, 1, n)

	ans := make([]int, n)
	var sum int
	var m int
	for i, x := range a {
		m++
		sum += x
		modify(0, 1, n, x, n, -1)
		if nodes[0].first < 0 {
			p := nodes[0].second
			sum -= p
			m--
			modify(0, 1, n, p, n, 1)
		}
		ans[i] = sum - m*(m+1)/2
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func (p Pair) AddFirst(v int) Pair {
	return Pair{p.first + v, p.second}
}

func min_pair(a, b Pair) Pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}
