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
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(s, 0, &n) + 1
		var L, R uint64
		pos = readUint64(s, pos, &L) + 1
		readUint64(s, pos, &R)
		A := readNNums(reader, n)
		res := solve(A, int64(L), int64(R))

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

const MAX_A = 10000010

func solve(A []int, L, R int64) int64 {
	// concate(A[i], A[j]) >= L <= R
	// find count_int_con(A) > R
	// if we fix at some position j, let it is the tail part,
	// then we nee to know how many numbers before j, which >= X / L[j]
	sort.Ints(A)
	// order of A donesn't matter
	n := len(A)

	B := make([]int64, n)
	for i := 0; i < n; i++ {
		B[i] = getBase(A[i])
	}

	count := func(X int64) int64 {
		var res int64
		for j := 0; j < n; j++ {
			// if A[j] is the tail part
			// we need find some A[i] * B[j] + A[j] >= X
			// A[i] * B[j] >= X - A[j]
			i := search(n, func(i int) bool {
				return int64(A[i])*B[j]+int64(A[j]) >= X
			})
			res += int64(n - i)
		}
		return res
	}

	return count(L) - count(R+1)
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func solve1(A []int, L, R int64) int64 {
	// concate(A[i], A[j]) >= L <= R
	// find count_int_con(A) > R
	// if we fix at some position j, let it is the tail part,
	// then we nee to know how many numbers before j, which >= X / L[j]
	bit := NewBit(MAX_A)

	n := len(A)

	B := make([]int64, n)
	for i := 0; i < n; i++ {
		B[i] = getBase(A[i])
	}

	count := func(X int64) int64 {
		var res int64
		// if we fix j as the tail part, not including (j, j)
		for j := 0; j < n; j++ {
			x, r := X/B[j], X%B[j]
			var c int
			if r > int64(A[j]) {
				x++
			}
			if x < MAX_A {
				c = bit.GetRange(int(x), MAX_A-1)
			}
			res += int64(c)
			bit.Add(A[j], 1)
		}
		// reset bit
		for j := 0; j < n; j++ {
			bit.Add(A[j], -1)
		}
		// if we fix i as the tail part, including (i, i)
		for i := n - 1; i >= 0; i-- {
			bit.Add(A[i], 1)
			x, r := X/B[i], X%B[i]
			if r > int64(A[i]) {
				x++
			}
			var c int
			if x < MAX_A {
				c = bit.GetRange(int(x), MAX_A-1)
			}
			res += int64(c)
		}

		for j := 0; j < n; j++ {
			bit.Add(A[j], -1)
		}
		return res
	}

	return count(L) - count(R+1)
}

func getBase(num int) int64 {
	var res int64 = 1
	for num > 0 {
		res *= 10
		num /= 10
	}
	return res
}

type Bit struct {
	arr []int
	n   int
}

func NewBit(n int) *Bit {
	arr := make([]int, n+1)
	return &Bit{arr, n}
}

func (bit *Bit) Add(p int, v int) {
	p++
	for p <= bit.n {
		bit.arr[p] += v
		p += p & -p
	}
}

func (bit *Bit) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit.arr[p]
		p -= p & -p
	}
	return res
}

func (bit *Bit) GetRange(l, r int) int {
	res := bit.Get(r)
	if l > 0 {
		res -= bit.Get(l - 1)
	}
	return res
}
