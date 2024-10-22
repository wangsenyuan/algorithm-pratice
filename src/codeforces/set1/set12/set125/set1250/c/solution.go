package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	tasks := make([][]int, n)
	for i := 0; i < n; i++ {
		tasks[i] = readNNums(reader, 3)
	}
	profit, L, R, set := solve(k, tasks)

	if profit == 0 {
		fmt.Println(0)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d %d %d\n", profit, L, R, len(set)))
	for _, x := range set {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
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

func solve(k int, tasks [][]int) (profit int, L int, R int, set []int) {
	var m int
	for _, task := range tasks {
		m = max(m, task[1])
	}

	m++
	tr := Build(m)

	at := make([][]pair, m)

	for _, task := range tasks {
		l, r, p := task[0], task[1], task[2]
		at[r] = append(at[r], pair{l, p})
	}

	ans := []int{-1, -1}

	for i := 1; i < m; i++ {
		tr.Update(0, i, -k)

		for _, cur := range at[i] {
			l, p := cur.first, cur.second
			tr.Update(0, l, p)
		}

		tmp := tr.Get(0, i)

		if tmp.first > profit {
			ans = []int{tmp.second, i}
			profit = tmp.first
		}
	}

	if profit == 0 {
		return
	}

	L = ans[0]
	R = ans[1]

	for i, task := range tasks {
		l, r := task[0], task[1]
		if L <= l && r <= R {
			set = append(set, i+1)
		}
	}

	return
}

const inf = 1e18

type pair struct {
	first  int
	second int
}

func (this pair) Max(that pair) pair {
	if this.first > that.first {
		return this
	}
	return that
}

type SegTree struct {
	val  []pair
	lazy []int
	sz   int
}

func Build(n int) *SegTree {
	tr := new(SegTree)
	tr.val = make([]pair, 4*n)
	tr.lazy = make([]int, 4*n)
	tr.sz = n

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = pair{0, l}
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr.pull(i)
	}

	build(0, 0, n-1)

	return tr
}

func (tr *SegTree) update(i int, v int) {
	tr.val[i].first += v
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(i*2+1, tr.lazy[i])
		tr.update(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) pull(i int) {
	tr.val[i] = tr.val[i*2+1].Max(tr.val[i*2+2])
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		tr.pull(i)
	}
	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) pair {
	var loop func(i int, l int, r int) pair
	loop = func(i int, l int, r int) pair {
		if r < L || R < l {
			return pair{-inf, -1}
		}
		if L <= l && r <= R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		a := loop(i*2+1, l, mid)
		b := loop(i*2+2, mid+1, r)
		return a.Max(b)
	}

	return loop(0, 0, tr.sz-1)
}
