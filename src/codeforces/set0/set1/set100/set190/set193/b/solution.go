package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, u, r := readThreeNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	K := readNNums(reader, n)
	P := readNNums(reader, n)
	res := solve(u, r, A, B, K, P)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(u int, r int, A []int, B []int, K []int, P []int) int64 {
	// [a, b] * [c, d] when a < b && c < d
	// 1) a * c + b * d  1 * 1 + 2 * 2 = 5
	// 2) a * d + b * c  1 * 2 + 1 * 2 = 4
	//  a * c + b * d - a * d - b * c = a * (c - d) + b * (d - c) = (a - b) * (c - d) >= 0
	// 从例子中看到排序后相乘更合适
	// 把P中循环的部分放在一起，按照K排序
	// A[i] ^= B[i] 多次的xor操作没有意义，
	// 但是 (A[i] + r) ^ B[i] 就不一定了

	n := len(A)
	for i := 0; i < n; i++ {
		P[i]--
	}

	op1 := func(a []int) []int {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = a[i] ^ B[i]
		}
		return res
	}

	op2 := func(a []int) []int {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = a[P[i]] + r
		}
		return res
	}

	calc := func(a []int) int64 {
		var res int64
		for i := 0; i < n; i++ {
			res += int64(a[i]) * int64(K[i])
		}
		return res
	}

	var dfs func(a []int, l int, x bool) int64

	dfs = func(a []int, l int, x bool) int64 {

		// xor + shift
		// or just shift
		var res int64 = math.MinInt64

		if (u-l)&1 == 0 {
			res = calc(a)
		}

		if l == u {
			return res
		}

		res = max(res, dfs(op2(a), l+1, false))
		if !x {
			res = max(res, dfs(op1(a), l+1, true))
		}
		return res
	}

	return dfs(A, 0, false)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
