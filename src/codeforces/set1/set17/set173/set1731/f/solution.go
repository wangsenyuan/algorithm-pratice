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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, mod-2)
}

const N = 210

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = inv(F[N-1])

	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(n int, k int) int {
	values_of_polynomial := make([]int, n+3)

	for i := 1; i <= n; i++ {
		for x := 0; x < i; x++ {
			for y := x + 1; y <= n-i; y++ {
				for v := 1; v <= n+3; v++ {
					// 左边x个取值范围 [1...v-1]
					cur := mul(nCr(i-1, x), pow(v-1, x))
					cur = mul(cur, pow(k+1-v, i-1-x))
					cur = mul(cur, mul(nCr(n-i, y), pow(k-v, y)))
					cur = mul(cur, pow(v, n-i-y))
					cur = mul(cur, v)
					values_of_polynomial[v-1] = add(values_of_polynomial[v-1], cur)
				}
			}
		}
	}
	for i := 1; i < n+3; i++ {
		values_of_polynomial[i] = add(values_of_polynomial[i], values_of_polynomial[i-1])
	}

	if k <= n+3 {
		return values_of_polynomial[k-1]
	}
	return interpolate(values_of_polynomial, n+3, k)
}

func interpolate(y []int, r int, n int) int {
	var ans, prod = 0, 1
	for i := 1; i <= r; i++ {
		prod = mul(prod, n-i)
	}
	for i := 0; i < r; i++ {
		tmp := mul(prod, inv(n-i-1))
		tmp = mul(tmp, y[i])
		tmp = mul(tmp, I[i])
		tmp = mul(tmp, I[r-i-1])
		if (r-i)%2 == 0 {
			tmp = mod - tmp
		}
		ans = add(ans, tmp)
	}
	return ans
}
