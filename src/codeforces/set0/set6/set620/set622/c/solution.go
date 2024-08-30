package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}
	res := solve(a, queries)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(a []int, queries [][]int) []int {
	n := len(a)

	left := make([]int, n)

	for i := 0; i < n; i++ {
		left[i] = i - 1
		if i > 0 && a[i-1] == a[i] {
			left[i] = left[i-1]
		}
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r, x := cur[0], cur[1], cur[2]
		l--
		r--
		if a[r] != x {
			ans[i] = r + 1
		} else if left[r] >= l {
			ans[i] = left[r] + 1
		} else {
			ans[i] = -1
		}
	}

	return ans
}
func solve1(a []int, queries [][]int) []int {

	ma := slices.Max(a)

	pos := make([][]int, ma+1)

	for i, x := range a {
		pos[x] = append(pos[x], i)
	}

	n := len(a)

	id := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]][2] < queries[id[j]][2]
	})

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = queries[i][0]
	}
	tr := NewSegTree(n, -2, max)

	for i := 0; i < n; i++ {
		tr.Update(i, i)
	}

	for i := 0; i < len(id); {
		x := queries[id[i]][2]

		if x > ma {
			break
		}

		for _, j := range pos[x] {
			tr.Update(j, -2)
		}

		for i < len(id) && queries[id[i]][2] == x {
			l, r := queries[id[i]][0], queries[id[i]][1]
			v := tr.Get(l-1, r, -2)
			ans[id[i]] = v + 1
			i++
		}

		for _, j := range pos[x] {
			tr.Update(j, j)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type SegTree struct {
	arr []int
	sz  int
	f   func(int, int) int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{arr, n, f}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int, iv int) int {
	res := iv
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
