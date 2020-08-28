package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	fmt.Println(solve(n, A))
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

const MAX_N = 100010
const MOD = 1000000007

var M [MAX_N]int
var m [MAX_N]int

const INF = 1 << 30
const N_INF = -INF

var prod [MAX_N]int64

const H = 15

func solve(n int, A []int) int {
	rmq := make([][]int, n)
	rxq := make([][]int, n)
	for i := 0; i < n; i++ {
		rmq[i] = make([]int, H)
		rmq[i][0] = i
		rxq[i] = make([]int, H)
		rxq[i][0] = i
	}
	for j := 1; j < H; j++ {
		for i := 0; i < n; i++ {
			rmq[i][j] = rmq[i][j-1]
			k := i + (1 << uint(j-1))
			if k < n && A[rmq[i][j-1]] > A[rmq[k][j-1]] {
				rmq[i][j] = rmq[k][j-1]
			}
			rxq[i][j] = rxq[i][j-1]

			if k < n && A[rxq[i][j-1]] < A[rxq[k][j-1]] {
				rxq[i][j] = rxq[k][j-1]
			}
		}
	}

	K := make([]int, n+1)
	var cur int
	for i := 1; i <= n; i++ {
		if 1<<uint(cur+1) == i {
			cur++
		}
		K[i] = cur
	}

	findMaxPos := func(i, j int) int {
		l := j - i + 1
		k := K[l]
		u := rxq[i][k]
		v := rxq[j-(1<<uint(k))+1][k]
		if A[u] > A[v] {
			return u
		}
		return v
	}

	findMinPos := func(i, j int) int {
		l := j - i + 1
		k := K[l]
		u := rmq[i][k]
		v := rmq[j-(1<<uint(k))+1][k]
		if A[u] < A[v] {
			return u
		}
		return v
	}

	mem := make(map[int64]int64)

	var dp func(i, j int) int64

	dp = func(i, j int) int64 {
		ij := int64(i)*int64(n) + int64(j)
		if r, found := mem[ij]; found {
			return r
		}
		if j <= i {
			return 1
		}
		a := findMaxPos(i, j)
		b := findMinPos(i, j)
		d := int64(A[a] - A[b])
		a, b = min(a, b), max(a, b)
		res := pow(d, int64(a-i+1)*int64(j-b+1))
		res = (res * dp(i, b-1)) % MOD
		res = (res * dp(a+1, j)) % MOD
		res = (res * inverse(dp(a+1, b-1))) % MOD
		mem[ij] = res
		return res
	}

	res := dp(0, n-1)

	return int(res)
}

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
			r %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return r
}

func inverse(num int64) int64 {
	return pow(num, MOD-2)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve2(n int, A []int) int {

	prod[0] = 1
	var res int64 = 1

	for j := 0; j < n; j++ {
		M[j] = N_INF
		m[j] = INF
		cur := A[j]
		k1 := j
		for k1 >= 0 && M[k1] < cur {
			M[k1] = cur
			k1--
		}
		k2 := j
		for k2 >= 0 && m[k2] > cur {
			m[k2] = cur
			k2--
		}
		k := min(k1, k2)
		if k < 0 {
			k = 0
		}
		for i := k; i < j; i++ {
			prod[i+1] = (prod[i] * int64(M[i]-m[i])) % MOD
		}

		res = (res * prod[j]) % MOD
	}
	return int(res)
}

func solve1(n int, A []int) int {

	prod[0] = 1
	var res int64 = 1
	M[0] = INF
	m[0] = N_INF
	M[1] = A[0]
	m[1] = A[0]
	for j := 2; j <= n; j++ {
		cur := A[j-1]
		M[j] = cur
		m[j] = cur
		k := j - 1
		for M[k] < cur {
			M[k] = cur
			k--
		}
		for m[k] > cur {
			m[k] = cur
			k--
		}
		k++
		for i := k; i < j; i++ {
			prod[i] = (prod[i-1] * int64(M[i]-m[i])) % MOD
		}

		res = (res * prod[j-1]) % MOD
	}

	return int(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
