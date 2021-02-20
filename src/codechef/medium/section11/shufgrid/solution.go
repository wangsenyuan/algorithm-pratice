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
		line := readNNums(reader, 4)
		n, k, r, c := line[0], line[1], line[2], line[3]
		res := solve(n, k, r, c)
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

const MOD = 998244353

func binpow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
			res %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return res
}

type Mat [4][4]int64

func (this Mat) Mul(that Mat) Mat {
	var res [4][4]int64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				res[i][j] += (this[i][k] * that[k][j]) % MOD
				res[i][j] %= MOD
			}
		}
	}
	return res
}

var I Mat

func init() {
	for i := 0; i < 4; i++ {
		I[i][i] = 1
	}
}

func MatrixPow(m Mat, k int) Mat {
	if k == 0 {
		return I
	}
	r := MatrixPow(m, k/2)
	r = r.Mul(r)
	if k&1 == 1 {
		r = r.Mul(m)
	}
	return r
}

func solve(N, k, r, c int) int {
	n := int64(N)
	var T Mat
	T[0][0] = (n*n - n + 1) % MOD * binpow(n*n, MOD-2) % MOD
	T[0][1] = (n - 1) * binpow(2*n*n, MOD-2) % MOD
	T[0][2] = (n - 1) * binpow(2*n*n, MOD-2) % MOD
	T[0][3] = 0

	T[1][0] = binpow(2*n*n, MOD-2)
	T[1][1] = (2*n - 1) * binpow(2*n, MOD-2) % MOD
	T[1][3] = (n - 1) * binpow(2*n*n, MOD-2) % MOD

	T[2][0] = binpow(2*n*n, MOD-2)
	T[2][1] = 0
	T[2][2] = (2*n - 1) * binpow(2*n, MOD-2) % MOD
	T[2][3] = (n - 1) * binpow(2*n*n, MOD-2) % MOD

	T[3][0] = 0
	T[3][1] = binpow(2*n*n, MOD-2)
	T[3][2] = binpow(2*n*n, MOD-2)
	T[3][3] = (n*n - 1) * binpow(n*n, MOD-2) % MOD

	T = MatrixPow(T, k)

	return int(T[0][0])
}
