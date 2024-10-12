package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readNNums(reader, n)
	res := solve(s)

	ss := fmt.Sprintf("%v", res)

	fmt.Println(ss[1 : len(ss)-1])
}

func solve(s []int) []int {
	n := len(s)

	arr := make([]pair, n)

	for i := 0; i < n; i++ {
		arr[i] = pair{s[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		if a.first != b.first {
			return b.first - a.first
		}
		return a.second - b.second
	})

	res := make([]int, n)

	for i, cur := range arr {
		res[n-1-i] = cur.second + 1
	}

	return res
}

func solve1(s []int) []int {
	n := len(s)

	p := make([]int, n)

	tr := NewSegTree(s)

	for iter := 0; iter < n; iter++ {
		tmp := tr.Get(0, n-1)
		i := tmp.second
		p[i] = iter + 1
		// clean it
		tr.Update(i, i, inf)
		tr.Update(i, n-1, -(iter + 1))
	}

	return p
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func min_of(l, r pair) pair {
	// 不同值时，返回小的那个，相同值时，返回下标大的那个
	if l.first < r.first || l.first == r.first && l.second > r.second {
		return l
	}
	return r
}

func (this pair) addFirst(v int) pair {
	return pair{this.first + v, this.second}
}

type SegTree struct {
	arr  []pair
	lazy []int
	sz   int
}

func NewSegTree(arr []int) *SegTree {
	n := len(arr)

	res := make([]pair, n*4)

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			res[i] = pair{arr[l], l}
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		res[i] = min_of(res[2*i+1], res[2*i+2])
	}

	build(0, 0, n-1)

	lazy := make([]int, 4*n)

	return &SegTree{res, lazy, n}
}

func (tr *SegTree) update(i int, v int) {
	tr.arr[i] = tr.arr[i].addFirst(v)
	tr.lazy[i] += v
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *SegTree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		loop(2*i+1, l, mid)
		loop(2*i+2, mid+1, r)
		tr.arr[i] = min_of(tr.arr[2*i+1], tr.arr[2*i+2])
	}

	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) pair {
	var loop func(i int, l int, r int) pair
	loop = func(i int, l int, r int) pair {
		if R < l || r < L {
			return pair{inf, -1}
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}

		tr.push(i)

		mid := (l + r) / 2
		a := loop(2*i+1, l, mid)
		b := loop(2*i+2, mid+1, r)
		return min_of(a, b)
	}

	return loop(0, 0, tr.sz-1)
}
