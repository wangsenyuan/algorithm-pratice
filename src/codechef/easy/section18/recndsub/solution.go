package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)

		res := solve(n, A)
		for i := 0; i < len(res); i++ {
			fmt.Printf("%d ", res[i])
		}
		fmt.Println()
	}
}

const MOD = 163577857

const MAX_N = 100001

var F []int64
var IF []int64

func init() {
	F = make([]int64, MAX_N)

	F[0] = 1

	for i := 1; i < MAX_N; i++ {
		F[i] = (int64(i) * F[i-1]) % MOD
	}

	IF = make([]int64, MAX_N)
	IF[MAX_N-1] = pow(int(F[MAX_N-1]), MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = (int64(i+1) * IF[i+1]) % MOD
	}
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

func solve(n int, A []int) []int {
	cnt := []int{0, 0, 0}

	for i := 0; i < n; i++ {
		cnt[1+A[i]]++
	}

	X := pow(2, cnt[1])

	C := make([]int, 2*n+1)

	var i int
	for k := n; k > 0; k-- {
		C[i] = int((X * nCr(cnt[0]+cnt[2], cnt[2]+k)) % MOD)
		i++
	}

	C[i] = int((X * nCr(cnt[0]+cnt[2], cnt[2])) % MOD)
	C[i] = (C[i] + MOD - 1) % MOD
	i++

	for k := 1; k <= n; k++ {
		C[i] = int((X * nCr(cnt[0]+cnt[2], cnt[0]+k)) % MOD)
		i++
	}

	return C
}

func solve1(n int, A []int) []int {

	cnt := []int{0, 0, 0}

	for i := 0; i < n; i++ {
		cnt[1+A[i]]++
	}

	F := pow(2, cnt[1])

	// row := int64(cnt[0] + cnt[2])

	C := make([]int, 2*n+1)

	var prod int64 = 1

	for i := -n; i <= n; i++ {
		var x int64

		if i >= -cnt[0] && i <= cnt[2] {
			x = prod
			prod = (prod * int64(cnt[2]-i)) % MOD
			prod = (prod * pow(i+cnt[0]+1, MOD-2)) % MOD
		}

		x = (x * F) % MOD

		if i == 0 {
			x = (x + MOD - 1) % MOD
		}

		C[i+n] = int(x)
	}

	return C
}

func pow(a, b int) int64 {
	A := int64(a)
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
