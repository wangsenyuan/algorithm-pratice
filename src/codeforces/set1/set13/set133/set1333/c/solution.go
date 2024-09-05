package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(a []int) int {

	pref := make(map[int]bool)
	pref[0] = true
	var res int
	var si, sj int
	for i, j := 0, 0; i < len(a); i++ {
		for j < len(a) && !pref[sj+a[j]] {
			sj += a[j]
			pref[sj] = true
			j++
		}

		res += j - i
		delete(pref, si)
		si += a[i]
	}

	return res
}
func solve1(a []int) int {
	n := len(a)

	var sum int

	// -1, 2, -1, 1
	// [0, 2], [2, 3]
	// [1, 3], [0, 3]
	// 这类重叠的bad区域要怎么处理呢？
	// 1, -1, 1, -1, 1, -1
	// 如果不重叠，[l1, r1] [l2, r2] 那么只要 l1 * (r2 - r1) + l2 * (n - r2)

	pos := make(map[int]int)
	pos[0] = -1

	type pair struct {
		first  int
		second int
	}

	var arr []pair

	for i, num := range a {
		sum += num
		if num == 0 {
			arr = append(arr, pair{i, i + 1})
		} else if j, ok := pos[sum]; ok {
			arr = append(arr, pair{j + 1, i + 1})
		}
		pos[sum] = i
	}

	slices.SortFunc(arr, func(a, b pair) int {
		if a.second != b.second {
			return a.second - b.second
		}
		return b.first - a.first
	})

	// 那些包含一个更小区间的，需要排除掉，它们会造成重复计算

	set := NewSegTree(n, 0, max)

	var rra []pair

	for _, cur := range arr {
		l, r := cur.first, cur.second
		v := set.Get(l, r)
		if v == 0 {
			rra = append(rra, cur)
		}
		set.Update(l, 1)
	}

	arr = rra

	res := n * (n + 1) / 2

	for i := 0; i < len(arr); i++ {
		next := n + 1
		if i+1 < len(arr) {
			next = arr[i+1].second
		}
		res -= (arr[i].first + 1) * (next - arr[i].second)
	}

	return res
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
