package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	fs := make([][]int, n)
	for i := 0; i < n; i++ {
		fs[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 3)
	}
	res := solve(A, fs, Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const block = 200

const SZ = 10101010

func solve(A []int, fns [][]int, Q [][]int) []int64 {
	n := len(A)
	roots := make([]int, n+1)
	roots[0] = 0

	tr := new(Tree)
	tr.val = make([]int64, SZ)
	tr.left = make([]int, SZ)
	tr.right = make([]int, SZ)
	tr.sz = 0
	tr.n = n

	for i := 1; i <= n; i++ {
		roots[i] = tr.Update(roots[i-1], fns[i-1][0], 1)
		if fns[i-1][1] < n {
			roots[i] = tr.Update(roots[i], fns[i-1][1]+1, -1)
		}
	}

	acc := make([]int64, n+1)
	fv := make([]int64, n+1)

	build := func() {
		for i := 1; i <= n; i++ {
			acc[i] = acc[i-1] + int64(A[i-1])
		}

		for i := 1; i <= n; i++ {
			fv[i] = fv[i-1] + acc[fns[i-1][1]] - acc[fns[i-1][0]-1]
		}
	}

	seen := make([]bool, n+1)

	var res []int64

	upd := make([]Pair, block+1)

	for id := 0; id < len(Q); {
		build()
		var it int
		for k := 0; k < block && id < len(Q); k, id = k+1, id+1 {
			x, y := Q[id][1], Q[id][2]
			if Q[id][0] == 1 {
				upd[it] = Pair{x, y}
				it++
				continue
			}
			ans := fv[y] - fv[x-1]
			for i := it - 1; i >= 0; i-- {
				cur := upd[i]
				u, h := cur.first, cur.second
				if seen[u] {
					continue
				}
				seen[u] = true
				tmp := tr.Query(roots[y], u) - tr.Query(roots[x-1], u)
				ans -= tmp * int64(A[u-1])
				ans += tmp * int64(h)
			}

			for i := it - 1; i >= 0; i-- {
				seen[upd[i].first] = false
			}
			res = append(res, ans)
		}

		for i := 0; i < it; i++ {
			p := upd[i]
			A[p.first-1] = p.second
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Tree struct {
	val   []int64
	left  []int
	right []int
	sz    int
	n     int
}

func (t *Tree) nextNode() int {
	t.sz++
	return t.sz
}

func (t *Tree) copy(src int) int {
	dst := t.nextNode()
	t.val[dst] = t.val[src]
	t.left[dst] = t.left[src]
	t.right[dst] = t.right[src]
	return dst
}

func (t *Tree) Update(root int, pos int, v int) int {

	var loop func(i int, l int, r int) int

	loop = func(i int, l int, r int) int {
		j := t.copy(i)
		if l == r {
			t.val[j] += int64(v)
			return j
		}
		mid := (l + r) / 2
		if pos <= mid {
			t.left[j] = loop(t.left[j], l, mid)
		} else {
			t.right[j] = loop(t.right[j], mid+1, r)
		}
		t.val[j] = t.val[t.left[j]] + t.val[t.right[j]]
		return j
	}

	return loop(root, 1, t.n)
}

func (t *Tree) Query(root int, pos int) int64 {
	var loop func(i int, l int, r int) int64

	loop = func(i int, l int, r int) int64 {
		if l > pos {
			return 0
		}
		if r <= pos {
			return t.val[i]
		}
		mid := (l + r) / 2

		return loop(t.left[i], l, mid) + loop(t.right[i], mid+1, r)
	}

	return loop(root, 1, t.n)
}
