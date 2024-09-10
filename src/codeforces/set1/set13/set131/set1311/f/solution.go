package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	res := solve(a, b)

	fmt.Println(res)
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

type pair struct {
	first  int
	second int
}

func solve(x []int, v []int) int {
	// v[i] >= 0
	n := len(x)
	w := make([]int, n)
	copy(w, v)
	sort.Ints(w)

	find := func(speed int) int {
		return sort.SearchInts(w, speed)
	}

	players := make([]pair, n)
	for i := 0; i < n; i++ {
		players[i] = pair{x[i], find(v[i])}
	}

	sum := NewTree(n, 0, add)
	cnt := NewTree(n, 0, add)

	slices.SortFunc(players, func(a, b pair) int {
		return a.first - b.first
	})

	var res int

	for _, p := range players {
		// i 有可能会重复（相同速度的，有相同的i)
		i := p.second
		c := cnt.Find(0, i+1)
		s := sum.Find(0, i+1)
		res += c*p.first - s
		cnt.Add(i, 1)
		sum.Add(i, p.first)
	}

	return res
}

func add(a, b int) int {
	return a + b
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, f}
}

func (tree *SegTree) Add(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] += v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
