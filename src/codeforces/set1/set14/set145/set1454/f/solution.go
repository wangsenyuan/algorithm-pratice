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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < 3; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
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

func solve(A []int) []int {
	// max(1...l) = min(l + 1.... r - 1) = max(r...n)
	// 从左往右，对于i,如果以i为最大值，然后找到一个连续的区间[l...r] >= A[i], 且在区间 i...l中间都 <= A[i]
	n := len(A)
	x := A[0]
	for i := 1; i < n; i++ {
		x = max(x, A[i])
	}

	var idx []int
	for i := 0; i < n; i++ {
		if A[i] == x {
			idx = append(idx, i)
		}
	}

	if len(idx) >= 3 {
		return []int{idx[1], 1, n - idx[1] - 1}
	}
	mn := NewTree(1<<30, A, func(a, b int) int {
		return min(a, b)
	})
	mx := NewTree(0, A, func(a, b int) int {
		return max(a, b)
	})
	l, r := idx[0], idx[len(idx)-1]+1
	v := mn.Get(l, r)

	for l > 0 && r < n {
		if mx.Get(0, l) == v && v == mx.Get(r, n) {
			return []int{l, r - l, n - r}
		}
		// if expand to left
		var a int
		if l > 1 {
			a = min(mx.Get(0, l-1), mn.Get(l-1, r), mx.Get(r, n))
		}
		var b int
		if r+1 < n {
			b = min(mx.Get(0, l), mn.Get(l, r+1), mx.Get(r+1, n))
		}
		if a >= b {
			l--
		} else {
			r++
		}

		v = mn.Get(l, r)
	}

	return nil
}

type Tree struct {
	arr []int
	n   int
	iv  int
	fn  func(int, int) int
}

func NewTree(iv int, arr []int, fn func(int, int) int) *Tree {
	res := make([]int, len(arr)*2)
	copy(res[len(arr):], arr)
	for i := len(arr) - 1; i > 0; i-- {
		res[i] = fn(res[2*i], res[2*i+1])
	}

	return &Tree{res, len(arr), iv, fn}
}

func (t *Tree) Get(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = t.fn(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.fn(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func reverse(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
