package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	var buf bytes.Buffer

	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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
	L, b, f := readThreeNums(reader)
	n := readNum(reader)
	events := make([][]int, n)
	for i := range n {
		events[i] = readNNums(reader, 2)
	}
	return solve(L, b, f, events)
}

const inf = 1 << 60

func solve(L int, b int, f int, events [][]int) []int {

	type car struct {
		id    int
		start int
		end   int
	}
	var cars []car

	park := func(l int) int {
		if l > L {
			return -1
		}
		if len(cars) == 0 || l+f <= cars[0].start {
			return 0
		}
		for i, cur := range cars {
			if i+1 < len(cars) {
				if cars[i+1].start-cur.end >= l+b+f {
					return cur.end + b
				}
			} else {
				if L-cur.end >= l+b {
					return cur.end + b
				}
			}
		}
		return -1
	}

	leave := func(i int) {
		for j := 0; j < len(cars); j++ {
			if cars[j].id == i {
				copy(cars[j:], cars[j+1:])
				cars = cars[:len(cars)-1]
				return
			}
		}
	}

	var res []int
	for i, cur := range events {
		if cur[0] == 1 {
			pos := park(cur[1])
			res = append(res, pos)
			if pos >= 0 {
				cars = append(cars, car{i, pos, pos + cur[1]})
			}
			slices.SortFunc(cars, func(a, b car) int {
				return a.start - b.start
			})
		} else {
			leave(cur[1] - 1)
		}
	}

	return res
}

func solve1(L int, b int, f int, events [][]int) []int {
	t := Build(L)

	var res []int
	var car []int

	s1 := NewSegTree(L, inf, func(a, b int) int {
		return min(a, b)
	})

	s2 := NewSegTree(L, -inf, func(a, b int) int {
		return max(a, b)
	})
	add := func(i int) {
		if i < L {
			s1.Update(i, i)
			s2.Update(i, i)
		}

	}

	rem := func(i int) {
		if i < L {
			s1.Update(i, inf)
			s2.Update(i, -inf)
		}
	}
	t.Update(0, L)

	park := func(l int) {
		car = append(car, l)

		if l > L {
			res = append(res, -1)
			return
		}
		// 先看能否停在第一个位置
		tmp := t.Get(L)
		if tmp.first == 0 {
			// 可以停在第一个位置
			res = append(res, 0)
			// 0后面没有空位了
			t.Update(0, 0)
			t.Update(l, L-l)
			add(0)
			add(l)
			return
		}
		// 还需要检查一次
		tmp = t.Get(l + f)
		if tmp.first == 0 {
			res = append(res, 0)
			t.Update(0, 0)
			t.Update(l, tmp.second-l)
			add(0)
			add(l)
			return
		}
		// 现在处理 b + l + f 的情况
		tmp = t.Get(b + l + f)
		if tmp.first < 0 {
			x := s2.Get(0, L)
			if x < 0 || x+b+l > L {
				res = append(res, -1)
				return
			}
			tmp = pair{x, L - x}
		}
		i := tmp.first
		res = append(res, i+b)
		t.Update(i, b)
		t.Update(i+b+l, tmp.second-(b+l))
		add(i + b)
		add(i + b + l)
	}
	id := make([]int, len(events))

	leave := func(j int) {
		j = id[j-1]
		i := res[j]
		if i == 0 {
			t.Update(i+car[j], 0)
			rem(i)
			rem(i + car[j])
			r := min(L, s1.Get(0, L))
			t.Update(0, r)
			return
		}

		t.Update(i+car[j], 0)
		rem(i)
		rem(i + car[j])
		r := min(L, s1.Get(i, L))
		k := max(0, s2.Get(0, i))
		t.Update(k, r-k)
	}

	for i, cur := range events {
		if cur[0] == 1 {
			id[i] = len(car)
			park(cur[1])
		} else {
			leave(cur[1])
		}
	}

	return res
}

type Tree struct {
	val []int
	sz  int
}

func Build(n int) *Tree {
	val := make([]int, 4*n)
	return &Tree{val, n}
}

func (t *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.val[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
	}
	loop(0, 0, t.sz-1)
}

type pair struct {
	first  int
	second int
}

func (t *Tree) Get(need int) pair {
	if t.val[0] < need {
		return pair{-1, 0}
	}
	var loop func(i int, l int, r int) pair

	loop = func(i int, l int, r int) pair {
		if l == r {
			return pair{l, t.val[i]}
		}
		mid := (l + r) / 2
		if t.val[2*i+1] >= need {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}
	return loop(0, 0, t.sz-1)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
