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
	for range tc {
		res := process(reader)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

const inf = 1 << 60

func solve(a []int, b []int) []int {
	n := len(a)
	m := len(b)

	dp := make([]*Tree, m)
	for i := range m {
		dp[i] = NewTree(n + 1)
		dp[i].Update(n, 1, 0)
	}

	next := make([][]int, m)
	for i := range m {
		next[i] = make([]int, n)
		var r int
		var sum int
		for j := range n {
			for r < n && sum+a[r] <= b[i] {
				sum += a[r]
				r++
			}
			next[i][j] = r
			sum -= a[j]
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i+1 <= next[j][i] {
				tmp := dp[j].Get(i+1, next[j][i])
				dp[j].Update(i, tmp.cnt, tmp.val+m-(j+1))
			}
			if j != m-1 {
				tmp1 := dp[j+1].Get(i, i)
				tmp2 := dp[j].Get(i, i)
				if tmp1.val < tmp2.val {
					dp[j].Update(i, tmp1.cnt, tmp1.val)
				} else if tmp1.val == tmp2.val {
					dp[j].Update(i, add(tmp1.cnt, tmp2.cnt), tmp1.val)
				}
			}
		}
	}
	res := dp[0].Get(0, 0)
	if res.val == inf {
		return []int{-1}
	}
	return []int{res.val, res.cnt}
}

type state struct {
	cnt int
	val int
}

func merge(a state, b state) state {
	if a.val < b.val {
		return a
	}
	if a.val > b.val {
		return b
	}
	return state{add(a.cnt, b.cnt), a.val}
}

type Tree struct {
	states []state
	sz     int
}

func NewTree(n int) *Tree {
	tree := new(Tree)
	tree.states = make([]state, 4*n)
	tree.sz = n
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tree.states[i] = state{1, inf}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		tree.states[i] = merge(tree.states[2*i+1], tree.states[2*i+2])
	}
	build(0, 0, n-1)
	return tree
}

func (tr *Tree) Update(p int, c int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.states[i] = state{c, v}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.states[i] = merge(tr.states[2*i+1], tr.states[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr Tree) Get(L int, R int) state {
	var loop func(i int, l int, r int) state
	loop = func(i int, l int, r int) state {
		if R < l || r < L {
			return state{0, inf}
		}
		if L <= l && r <= R {
			return tr.states[i]
		}
		mid := (l + r) / 2
		a := loop(2*i+1, l, mid)
		b := loop(2*i+2, mid+1, r)
		return merge(a, b)
	}
	return loop(0, 0, tr.sz-1)
}
