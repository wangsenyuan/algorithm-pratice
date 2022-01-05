package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(n, m, A)
	fmt.Printf("%d\n", res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const MOD = 1000000007

const MAX_N = 200010

var fact []int64
var ifact []int64

func init() {
	fact = make([]int64, MAX_N+1)
	ifact = make([]int64, MAX_N+1)

	fact[0] = 1
	for i := 1; i <= MAX_N; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	ifact[MAX_N] = inverse(fact[MAX_N])

	for i := MAX_N; i > 0; i-- {
		ifact[i-1] = (int64(i) * ifact[i]) % MOD
	}
}

func inverse(q int64) int64 {
	return pow(q, MOD-2)
}

func pow(a, b int64) int64 {
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func nCr(n int, r int) int64 {
	res := (ifact[r] * ifact[n-r]) % MOD
	res *= fact[n]
	return res % MOD
}

func solve(n int, m int, B []int) int {
	if m < n {
		return 0
	}
	var p int
	for _, b := range B {
		if b <= m {
			p++
		}
	}
	ans := modMul(nCr(m, n), fact[n])

	sign := int64(-1)

	for x := 1; x <= p; x++ {
		tmp := modMul(nCr(p, x), modMul(nCr(m-x, n-x), fact[n-x]))

		if sign < 0 {
			ans = modSub(ans, tmp)
		} else {
			ans = modAdd(ans, tmp)
		}
		sign *= -1
	}

	return int(ans)
}

func modMul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int64) int64 {
	return modAdd(a, MOD-b)
}
