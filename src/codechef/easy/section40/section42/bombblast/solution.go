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
		X, Y := make([]int, m), make([]int, m)

		for i := 0; i < m; i++ {
			X[i], Y[i] = readTwoNums(reader)
		}

		res := solve(n, m, X, Y)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, m int, X []int, Y []int) []int64 {
	inc := make([][]int, n+1)
	dec := make([][]int, n)

	updateInc := func(pos, x int) {
		if inc[pos+1] == nil {
			inc[pos+1] = make([]int, 0, 1)
		}
		inc[pos+1] = append(inc[pos+1], 1)
		if pos+x+1 >= n {
			return
		}
		if inc[pos+x+1] == nil {
			inc[pos+x+1] = make([]int, 0, 1)
		}
		inc[pos+x+1] = append(inc[pos+x+1], -x-1)
	}

	updateDec := func(pos int, x int) {
		if pos > 0 {
			if dec[pos-1] == nil {
				dec[pos-1] = make([]int, 0, 1)
			}
			dec[pos-1] = append(dec[pos-1], 1)
		}
		if pos-x-1 >= 0 {
			if dec[pos-x-1] == nil {
				dec[pos-x-1] = make([]int, 0, 1)
			}
			dec[pos-x-1] = append(dec[pos-x-1], -x-1)
		}
	}

	for i := 0; i < m; i++ {
		x, y := X[i], Y[i]
		x--
		updateInc(x, y)
		updateDec(x, y)
	}

	ans := make([]int64, n)
	var cnt int64
	var sum int64
	var sqsum int64
	for i := 0; i < n; i++ {
		sqsum += 2*sum + cnt
		sum += cnt
		for _, x := range inc[i] {
			if x == 1 {
				cnt++
				sum++
				sqsum++
			} else {
				cnt--
				sum += int64(x)
				sqsum -= int64(x) * int64(x)
			}
		}
		ans[i] = sum + sqsum
	}
	cnt = 0
	sum = 0
	sqsum = 0
	for i := n - 1; i >= 0; i-- {
		sqsum += 2*sum + cnt
		sum += cnt
		for _, x := range dec[i] {
			if x == 1 {
				cnt++
				sum++
				sqsum++
			} else {
				cnt--
				sum += int64(x)
				sqsum -= int64(x) * int64(x)
			}
		}
		ans[i] -= sum + sqsum
	}
	return ans
}

func solve1(n int, m int, X []int, Y []int) []int64 {
	// dec (Xi - j) * (Xi - j + 1)
	// (Xi - j) * (Xi - j) + (Xi - j)
	// d * d + d
	// d * d = (xi - j) * (xi - j) = xi * xi - 2 * j * xi + j * j
	// d * d + d = xi * xi - 2 * j * xi + j * j + xi - j
	//           = xi * xi - (2 * j - 1) * xi + (j * j - j)
	// 分三个seg tree, 第一个部分记录 xi * xi
	// 第二部分记录 xi
	// 第三部分记录更新的次数 （这个可以用diff数组)
	// range update, point query
	// (k - xi) * (k - xi + 1)
	// d = k - xi
	// d * d + d = k * k - 2 * k * xi + xi * xi + k - xi
	//           = k * k + k - (2 * k + 1) * xi + xi * xi
	// diff1 for j * j
	diff1 := make([]int, n+2)
	// diff2 for j
	diff2 := make([]int, n+2)
	// for xi * xi
	xt := NewSegTree(n)
	// for xi * (2 * j - 1)
	jt := NewSegTree(n)
	// for xi * (2 * k + 1)
	kt := NewSegTree(n)

	for i := 0; i < m; i++ {
		x, y := X[i], Y[i]
		a := max(1, x-y)
		b := min(n, x+y)

		a--
		x--
		b--

		diff1[a]--
		diff1[x]++
		diff2[a]++
		diff2[x]--
		xt.RangeUpdate(a, x, -int64(X[i])*int64(X[i]))
		jt.RangeUpdate(a, x, int64(X[i]))
		diff1[x+1]++
		diff1[b+1]--
		diff2[x+1]++
		diff2[b+1]--
		xt.RangeUpdate(x+1, b+1, int64(X[i])*int64(X[i]))
		kt.RangeUpdate(x+1, b+1, -int64(X[i]))
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			diff1[i] += diff1[i-1]
			diff2[i] += diff2[i-1]
		}
		j := int64(i + 1)
		res[i] = int64(diff1[i]) * j * j
		res[i] += int64(diff2[i]) * j
		res[i] += xt.Get(i)
		res[i] += jt.Get(i) * (2*j - 1)
		res[i] += kt.Get(i) * (2*j + 1)
	}

	return res
}

type SegTree struct {
	arr []int64
	n   int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int64, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) RangeUpdate(l int, r int, v int64) {
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			t.arr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			t.arr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (t *SegTree) Get(p int) int64 {
	p += t.n

	var res int64

	for p >= 1 {
		res += t.arr[p]
		p >>= 1
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
