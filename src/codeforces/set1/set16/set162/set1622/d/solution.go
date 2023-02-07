package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	s, _ := reader.ReadString('\n')
	res := solve(n, k, s)
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

const N = 5010

var F []int64
var I []int64

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func init() {
	F = make([]int64, N)
	I = make([]int64, N)
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(int64(i), F[i-1])
	}
	I[N-1] = int64(inverse(int(F[N-1])))

	for i := N - 2; i >= 0; i-- {
		I[i] = mul(int64(i+1), I[i+1])
	}
}

func nCr(n, r int) int64 {
	if n < 0 || n < r {
		return 0
	}
	res := F[n]
	res = mul(res, I[r])
	res = mul(res, I[n-r])
	return res
}

func mul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}

func add(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int64) int64 {
	return add(a, MOD-b)
}

func solve(n int, k int, s string) int {
	var ones []int
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ones = append(ones, i)
		}
	}

	if len(ones) < k || k == 0 {
		return 1
	}

	var tot int64 = 1
	for i, cur := range ones {
		var start int
		if i > 0 {
			start = ones[i-1] + 1
		}
		var end int = n - 1
		if i+k < len(ones) {
			end = ones[i+k] - 1
		}
		tot = add(tot, nCr(end-start+1, min(k, len(ones)-i)))
		tot = sub(tot, nCr(end-cur, min(k, len(ones)-i)-1))
	}

	return int(tot)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
