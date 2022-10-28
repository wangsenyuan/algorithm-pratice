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
		line := readString(reader)
		var i int
		for line[i] != ' ' {
			i++
		}
		n := line[:i]
		i++
		var x, m int
		i = readInt([]byte(line), i, &x) + 1
		readInt([]byte(line), i, &m)

		res := solve(n, m, x)

		buf.WriteString(fmt.Sprintf("%d\n", res))

	}

	fmt.Print(buf.String())
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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
const INV_MOD = 499122177

func pow(a, b, mod int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		b >>= 1
	}
	return res
}

func pow2(a int64, b string, mod int64) int64 {
	var res int64 = 1
	// b^1234 = (b^123)^10*b^4

	for i := 0; i < len(b); i++ {
		res = pow(res, 10, mod) * pow(a, int64(b[i]-'0'), mod) % mod
	}

	return res
}

func solve(n string, m int, x int) int {
	if x == 0 && n != "1" {
		return 1 % m
	}

	M := int64(m)
	X := int64(x)
	// (2 ** n - 1) % M
	b := (pow2(2, n, M) + M - 1) % M

	pc := baby_step_giant_step(b, X, M)
	if pc < 0 {
		return -1
	}

	res := pow2(2, n, MOD) * INV_MOD % MOD * (pow(2, pc, MOD) + MOD - 1)
	res %= MOD

	return int(res)
}

func baby_step_giant_step(g, t, m int64) int64 {
	sp := 1 % m
	pn := msb_pos(m)
	for i := 0; i < pn; i++ {
		if sp == t {
			return int64(i)
		}
		sp = sp * g % m
	}

	d := gcd(sp, m)

	if gcd(t, m)^d > 0 {
		return -1
	}

	var b int64 = 1
	gp := g

	for b*b < m {
		gp = gp * g % m
		b++
	}

	mp := make(map[int64]int64)

	for i, p := int64(1), t*g%m; i <= b; i++ {
		mp[p] = i
		p = p * g % m
	}

	for i, p := int64(1), sp*gp%m; i <= b; i++ {
		if mp[p] > 0 {
			return i*b - mp[p] + int64(pn)
		}
		p = p * gp % m
	}

	return -1
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func msb_pos(num int64) int {
	var res int
	for num > 0 {
		res++
		num >>= 1
	}
	return res
}
