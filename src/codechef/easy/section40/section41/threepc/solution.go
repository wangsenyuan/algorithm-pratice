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

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		C := readNNums(reader, n)
		res := solve(A, B, C)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(A, B, C []int) []int {
	a := find(A, B, C)
	b := find(B, A, C)
	c := find(C, A, B)
	return []int{a, b, c}
}

func find(A, B, C []int) int {
	n := len(A)
	X := make([]Pair, n+1)
	Y := make([]Pair, n+1)

	for i := 1; i <= n; i++ {
		X[i] = Pair{X[i-1].first + int64(A[i-1]-B[i-1]), i}
		Y[i] = Pair{Y[i-1].first + int64(A[i-1]-C[i-1]), i}
	}

	sort.Slice(Y, func(i, j int) bool {
		return Y[i].first < Y[j].first || Y[i].first == Y[j].first && Y[i].second < Y[j].second
	})

	rk := make([]int, n+1)

	for i, y := range Y {
		rk[y.second] = i
	}

	m := n + 1

	tree := make([]int, 2*m)
	for i := 0; i < 2*m; i++ {
		tree[i] = m * 2
	}

	set := func(p int, v int) {
		p += m
		tree[p] = min(tree[p], v)
		for p > 1 {
			tree[p>>1] = min(tree[p], tree[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += m
		r += m
		res := m * 2

		for l < r {
			if l&1 == 1 {
				res = min(res, tree[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, tree[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	sort.Slice(X, func(i, j int) bool {
		return X[i].first < X[j].first || X[i].first == X[j].first && X[i].second < X[j].second
	})

	var best int

	// X已经按照first，升序排列，所以X[r].first >= X[..].first for all before r
	//
	for _, v := range X {
		r := v.second
		// i是y的位置，zai i之前出现的
		l := get(0, rk[r])
		best = max(best, r-l)
		set(rk[r], r)
	}

	return best
}

type Pair struct {
	first  int64
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
