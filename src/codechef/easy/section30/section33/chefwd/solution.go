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
		var n uint64
		s, _ := reader.ReadBytes('\n')
		readUint64(s, 0, &n)
		res := solve(int64(n))
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	r %= MOD
	return int(r)
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

var inv_5 int

var F [6]int

func init() {
	inv_5 = pow(5, MOD-2)
	F[0] = 0
	F[1] = 1
	F[2] = 1
	F[3] = 2
	F[4] = 3
	F[5] = 5
}

func solve(n int64) int {
	m := int(n % MOD)
	n += 5
	A := make([][]int, 2)
	B := make([][]int, 2)
	C := make([][]int, 2)
	for i := 0; i < 2; i++ {
		A[i] = make([]int, 2)
		B[i] = make([]int, 2)
		C[i] = make([]int, 2)
	}
	A[0][0] = 1
	A[1][1] = 1
	B[0][0] = 1
	B[0][1] = 1
	B[1][0] = 1
	for n > 0 {
		if n&1 == 1 {
			C[0][0] = A[0][0]
			C[0][1] = A[0][1]
			C[1][0] = A[1][0]
			C[1][1] = A[1][1]
			A[0][0] = 0
			A[0][1] = 0
			A[1][0] = 0
			A[1][1] = 0
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					for k := 0; k < 2; k++ {
						A[i][j] = modAdd(A[i][j], modMul(C[i][k], B[k][j]))
					}
				}
			}
		}
		C[0][0] = B[0][0]
		C[0][1] = B[0][1]
		C[1][0] = B[1][0]
		C[1][1] = B[1][1]
		B[0][0] = 0
		B[0][1] = 0
		B[1][0] = 0
		B[1][1] = 0

		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					B[i][j] = modAdd(B[i][j], modMul(C[i][k], C[k][j]))
				}
			}
		}
		n >>= 1
	}

	fa := A[0][0]
	fb := A[1][0]

	res := modMul(modAdd(modMul(2, m), 10), fa)
	res = modSub(res, modMul(modAdd(m, 6), fb))

	res = modMul(res, inv_5)

	for i := 0; i < 5; i++ {
		fa, fb = fb, modSub(fa, fb)
		if i < 2 {
			res = modSub(res, modMul(2, modMul(F[i], fa)))
		} else {
			res = modSub(res, modMul(F[i], fa))
		}
	}

	return res
}
