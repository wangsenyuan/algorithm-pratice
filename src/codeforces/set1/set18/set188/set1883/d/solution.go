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

	n := readNum(reader)
	ops := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var x int
		if s[0] == '-' {
			x = 1
		}
		var l, r, pos int
		pos = readInt(s, 2, &l)
		readInt(s, pos+1, &r)
		ops[i] = []int{x, l, r}
	}

	res := solve(ops)

	var buf bytes.Buffer

	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const oo = 1 << 30

func solve(ops [][]int) []bool {
	nums := make(map[int]int)
	for _, op := range ops {
		nums[op[1]]++
		nums[op[2]]++
	}
	arr := make([]int, 0, len(nums))
	for k := range nums {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	n := len(arr)
	for i := 0; i < n; i++ {
		nums[arr[i]] = i
	}
	// 需要直到最左边的右端点，和最右边的左端点，
	// 只要它们不相交，就可以了
	ans := make([]bool, len(ops))

	tr := NewSegTree(n, oo, min)
	tl := NewSegTree(n, 0, max)

	cnt1 := make([]int, n)
	cnt2 := make([]int, n)

	for i, cur := range ops {
		x, l, r := cur[0], nums[cur[1]], nums[cur[2]]
		if x == 0 {
			// add
			cnt1[r]++
			if cnt1[r] == 1 {
				tr.Update(r, r)
			}
			cnt2[l]++
			if cnt2[l] == 1 {
				tl.Update(l, l)
			}
		} else {
			if cnt1[r] == 1 {
				tr.Update(r, oo)
			}
			cnt1[r]--
			if cnt2[l] == 1 {
				tl.Update(l, 0)
			}
			cnt2[l]--
		}
		l = tr.Get(0, n)
		r = tl.Get(0, n)
		ans[i] = l < r
	}
	return ans
}

type SegTree struct {
	size      int
	arr       []int
	initValue int
	op        func(int, int) int
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
	res := seg.initValue
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
