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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(A, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(A []int, E [][]int) int64 {
	// 假设 E是空的，如何找到最大值呢?
	// (cnt[x] + cnt[y]) * (x + y) (0)
	// cnt[x] * x + cnt[y] * y + cnt[x] * y + cnt[y] * x
	// 给定x的情况下，如何找到y使得上面的结果最大呢？
	// 好像很难搞哎。
	// 如果在 x, cnt[x], cnt[y]都确定的情况下，
	// 找到最大的y即可
	// x, cnt[x]是确定的，cnt[y] <= cnt[x]， 似乎是有限的几个
	// 最多有 Logn 不同的cnt[x]
	adj := make(map[int][]int)

	addEdge := func(u int, v int) {
		if len(adj[u]) == 0 {
			adj[u] = make([]int, 0, 1)
		}
		adj[u] = append(adj[u], v)
	}

	cnt := make(map[int]int)

	for _, x := range A {
		cnt[x]++
	}

	freqs := make([]Num, 0, len(cnt))
	for k, v := range cnt {
		freqs = append(freqs, Num{k, v})
	}

	sort.Sort(Nums(freqs))
	pos := make(map[int]int)
	ids := make(map[int]int)

	for _, cur := range freqs {
		pos[cur.val] = ids[cur.freq]
		ids[cur.freq]++
	}

	for _, e := range E {
		u, v := e[0], e[1]
		if cnt[u] == cnt[v] {
			if pos[u] < pos[v] {
				u, v = v, u
			}
		} else if cnt[u] < cnt[v] {
			u, v = v, u
		}
		// cnt[u]
		addEdge(u, v)
	}

	trs := make(map[int]*Tree)

	var best int64

	for _, cur := range freqs {
		x := cur.val
		cx := cur.freq
		// process bad
		for _, y := range adj[x] {
			cy := cnt[y]
			// some num later
			//i is the position in array of cy
			trs[cy].Set(pos[y], 0)
		}

		for cy, t := range trs {
			// (cx + cy) * (x + y) 最大
			y := t.Get(0, ids[cy])
			if y > 0 {
				best = max2(best, int64(cx+cy)*int64(x+y))
			}
		}

		// restore
		for _, y := range adj[x] {
			cy := cnt[y]
			//i is the position in array of cy
			trs[cy].Set(pos[y], y)
		}

		if trs[cx] == nil {
			trs[cx] = NewTree(ids[cx])
		}
		trs[cx].Set(pos[x], x)
	}

	return best
}

type Tree struct {
	arr []int
	sz  int
}

func NewTree(n int) *Tree {
	t := new(Tree)
	t.arr = make([]int, 2*n)
	t.sz = n
	return t
}

func (t *Tree) Set(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *Tree) Get(l int, r int) int {
	l += t.sz
	r += t.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type Num struct {
	val  int
	freq int
}

func (this Num) Less(that Num) bool {
	return this.freq < that.freq || this.freq == that.freq && this.val > that.val
}

type Nums []Num

func (a Nums) Len() int {
	return len(a)
}

func (a Nums) Less(i, j int) bool {
	return a[i].Less(a[j])
}

func (a Nums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
