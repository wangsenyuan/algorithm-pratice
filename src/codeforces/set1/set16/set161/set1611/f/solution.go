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
		n, s := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(s, A)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

func solve(s int, A []int) []int {
	// 假设从i开始，找到最远的j, 满足，在任何时刻 -A[j] >= sum (when A[j] < 0)
	// sum[j] = sum[j-1] + A[j] >= 0
	// sum[i] = s + A[i]
	//     => s + A[i] + A[i+1] + A[i+2] + ... + A[j] < 0
	//  A[i] + A[i+1] + A[i+2] + .. + A[j] 找到靠近左边的，pref_sum < -s 的位置
	n := len(A)
	res := []int{-1, -1}

	var sum int64
	v := int64(-s)

	for i, r := 0, 0; i < n; i++ {
		for sum >= v && r < n {
			sum += int64(A[r])
			r++
			if sum >= v && (res[0] < 0 || res[1]-res[0] < r-i) {
				res = []int{i + 1, r}
			}
		}

		sum -= int64(A[i])
	}

	if res[0] < 0 {
		return nil
	}

	return res
}

func solve1(s int, A []int) []int {
	// 假设从i开始，找到最远的j, 满足，在任何时刻 -A[j] >= sum (when A[j] < 0)
	// sum[j] = sum[j-1] + A[j] >= 0
	// sum[i] = s + A[i]
	//     => s + A[i] + A[i+1] + A[i+2] + ... + A[j] < 0
	//  A[i] + A[i+1] + A[i+2] + .. + A[j] 找到靠近左边的，pref_sum < -s 的位置
	n := len(A)
	t := NewSegTree(n)
	v := -int64(s)
	res := []int{-1, -1}
	for i := n - 1; i >= 0; i-- {
		t.Set(i, int64(A[i]))
		j := t.FindLeftPositionLessThanValue(v)
		if j > i {
			if res[0] < 0 || j-i > res[1]-res[0]+1 {
				res[0], res[1] = i+1, j
			}
		}
	}

	if res[0] < 0 {
		return nil
	}

	return res
}

type SegTree struct {
	val []int64
	sum []int64
	n   int
}

func NewSegTree(n int) *SegTree {
	val := make([]int64, 4*n)
	sum := make([]int64, 4*n)
	return &SegTree{val, sum, n}
}

func (t *SegTree) Set(p int, v int64) {

	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l == r {
			t.sum[i] = v
			t.val[i] = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i, l, mid)
		} else {
			loop(2*i+1, mid+1, r)
		}
		// 要找到小于t.s的最小的sum
		t.sum[i] = t.sum[2*i] + t.sum[2*i+1]
		t.val[i] = t.val[2*i]
		if t.sum[2*i]+t.val[2*i+1] < t.val[i] {
			t.val[i] = t.sum[2*i] + t.val[2*i+1]
		}
	}

	loop(1, 0, t.n-1)
}

// 在区间[L..R]上找到最靠左边的，< v的前缀和小标
func (t *SegTree) FindLeftPositionLessThanValue(v int64) int {

	var loop func(i int, l int, r int, v int64) int

	loop = func(i int, l int, r int, v int64) int {
		if l == r {
			if t.val[i] >= v {
				return l + 1
			}
			return l
		}
		mid := (l + r) / 2
		if t.val[2*i] < v {
			return loop(i*2, l, mid, v)
		}
		// t.val[2 * i] >= v
		// t.sum[2 * i] + t.val[2 * i + 1] < v
		// t.val[2 * i + 1] < v - t.sum[2 * i]
		return loop(2*i+1, mid+1, r, v-t.sum[2*i])
	}

	if t.val[1] > v {
		return t.n
	}
	return loop(1, 0, t.n-1, v)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
