package main

import (
	"bufio"
	"bytes"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func process(reader *bufio.Reader) []bool {
	n, q := readTwoNums(reader)
	a := readNNums(reader, n)
	qs := make([][]int, q)
	for i := range q {
		qs[i] = readNNums(reader, 2)
	}
	return solve(a, qs)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, queries [][]int) []bool {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})

	q := make([][]pair, n+1)
	for i, cur := range queries {
		j, x := cur[0]-1, cur[1]
		q[x] = append(q[x], pair{j, i})
	}
	for i := 1; i <= n; i++ {
		slices.SortFunc(q[i], func(x, y pair) int {
			return x.first - y.first
		})
	}

	ans := make([]bool, len(queries))
	cur := make([]int, n+1)

	active := NewTree(n)
	for i := 0; i < n; i++ {
		active.Update(i, 1)
	}
	var p int
	for lvl := 1; lvl <= n; lvl++ {
		for k := 1; k <= n; k++ {
			if cur[k] >= n {
				break
			}
			x := active.OrderOfKey(cur[k])
			var nxt int
			if x+k-1 >= active.arr[0] {
				nxt = n
			} else {
				nxt = active.FindByOrder(x + k)
			}
			for len(q[k]) > 0 && q[k][0].first <= nxt {
				w := q[k][0]
				ans[w.second] = a[w.first] >= lvl
				q[k] = q[k][1:]
			}
			cur[k] = nxt + 1
		}
		for p < n && arr[p].first == lvl {
			active.Update(arr[p].second, 0)
			p++
		}
	}

	return ans
}

type Tree struct {
	arr []int
	sz  int
}

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	return &Tree{arr, n}
}

func (tr *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.arr[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		tr.arr[i] = tr.arr[2*i+1] + tr.arr[2*i+2]
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) OrderOfKey(k int) int {
	if k == 0 {
		return 0
	}
	k--
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return tr.arr[i]
		}
		mid := (l + r) / 2
		if k <= mid {
			return loop(2*i+1, l, mid)
		}
		return tr.arr[2*i+1] + loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, tr.sz-1)
}

func (tr *Tree) FindByOrder(ord int) int {
	var loop func(i int, l int, r int, ord int) int
	loop = func(i int, l int, r int, ord int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if tr.arr[2*i+1] >= ord {
			return loop(2*i+1, l, mid, ord)
		}
		return loop(2*i+2, mid+1, r, ord-tr.arr[2*i+1])
	}
	return loop(0, 0, tr.sz-1, ord)
}
