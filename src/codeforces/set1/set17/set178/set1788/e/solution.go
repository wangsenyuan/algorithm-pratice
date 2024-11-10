package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	fmt.Println(res)
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

func solve(a []int) int {
	n := len(a)
	pref := make([]int, n+1)
	for i, x := range a {
		pref[i+1] = pref[i] + x
	}

	arr, pos := compact(pref)

	m := len(arr)

	set := NewSegTree(m)

	dp := make([]int, n+1)

	var sum int

	dp[0] = 0

	set.Update(pos[0], 0)

	for i, x := range a {
		dp[i+1] = dp[i]
		sum += x
		j := pos[sum]
		// 找到左边第一个l， sum(l....i) >= 0
		// 0....l-1 sum(j...i) < 0
		v := set.Get(0, j+1)
		dp[i+1] = max(dp[i+1], i+1+v)
		set.Update(j, dp[i+1]-i-1)
	}

	return dp[n]
}

func compact(arr []int) (res []int, pos map[int]int) {
	pos = make(map[int]int)
	for _, x := range arr {
		pos[x]++
	}
	for k := range pos {
		res = append(res, k)
	}
	sort.Ints(res)

	for i := 0; i < len(res); i++ {
		pos[res[i]] = i
	}

	return
}

type SegTree struct {
	arr []int
	sz  int
}

const inf = 1 << 30

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		arr[i] = -inf
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = max(t.arr[p], v)
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := -inf
	l += t.sz
	r += t.sz

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
