package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, S := readTwoNums(scanner)
		A := readNNums(scanner, N)
		res := solve(N, S, A)
		fmt.Println(res)
	}
}

const MAX = 1e5 + 5

var PA Pairs
var W Pairs
var seg []int

func init() {
	PA = make(Pairs, MAX)
	seg = make([]int, 2*MAX)
	W = make(Pairs, MAX)
}

func solve(N, S int, A []int) int {
	P := zipWithIndex(A)
	sort.Sort(P)

	build := func() {
		for i := 0; i < 2*N; i++ {
			seg[i] = 0
		}
	}

	update := func(pos, val int) {
		pos--
		pos += N
		seg[pos] = val
		for pos > 1 {
			seg[pos>>1] = max(seg[pos], seg[pos^1])
			pos >>= 1
		}
	}

	query := func(l, r int) int {
		l--
		l += N
		r += N
		var res int

		for l < r {
			if l&1 == 1 {
				res = max(res, seg[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, seg[r])
			}
			l >>= 1
			r >>= 1
		}

		return res
	}

	check := func(k int) bool {
		build()
		var sum int64

		for i := 1; i <= N; i++ {
			j := i
			for j <= N && P[i-1].first == P[j-1].first {
				j++
			}
			var idx int
			for u := i; u < j; u++ {
				l := max(1, P[u-1].second-k)
				r := min(N, P[u-1].second+k)
				x := query(l, r) + 1
				W[idx] = Pair{x, P[u-1].second}
				idx++
			}

			for u := 0; u < idx; u++ {
				v := u + 1
				for v < idx && W[v].second <= W[v-1].second+k {
					v++
				}
				var value int
				for x := u; x < v; x++ {
					value = max(value, W[x].first)
				}
				for x := u; x < v; x++ {
					sum += int64(value)
					update(W[x].second, value)
				}
				u = v - 1
			}
			i = j - 1
		}
		return sum <= int64(S)
	}

	lo, hi := 0, N+1

	for lo < hi {
		mid := (lo + hi) >> 1
		if !check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
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

func zipWithIndex(A []int) Pairs {
	n := len(A)
	for i := 0; i < n; i++ {
		PA[i] = Pair{A[i], i + 1}
	}
	return PA[:n]
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
