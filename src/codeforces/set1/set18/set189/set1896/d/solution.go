package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				var x int
				readInt(s, 2, &x)
				ops[i] = []int{1, x}
			} else {
				var j, x int
				pos := readInt(s, 2, &j)
				readInt(s, pos+1, &x)
				ops[i] = []int{2, j, x}
			}
		}
		res := solve(a, ops)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func solve(a []int, ops [][]int) []bool {
	var sum int
	n := len(a)

	mn := NewSegTree(n, n, min)
	mx := NewSegTree(n, -1, max)

	for i := 0; i < n; i++ {
		sum += a[i]
		if a[i] == 1 {
			mn.Update(i, i)
			mx.Update(i, i)
		}
	}
	best := make([]int, 2)

	update := func() {
		l := mn.Get(0, n)
		r := mx.Get(0, n)
		if sum%2 == 0 {
			best[0] = sum
			if l < n {
				best[1] = max(sum-2*l-1, 2*l+1)
				best[1] = max(best[1], max(sum-2*(n-r-1)-1, 2*(n-r-1)+1))
			} else {
				best[1] = 0
			}
		} else {
			best[1] = sum
			if l < n {
				best[0] = max(sum-2*l-1, 2*l)
				best[0] = max(best[0], max(sum-2*(n-r-1)-1, 2*(n-r-1)))
			} else {
				best[0] = 0
			}
		}
	}

	update()

	check := func(s int) bool {
		return s <= best[s&1]
	}

	var ans []bool

	for _, cur := range ops {
		if cur[0] == 1 {
			ans = append(ans, check(cur[1]))
		} else {
			i, x := cur[1]-1, cur[2]
			if a[i] == 1 {
				mn.Update(i, n)
				mx.Update(i, -1)
			}
			sum -= a[i]
			a[i] = x
			sum += a[i]
			if a[i] == 1 {
				mn.Update(i, i)
				mx.Update(i, i)
			}
			update()
		}
	}

	return ans
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
