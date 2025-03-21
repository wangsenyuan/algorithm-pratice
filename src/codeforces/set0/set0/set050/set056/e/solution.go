package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := process(reader)
	s := fmt.Sprintf("%v", result)
	fmt.Println(s[1 : len(s)-1])
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
	n := readNum(reader)
	cards := make([][]int, n)

	for i := 0; i < n; i++ {
		cards[i] = readNNums(reader, 2)
	}

	return solve(cards)
}

type item struct {
	id     int
	pos    int
	height int
}

func getAndSort(cards [][]int) []item {
	n := len(cards)
	items := make([]item, n)
	for i := 0; i < n; i++ {
		items[i] = item{i, cards[i][0], cards[i][1]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].pos < items[j].pos
	})
	return items
}

func solve(cards [][]int) []int {
	items := getAndSort(cards)
	n := len(cards)

	st := NewSegTree(n)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		it := items[i]
		j := sort.Search(n, func(j int) bool {
			return items[j].pos >= it.pos+it.height
		})
		u := st.Query(i+1, j)
		u = max(u, i+1)
		ans[it.id] = u - i
		st.Update(i, u)
	}
	return ans
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (st *SegTree) Update(p, v int) {
	p += st.sz
	st.arr[p] = v
	for p > 1 {
		st.arr[p>>1] = max(st.arr[p], st.arr[p^1])
		p >>= 1
	}
}

func (st *SegTree) Query(l, r int) int {
	l += st.sz
	r += st.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, st.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, st.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
