package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		x, y := readTwoNums(reader)
		res := solve(x, y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const N = 2000010
const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
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

var lps [N]int
var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}

	I[N-1] = pow(F[N-1], mod-2)

	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	var primes []int
	for i := 2; i < N; i++ {
		if lps[i] == 0 {
			lps[i] = i
			primes = append(primes, i)
		}

		for j := 0; j < len(primes); j++ {
			if i*primes[j] >= N {
				break
			}
			lps[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
}

func nCr(n, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(x int, y int) int {

	var arr []int

	for tmp := x; tmp > 1; {
		u := lps[tmp]
		var cnt int
		for tmp%u == 0 {
			cnt++
			tmp /= u
		}
		arr = append(arr, cnt)
	}

	res := 1

	for _, f := range arr {
		res = mul(res, nCr(f+y-1, f))
	}
	res = mul(res, pow(2, y-1))

	return res
}
