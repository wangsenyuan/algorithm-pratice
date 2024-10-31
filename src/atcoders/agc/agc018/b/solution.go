package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := make([][]int, n)

	for i := 0; i < n; i++ {
		A[i] = readNNums(reader, m)
	}

	res := solve(A)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}

func solve(A [][]int) int {
	n := len(A)
	m := len(A[0])
	pos := make([][]int, n)
	trs := make([]*SegTree, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, m)
		trs[i] = NewSegTree(m)
		for j := 0; j < m; j++ {
			A[i][j]--
			pos[i][A[i][j]] = j
			trs[i].Update(j, j)
		}
	}

	marked := make([]bool, m)

	cnt := make([]int, m)

	check := func() (int, int) {
		// 找出目前能够得到最的最大值
		clear(cnt)
		for i := 0; i < n; i++ {
			j := trs[i].Get(0, m)
			cnt[A[i][j]]++
		}
		ans := -1
		for i := 0; i < m; i++ {
			if marked[i] {
				continue
			}
			if ans < 0 || cnt[ans] < cnt[i] {
				ans = i
			}
		}
		return ans, cnt[ans]
	}

	best := n

	for j := 0; j < m; j++ {
		v, cur := check()
		best = min(best, cur)
		marked[v] = true
		for i := 0; i < n; i++ {
			trs[i].Update(pos[i][v], inf)
		}
	}

	return best
}

const inf = 1 << 30

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
