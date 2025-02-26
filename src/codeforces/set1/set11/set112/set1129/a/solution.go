package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
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
	n, m := readTwoNums(reader)
	candies := make([][]int, m)
	for i := range m {
		candies[i] = readNNums(reader, 2)
	}
	return solve(n, candies)
}

const inf = 1 << 30

func solve(n int, candies [][]int) []int {
	at := make([]int, n)

	for i := range n {
		at[i] = inf
	}

	cnt := make([]int, n)

	getDist := func(i int, j int) int {
		if i < j {
			return j - i
		}
		return j + n - i
	}

	for _, cur := range candies {
		a, b := cur[0], cur[1]
		a--
		b--
		// 最后一圈的时间
		at[a] = min(at[a], getDist(a, b))
		cnt[a]++
	}

	x := slices.Max(cnt)
	s1 := NewSegTree(2 * n)
	s2 := NewSegTree(2 * n)

	for i := range n {
		if cnt[i] == 0 {
			continue
		}
		if cnt[i] == x-1 {
			s1.Update(i, i+at[i])
			s1.Update(i+n, i+n+at[i])
		} else if cnt[i] == x {
			s2.Update(i, i+at[i])
			s2.Update(i+n, i+n+at[i])
		}
	}

	res := make([]int, n)

	for i := range n {
		res[i] = (x - 1) * n
		u := s2.Get(i, i+n)
		v := s1.Get(i, i+n)
		res[i] += max(u-i, v-(i+n))
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
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
