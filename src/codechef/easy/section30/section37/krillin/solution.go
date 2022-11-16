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
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 3)
		}
		res := solve(A, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int, Q [][]int) []int64 {
	n := len(A)

	nodes := make([]Node, 2*n)

	for i := n; i < 2*n; i++ {
		nodes[i] = Node{i - n + 1, A[i-n], int64(A[i-n])}
	}

	for i := n - 1; i > 0; i-- {
		nodes[i] = merge(nodes[i*2], nodes[i*2+1])
	}

	set := func(p int, v int) {
		p += n
		nodes[p] = Node{p - n + 1, v, int64(v)}

		for p > 1 {
			nodes[p>>1] = merge(nodes[p], nodes[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) Node {
		l += n
		r += n
		var res Node

		for l < r {
			if l&1 == 1 {
				res = merge(res, nodes[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = merge(res, nodes[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	query := func(l int, r int) int64 {
		x := get(l, r)
		mode := x.pos
		sum := x.sum

		sum = sum/2 + sum&1

		// binary search i, [l...r] >= sum
		median := search(l, r, func(i int) bool {
			return get(l, i).sum >= sum
		})
		return lcm(int64(median), int64(mode))
	}

	ans := make([]int64, 0, len(Q))

	for _, cur := range Q {
		if cur[0] == 1 {
			ans = append(ans, query(cur[1]-1, cur[2]))
		} else {
			p, v := cur[1], cur[2]
			set(p-1, v)
			A[p-1] = v
		}
	}

	return ans
}

func search(l int, r int, f func(int) bool) int {
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

func lcm(a, b int64) int64 {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Node struct {
	pos  int
	freq int
	sum  int64
}

func merge(a, b Node) Node {
	if a.freq > b.freq {
		return Node{a.pos, a.freq, a.sum + b.sum}
	}
	return Node{b.pos, b.freq, a.sum + b.sum}
}
