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

	n, k := readTwoNums(reader)

	fmt.Println(solve(n, k))
}

const MAX_N = 500043
const MOD = 998244353

var F []int64

// var PM []int

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

func init() {
	F = make([]int64, MAX_N)
	F[0] = 1

	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1]
		F[i] %= MOD
	}
}

func IF(n int) int64 {
	x := F[n]
	return pow(x, MOD-2)
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF(r)
	res %= MOD
	res *= IF(n - r)
	res %= MOD
	return res
}

func solve(n, k int) int {
	var res int64

	for x := 1; x <= n; x++ {
		// x := PM[i]
		d := n / x
		if d < k {
			break
		}
		res += nCr(d-1, k-1)
		res %= MOD
	}

	return int(res)
}
