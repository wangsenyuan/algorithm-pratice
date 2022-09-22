package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' || s[i] == '\r' {
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
func solve(n int, k int) int {
	// 假设 n 的 质因数是, a, b, c
	// 就是 n * a ** cnt(a) * b ** cnt(b) * c * cnt(c)
	// cnt(a) + cnt(b) + cnt(c) <= k
	ans := 1
	for i := 2; i <= n/i; i++ {
		if n%i > 0 {
			continue
		}
		var cnt int
		for n%i == 0 {
			n /= i
			cnt++
		}
		ans = modMul(ans, sum_of_prime_factors(i, cnt, k), MOD)
	}

	if n > 1 {
		ans = modMul(ans, sum_of_prime_factors(n, 1, k), MOD)
	}

	return ans
}

func sum_of_prime_factors(p int, a int, k int) int {
	first := int64(pow(p, a, MOD))
	ratio := int64(p)
	terms := (int64(pow(2, k, MOD-1))*int64(a) - int64(a) + 1) % int64(MOD-1)
	res := (first * int64(pow(int(ratio), int(terms), MOD)-1)) % MOD
	res *= int64(pow(int(ratio-1), MOD-2, MOD))
	return int(res % MOD)
}

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int, mod int) int {
	r := int64(a) * int64(b)
	return int(r % int64(mod))
}

func pow(a, b, mod int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a, mod)
		}
		a = modMul(a, a, mod)
		b >>= 1
	}
	return r
}
