package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if res {
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
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int) bool {
	// max(A[l...r]) >= sum(A[l...r])
	// for all (l, r) pair
	// 不能出现连续的正值，连续的负值是ok的，连续的负值，可以收缩成一个值
	n := len(A)
	pref := make([]int64, n)
	stack := make([]int, n)
	L := make([]int, n)
	R := make([]int, n)

	var p int
	for i := 0; i < n; i++ {
		pref[i] = int64(A[i])
		if i > 0 {
			pref[i] += pref[i-1]
		}
		R[i] = n
		for p > 0 && A[stack[p-1]] < A[i] {
			R[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}
	pt := NewSegTree(pref)
	p = 0
	suf := make([]int64, n)
	for i := n - 1; i >= 0; i-- {
		suf[i] = int64(A[i])
		if i+1 < n {
			suf[i] += suf[i+1]
		}
		L[i] = -1
		for p > 0 && A[stack[p-1]] < A[i] {
			L[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}

	st := NewSegTree(suf)

	for i := 0; i < n; i++ {
		a := st.Get(L[i]+1, i) - suf[i]
		b := pt.Get(i+1, R[i]) - pref[i]
		if max(a, b) > 0 {
			return false
		}
	}
	return true
}

type SegTree struct {
	arr []int64
	n   int
}

func NewSegTree(arr []int64) *SegTree {
	n := len(arr)
	nums := make([]int64, 2*n)
	copy(nums[n:], arr)
	for i := n - 1; i > 0; i-- {
		nums[i] = max(nums[i*2], nums[i*2+1])
	}
	return &SegTree{nums, n}
}

func (t *SegTree) Get(l int, r int) int64 {
	l += t.n
	r += t.n
	var res int64 = math.MinInt64 >> 1

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

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
