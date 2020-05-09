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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)

		A := readNNums(reader, n)

		Q := readNum(reader)

		queries := make([][]int, Q)

		for i := 0; i < Q; i++ {
			queries[i] = readNNums(reader, 2)
		}

		res := solve(n, A, Q, queries)

		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

func solve(n int, A []int, Q int, queries [][]int) []int {

	qs := make([]Query, Q)

	for i := 0; i < Q; i++ {
		a, b := queries[i][0], queries[i][1]
		a--
		b--
		qs[i] = Query{a, b, i}
	}

	sort.Sort(Queries(qs))

	var tx int

	C := make([]int, MAX_N+1)
	CC := make([]int, MAX_N+1)

	BD := make([]bool, MAX_N+1)

	var D []int

	update := func(x, y int) {
		tx ^= C[x]
		CC[C[x]]--
		C[x] += y
		tx ^= C[x]
		CC[C[x]]++

		if !BD[C[x]] {
			D = append(D, C[x])
			BD[C[x]] = true
		}
	}

	res := make([]int, Q)

	L, R := 0, -1

	for i := 0; i < Q; i++ {
		q := qs[i]

		for q.left < L {
			L--
			update(A[L], 1)
		}

		for q.right > R {
			R++
			update(A[R], 1)
		}

		for q.left > L {
			update(A[L], -1)
			L++
		}

		for q.right < R {
			update(A[R], -1)
			R--
		}

		ND := make([]int, 0, len(D))

		var ans int64

		for _, c := range D {
			if CC[c] == 0 {
				BD[c] = false
				continue
			}

			c2 := c ^ tx
			if c2 < c {
				ans += (nCr(c, c2) * int64(CC[c])) % MOD
				ans %= MOD
			}
			ND = append(ND, c)
		}

		D = ND

		res[q.i] = int(ans)
	}

	return res
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[r]
	res %= MOD
	res *= IF[n-r]
	res %= MOD
	return res
}

const BLOCK_SIZE = 300

const MAX_N = 100000

const MOD = 998244353

var F []int64
var IF []int64

func pow(A, b int64) int64 {
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func init() {
	F = make([]int64, MAX_N+1)
	IF = make([]int64, MAX_N+1)

	F[0] = 1
	for i := 1; i <= MAX_N; i++ {
		F[i] = (int64(i) * F[i-1]) % MOD
	}

	IF[MAX_N] = pow(F[MAX_N], MOD-2)

	for i := MAX_N - 1; i >= 0; i-- {
		IF[i] = (int64(i+1) * IF[i+1]) % MOD
	}
}

type Query struct {
	left, right int
	i           int
}

type Queries []Query

func (this Queries) Len() int {
	return len(this)
}

func (this Queries) Less(i, j int) bool {
	a, b := this[i], this[j]
	if a.left/BLOCK_SIZE < b.left/BLOCK_SIZE {
		return true
	}
	if a.left/BLOCK_SIZE == b.left/BLOCK_SIZE {
		return a.right < b.right
	}
	return false
}

func (this Queries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
