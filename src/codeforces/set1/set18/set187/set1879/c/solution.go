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
		s := readString(reader)
		a, b := solve(s)

		buf.WriteString(fmt.Sprintf("%d %d\n", a, b))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const N = 200010

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

func nCr(n, r int) int {
	if r < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(s string) (int, int) {
	// 01 or 10
	n := len(s)
	a, b := solve1(s, 0)
	c, d := solve1(s, 1)
	if a > c {
		return n - a, b
	}
	if a < c {
		return n - c, d
	}
	return n - a, add(b, d)
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
	return a * b % mod
}

func solve1(s string, expect int) (int, int) {
	n := len(s)

	var set []int

	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		if len(set)%2 == 0 {
			if x == expect {
				set = append(set, x)
			}
		} else {
			if x != expect {
				set = append(set, x)
			}
		}
	}
	ln := len(set)
	rem := n - ln
	// 求的是序列
	ways := 1
	for i, j := 0, 0; i < len(set); i++ {
		var cnt int
		for j < n && int(s[j]-'0') == set[i] {
			cnt++
			j++
		}
		// 从cnt中选择一个
		ways = mul(ways, cnt)
	}

	ways = mul(ways, F[rem])

	return ln, ways
}
