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

	res, _, _ := process(reader)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res []int, n int, segs [][]int) {
	n, m := readTwoNums(reader)
	segs = make([][]int, m)
	for i := range m {
		segs[i] = readNNums(reader, 3)
	}
	res = solve(n, segs)
	return
}

func solve(n int, segs [][]int) []int {
	pos := make(map[int]int)
	for _, seg := range segs {
		pos[seg[1]]++
		pos[seg[2]]++
	}
	arr := make([]int, 0, len(pos))
	for x := range pos {
		arr = append(arr, x)
	}
	sort.Ints(arr)
	for i, x := range arr {
		pos[x] = i
	}

	rows := make([][]pair, n)
	for _, seg := range segs {
		i, l, r := seg[0], seg[1], seg[2]
		rows[i-1] = append(rows[i-1], pair{pos[l], pos[r]})
	}
	m := len(arr)

	tr := NewSegTree(m)

	dp := make([]int, n)
	from := make([]int, n)

	for i := range n {
		var best pair
		for _, seg := range rows[i] {
			l, r := seg.first, seg.second
			tmp := tr.Get(l, r)
			best = max_of(best, tmp)
		}
		if best.first == 0 {
			// 没有和它重叠的
			dp[i] = 1
			from[i] = -1
		} else {
			dp[i] = best.first + 1
			from[i] = best.second
		}
		for _, seg := range rows[i] {
			l, r := seg.first, seg.second
			tr.Update(l, r, pair{dp[i], i})
		}
	}

	var keep []int
	it := tr.Get(0, m-1).second
	for it >= 0 {
		keep = append(keep, it)
		it = from[it]
	}

	reverse(keep)

	var ans []int

	for i, j := 0, 0; i < n; i++ {
		for j < len(keep) && keep[j] < i {
			j++
		}
		if j < len(keep) && keep[j] == i {
			j++
			continue
		}
		ans = append(ans, i+1)
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type pair struct {
	first  int
	second int
}

func max_of(a pair, b pair) pair {
	if a.first >= b.first {
		return a
	}
	return b
}

var zero = pair{0, 0}

type SegTree struct {
	arr  []pair
	lazy []pair
	sz   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]pair, 4*n)
	lazy := make([]pair, 4*n)
	for i := 0; i < 4*n; i++ {
		arr[i] = zero
		lazy[i] = zero
	}

	return &SegTree{arr, lazy, n}
}

func (tr *SegTree) update(i int, v pair) {
	tr.arr[i] = max_of(tr.arr[i], v)
	tr.lazy[i] = max_of(tr.lazy[i], v)
}

func (tr *SegTree) push(i int) {
	if tr.lazy[i] != zero {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = zero
	}
}

func (tr *SegTree) pull(i int) {
	tr.arr[i] = max_of(tr.arr[2*i+1], tr.arr[2*i+2])
}

func (tr *SegTree) Update(L int, R int, v pair) {
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
		tr.pull(i)
	}

	loop(0, 0, tr.sz-1)
}

func (tr *SegTree) Get(L int, R int) pair {
	var loop func(i int, l int, r int) pair
	loop = func(i int, l int, r int) pair {
		if R < l || r < L {
			return zero
		}
		if L <= l && r <= R {
			return tr.arr[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		a := loop(2*i+1, l, mid)
		b := loop(2*i+2, mid+1, r)
		return max_of(a, b)
	}
	return loop(0, 0, tr.sz-1)
}
