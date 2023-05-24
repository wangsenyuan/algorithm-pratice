package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	res := solve(n, m)
	fmt.Println(res)
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const MOD = 998244353
const MAX_N = 2021*2 + 10

var F [MAX_N]int
var I [MAX_N]int

func modAdd(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a int, b int) int {
	res := int64(a) * int64(b)
	res %= MOD
	return int(res)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = modMul(i, F[i-1])
	}
	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)
	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = modMul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func solve(n int, m int) int {
	// must be dp, but how?
	if n < m {
		n, m = m, n
	}
	var ans int
	for r := 1; r < n; r++ {
		var f int
		for c := m - 1; c >= 1; c-- {
			f = modAdd(f, modMul(nCr(n-r+m-c-1, n-r), nCr(c+n-r-1, n-r-1)))
			ans = modAdd(ans, modMul(nCr(r+c-1, r), modMul(nCr(m-c+r-1, r-1), f)))
		}
	}

	n, m = m, n

	for r := 1; r < n; r++ {
		var f int
		for c := m - 1; c >= 1; c-- {
			ans = modAdd(ans, modMul(nCr(r+c-1, r), modMul(nCr(m-c+r-1, r-1), f)))
			f = modAdd(f, modMul(nCr(n-r+m-c-1, n-r), nCr(c+n-r-1, n-r-1)))
		}
	}

	return modMul(ans, 2)
}
