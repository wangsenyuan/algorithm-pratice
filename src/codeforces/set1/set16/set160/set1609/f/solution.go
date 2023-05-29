package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNInt64s(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

var INF = int64(1) << 62

const H = 62

func solve(A []int64) int64 {
	n := len(A)
	pc := make([]int, n)
	for i := 0; i < n; i++ {
		pc[i] = digit_count(A[i])
	}
	segs_min := make([]*Stack[Triple], H)
	segs_max := make([]*Stack[Triple], H)
	for i := 0; i < H; i++ {
		segs_min[i] = NewStack[Triple](1)
		segs_max[i] = NewStack[Triple](1)
	}
	st_min := NewStack[int](n)
	st_min.Push(-1)
	st_max := NewStack[int](n)
	st_max.Push(-1)

	countOne := func(v *Stack[Triple], up int) int {
		it := v.LowerBound(func(x Triple) bool {
			return x.first >= up+1
		})
		var res int
		if it > 0 {
			res = v.Get(it - 1).third
		}

		for it < v.Size() && v.Get(it).second <= up {
			res += up - v.Get(it).second + 1
			it++
		}
		return res
	}

	count := func(v *Stack[Triple], L int, R int) int {
		return countOne(v, R) - countOne(v, L-1)
	}

	var ans int64
	var cc int

	for i := 0; i < n; i++ {
		for st_min.Size() > 1 && A[st_min.Back()] > A[i] {
			cc -= count(segs_max[pc[st_min.Back()]], st_min.BackGet(-2)+1, st_min.Back())
			segs_min[pc[st_min.Back()]].Pop()
			st_min.Pop()
		}
		cc += count(segs_max[pc[i]], st_min.Back()+1, i)
		tmp := i - st_min.Back()

		if segs_min[pc[i]].Size() > 0 {
			tmp += segs_min[pc[i]].Back().third
		}

		segs_min[pc[i]].Push(Triple{i, st_min.Back() + 1, tmp})
		st_min.Push(i)

		for st_max.Size() > 1 && A[st_max.Back()] < A[i] {
			cc -= count(segs_min[pc[st_max.Back()]], st_max.BackGet(-2)+1, st_max.Back())
			segs_max[pc[st_max.Back()]].Pop()
			st_max.Pop()
		}

		cc += count(segs_min[pc[i]], st_max.Back()+1, i)
		tmp = i - st_max.Back()

		if segs_max[pc[i]].Size() > 0 {
			tmp += segs_max[pc[i]].Back().third
		}

		segs_max[pc[i]].Push(Triple{i, st_max.Back() + 1, tmp})
		st_max.Push(i)

		ans += int64(cc)
	}

	return ans
}

type Triple struct {
	first  int
	second int
	third  int
}

func solve1(A []int64) int64 {
	n := len(A)

	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, H)
	}

	X := make([]int64, n)
	Y := make([]int64, n)

	count := func(lx []int64, rx []int64, add int64) int64 {
		var mx, mn int64 = 0, INF
		for i := 0; i < len(rx); i++ {
			mx = max(mx, rx[i])
			mn = min(mn, rx[i])
			X[i] = mx
			Y[i] = mn
			copy(cnt[i+1], cnt[i])
			cnt[i+1][digit_count(mx)]++
		}
		mx = 0
		mn = INF

		var res int64

		var L, R int
		for i := len(lx) - 1; i >= 0; i-- {
			mx = max(mx, lx[i])
			for L < len(rx) && mx >= X[L] {
				L++
			}
			mn = min(mn, lx[i])
			for R < len(rx) && mn <= Y[R]-add {
				R++
			}
			b := digit_count(mn)
			if b == digit_count(mx) {
				res += min(int64(L), int64(R))
			}
			if L <= R {
				res += int64(cnt[R][b] - cnt[L][b])
			}
		}

		return res
	}

	merge := func(l int, mid int, r int) int64 {
		a := count(A[l:mid+1], A[mid+1:r+1], 0)
		reverse(A[l : mid+1])
		reverse(A[mid+1 : r+1])
		b := count(A[mid+1:r+1], A[l:mid+1], 1)
		reverse(A[l : mid+1])
		reverse(A[mid+1 : r+1])
		return a + b
	}

	var f func(l int, r int) int64

	f = func(l int, r int) int64 {
		if l == r {
			// max = min
			return 1
		}

		mid := (l + r) / 2
		a := f(l, mid)
		b := f(mid+1, r)
		c := merge(l, mid, r)
		return a + b + c
	}

	return f(0, n-1)
}

func reverse(arr []int64) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func digit_count(num int64) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	return a + b - max(a, b)
}

type Item interface {
	int | Triple
}

type Stack[T Item] struct {
	arr []T
	n   int
}

func NewStack[T Item](expectSize int) *Stack[T] {
	arr := make([]T, 0, expectSize)
	stack := Stack[T]{arr, 0}
	return &stack
}

func (s *Stack[T]) Size() int {
	return s.n
}

func (s *Stack[T]) Push(u T) {
	if s.n < len(s.arr) {
		s.arr[s.n] = u
		s.n++
		return
	}
	s.arr = append(s.arr, u)
	s.n = len(s.arr)
}

func (s *Stack[T]) Pop() T {
	res := s.arr[s.n-1]
	s.n--
	return res
}

func (s *Stack[T]) Back() T {
	return s.arr[s.n-1]
}

func (s *Stack[T]) Get(i int) T {
	return s.arr[i]
}

func (s *Stack[T]) BackGet(i int) T {
	return s.arr[i+s.n]
}

func (s *Stack[T]) LowerBound(f func(T) bool) int {
	l, r := 0, s.n
	for l < r {
		mid := (l + r) / 2
		if f(s.arr[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
