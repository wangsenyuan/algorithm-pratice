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

	n, k := readTwoNums(reader)
	p := readNNums(reader, k)

	res := solve(n, p)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, p []int) []int {
	k := len(p)
	// k >= 1
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
	}
	for i := 0; i < k; i++ {
		p[i]--
		pos[p[i]] = i
	}

	x := n - 1
	for i := 0; i < n; i++ {
		if pos[i] < 0 {
			x = i
			break
		}
	}
	// x是未出现的最小值
	// y是已出现的最大值
	// x必须在y的后面
	suf := make([]int, k+1)
	suf[k] = x
	for i := k - 1; i >= 0; i-- {
		suf[i] = min(p[i], suf[i+1])
	}

	tr := NewSegTree(n, 0, plus)

	for i := 0; i < k; i++ {
		// 在前面是否存在中间的数
		if suf[i+1] < p[i] {
			tmp := tr.Get(suf[i+1], p[i])
			if tmp > 0 {
				return nil
			}
		}
		tr.Update(p[i], 1)
	}

	// do have a solution
	res := make([]int, n)
	copy(res, p)
	sort.Ints(p)

	check := func(v int, j int) bool {
		if j == k {
			return v < n
		}
		return v < p[j]
	}

	for i, j := n-1, k-1; i >= k; j-- {
		var v int
		if j >= 0 {
			v = p[j] + 1
		}
		for i >= k && check(v, j+1) {
			res[i] = v
			v++
			i--
		}
	}
	for i := 0; i < n; i++ {
		res[i]++
	}

	return res
}

func plus(a, b int) int {
	return a + b
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
