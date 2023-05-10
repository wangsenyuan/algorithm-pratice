package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s := readString(reader)
		n := readNum(reader)
		F := make([]int, n)
		for i := 0; i < n; i++ {
			F[i] = readNum(reader)
		}
		res := solve(s, F)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(s string, F []int) []int64 {
	L, P := suffixArray(s)
	n := len(s)
	A := make([]int, n-1)
	for i := 0; i+1 < n; i++ {
		A[i] = findLcp(P, n, L[i].id, L[i+1].id)
	}
	var mn []int
	mx_seg := NewSegTree(n, -1, max)
	for i := 0; i < n-1; i++ {
		v := mx_seg.Query(0, A[i])
		mn = append(mn, v)
		mx_seg.Update(A[i], i)
	}
	var mx []int
	mn_seg := NewSegTree(n, n-1, min)
	for i := n - 2; i >= 0; i-- {
		v := mn_seg.Query(0, A[i])
		mx = append(mx, v)
		mn_seg.Update(A[i], i)
	}
	reverse(mx)

	indices := make([][]int, n)
	for i := 0; i < len(A); i++ {
		id := A[i]
		if len(indices[id]) == 0 {
			indices[id] = make([]int, 0, 1)
		}
		indices[id] = append(indices[id], i)
	}

	D := make([]int64, n+1)

	for i := 1; i < n; i++ {
		var right int
		for j := 0; j < len(indices[i]); j++ {
			id := indices[i][j]
			if id >= right {
				lo, hi := mn[id], mx[id]
				t := hi - lo
				x := i
				if 0 <= hi && hi < len(A) {
					x = min(x, i-A[hi])
				}
				if lo >= 0 && lo < len(A) {
					x = min(x, i-A[lo])
				}
				D[t] += int64(t) * int64(x)
				right = hi
			}
		}
	}

	var tot int64
	for i := 2; i < len(D); i++ {
		tot += D[i]
	}

	D[1] = int64(n)*int64(n+1)/2 - tot

	totSum := D[1] + tot

	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + D[i]
	}

	res := make([]int64, len(F))
	for i, f := range F {
		if f > n {
			res[i] = 0
			continue
		}
		res[i] = totSum - sum[f-1]
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type SegTree struct {
	f   func(int, int) int
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{f, arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = t.f(t.arr[p], v)

	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
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

func min(a, b int) int {
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

type Entry struct {
	nr [2]int
	id int
}

func findLcp(P [][]int, n int, x int, y int) int {
	if x == y {
		return n - x
	}
	h := len(P)
	var res int
	for stp := h - 1; stp >= 0 && x < n && y < n; stp-- {
		if P[stp][x] == P[stp][y] {
			x += 1 << stp
			y += 1 << stp
			res += 1 << stp
		}
	}
	return res
}

func suffixArray(s string) ([]*Entry, [][]int) {
	n := len(s)
	var h int
	for 1<<h < n {
		h++
	}
	h++
	// 1 << h > n
	P := make([][]int, h)

	for i := 0; i < h; i++ {
		P[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		P[0][i] = int(s[i] - 'a')
	}

	L := make([]*Entry, n)

	for i := 0; i < n; i++ {
		L[i] = new(Entry)
	}

	cmp := func(i, j int) bool {
		return L[i].nr[0] < L[j].nr[0] || L[i].nr[0] == L[j].nr[0] && L[i].nr[1] < L[j].nr[1]
	}

	for stp := 1; stp < h; stp++ {
		for i := 0; i < n; i++ {
			L[i].nr[0] = P[stp-1][i]
			L[i].nr[1] = -1
			if i+(1<<(stp-1)) < n {
				L[i].nr[1] = P[stp-1][i+(1<<(stp-1))]
			}
			L[i].id = i
		}
		sort.Slice(L, cmp)
		for i := 0; i < n; i++ {
			if i > 0 && L[i].nr == L[i-1].nr {
				P[stp][L[i].id] = P[stp][L[i-1].id]
			} else {
				P[stp][L[i].id] = i
			}
		}
	}
	return L, P
}
