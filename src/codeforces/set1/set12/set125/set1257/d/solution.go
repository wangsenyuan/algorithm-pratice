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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		m := readNum(reader)
		b := make([][]int, m)
		for i := 0; i < m; i++ {
			b[i] = readNNums(reader, 2)
		}

		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Println(buf.String())
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

func solve(monsters []int, heros [][]int) int {
	// 对heros进行排序
	// 先按照power升序排
	slices.SortFunc(heros, func(a, b []int) int {
		return a[0] - b[0]
	})

	n := len(heros)
	arr := make([][]int, n)
	var top int
	for i := 0; i < n; {
		j := i
		best := heros[j][1]
		for i < n && heros[i][0] == heros[j][0] {
			best = max(best, heros[i][1])
			i++
		}

		for top > 0 && best >= arr[top-1][1] {
			// 当前的更持久且更强力
			top--
		}
		arr[top] = []int{heros[j][0], best}
		top++
	}

	n = top

	m := len(monsters)

	set := NewSegTree(m, 0, max)
	for i := 0; i < m; i++ {
		if monsters[i] > arr[n-1][0] {
			return -1
		}
		set.Update(i, monsters[i])
	}

	check := func(l int, r int) bool {
		v := set.Get(l, r+1)
		i := sort.Search(n, func(i int) bool {
			return arr[i][0] >= v
		})
		s := arr[i][1]
		return r-l+1 <= s
	}

	dp := make([]int, m+1)

	for i, r := m-1, m; i >= 0; i-- {
		for r > i && !check(i, r-1) {
			r--
		}
		dp[i] = dp[r] + 1
	}

	return dp[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
	seg.arr[p] = seg.op(v, seg.arr[p])
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
