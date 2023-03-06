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
		n := readNum(reader)
		ok, res := solve(n)
		if !ok {
			buf.WriteString(fmt.Sprintf("%d\n", res))
		} else {
			buf.WriteString(fmt.Sprintf("%04d\n", res))
		}
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

const MOD = 1e4

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
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

const MAX_X = 5*1e5 + 10

var ps []int
var ds []int

func init() {
	ps = make([]int, MAX_X)
	ds = make([]int, MAX_X)

	for i := 2; i < MAX_X; i++ {
		if i < MAX_X/i {
			ps[i*i] = i
		}

		for j := i; j < MAX_X; j += i {
			ds[j]++
		}
	}
}

func solve(n int) (bool, int) {
	// t is about 300000
	// n <= 5 * 1e5
	// 假设 n的 factors 是 1, a, b, c, d...
	// 需要计算 a * b * c * d ....  % 10000
	// 假设a的质数分解是 p1 ** x1 * p2 ** x2 ... pk ** xk
	// 1 * a * b * c ...
	// p1 ** sum(0 + x1) * (x1 + 1) / 2 * p2 ** sum(0 + x2) * (x2 + 1) / 2 ...
	// 处于n, 但是问题是 10000不是质数，所以，还要对10000进行质数分解

	if n == 1 || ds[n] == 1 {
		// a prime number
		return false, 1
	}

	res := bruteForce(n)

	if res < MOD {
		return false, res
	}

	res = pow(n, (ds[n]-1)/2)

	if ps[n] > 0 {
		res = modMul(res, ps[n])
	}
	return true, res
}

func bruteForce(n int) int {
	var res int64 = 1
	for i := 0; i < (ds[n]-1)/2; i++ {
		res *= int64(n)
		if res >= MOD {
			return MOD
		}
	}

	if ps[n] > 0 {
		res *= int64(ps[n])
	}
	if res >= MOD {
		return MOD
	}
	return int(res)
}
