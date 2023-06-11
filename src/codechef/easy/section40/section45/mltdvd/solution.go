package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	m := readNum(reader)
	Q := make([]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNum(reader)
	}
	res := solve(A, Q)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const MOD = 1e9 + 7
const MAX_A = 10_000_010

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
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

func inverse(a int) int {
	return pow(a, MOD-2)
}

var F []int

func init() {
	F = make([]int, MAX_A)
	for i := 2; i < MAX_A; i++ {
		if F[i] == 0 {
			F[i] = i
			if MAX_A/i < i {
				continue
			}
			for j := i * i; j < MAX_A; j += i {
				if F[j] == 0 {
					F[j] = i
				}
			}
		}
	}
}

func solve(A []int, Q []int) [][]int {
	lcm := make(map[int]int)
	var g int
	for _, num := range A {
		tmp := primeFactors(num)
		for k, v := range tmp {
			lcm[k] = max(lcm[k], v)
		}
		g = gcd(g, num)
	}

	L0 := prod(lcm)
	G0 := g

	gs := primeFactors(g)
	for k, v := range gs {
		lcm[k] -= v
	}

	L_G := prod(lcm)

	calc := func(x, k int) int {
		// x * (L / G)**k
		return modMul(x, pow(L_G, k))
	}

	ans := make([][]int, len(Q))

	for i, k := range Q {
		ans[i] = []int{calc(G0, k), calc(L0, k)}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func prod(fs map[int]int) int {
	res := 1
	for k, v := range fs {
		res = modMul(res, pow(k, v))
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func primeFactors(num int) map[int]int {
	res := make(map[int]int)

	for num > 1 {
		res[F[num]]++
		num /= F[num]
	}

	return res
}
