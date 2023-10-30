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
		n := readNum(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func sub(a, b int) int {
	return add(a, mod-b)
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

func inverse(x int) int {
	return pow(x, mod-2)
}

func solve(x string, y string) int {
	n := len(x)
	var diff int
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			diff++
		}
	}
	if diff == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a := make([]int, n+10)
	b := make([]int, n+10)
	a[1] = 1
	b[1] = mul(n-1, inverse(n))

	for i := 2; i <= n; i++ {
		tmp := inverse(sub(n, mul(i, b[i-1])))
		a[i] = mul(add(n, mul(i, a[i-1])), tmp)
		b[i] = mul(n-i, tmp)
	}

	c := make([]int, n+10)
	d := make([]int, n+10)
	c[n] = 1
	d[n] = 1
	for i := n - 1; i >= 1; i-- {
		tmp := inverse(sub(n, mul(n-i, d[i+1])))
		c[i] = mul(add(n, mul(n-i, c[i+1])), tmp)
		d[i] = mul(i, tmp)
	}

	var all bool
	if diff == n {
		diff--
		all = true
	}

	// c[i] + d[i] * a[i-1]
	u := add(c[diff], mul(d[diff], a[diff-1]))
	// 1 - d[i] * b[i-1]
	v := sub(1, mul(d[diff], b[diff-1]))
	ans := mul(u, inverse(v))
	if all {
		ans = add(ans, 1)
	}
	return ans
}
