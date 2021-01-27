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
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const MOD = 1000000009
const H = 30

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func inv(a int) int {
	return pow(a, MOD-2)
}

func modMul(a int, b int) int {
	A := int64(a)
	B := int64(b)
	return int(A * B % MOD)
}

func modAdd(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(n int, A []int) int {
	var res int

	for k := H - 1; k >= 0; k-- {
		// F[i, k, 0] & F[i, k, 1]
		a, b := 1, 0
		zero, one := 1, 1
		var m int
		for i := 0; i < n; i++ {
			if A[i]&(1<<uint(k)) > 0 {
				m ^= 1
				A[i] -= 1 << uint(k)
				one = modMul(one, A[i]+1)
				// F[i, k, 0] = F[i-1, k, 0] * 1 << k + F[i-1, k, 1] * (A[i] + 1)
				// 0 ^ 0 => 0, if set b[i] = 0, B[i] < A[i] for later bits, so, can take all 1 << k numbers
				// 1 ^ 1 => 0, if set[bi] = 1, B[i] = A[i] for now, can take [0...A[i]] numbers
				aa := modAdd(modMul(a, 1<<uint(k)), modMul(b, A[i]+1))
				// F[i, k, 1] = F[i-1, k, 1] * 1 << k + F[i-1, k, 0] * (A[i] + 1)
				// 1 ^ 0 => 1, 0 * 1 => 1
				bb := modAdd(modMul(b, 1<<uint(k)), modMul(a, A[i]+1))
				a, b = aa, bb
			} else {
				zero = modMul(zero, A[i]+1)
				// has to take B[i][k] = 0
				a = modMul(a, A[i]+1)
				b = modMul(b, A[i]+1)
			}
		}
		var c int
		if m == 0 {
			c = a
		} else {
			c = b
		}
		// res += (c - zero * one) / (1 << k)
		res = modAdd(res, modMul(modAdd(c, MOD-modMul(zero, one)), inv(1<<uint(k))))

	}

	return modAdd(res, 1)
}
