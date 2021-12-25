package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		var n, m uint64
		fmt.Scanf("%d %d\n", &n, &m)

		res := solve(n, int64(m))

		fmt.Println(res)
	}
}

const MOD = 1000000007

func solve(N uint64, M int64) int64 {
	P := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		P[i] = make([]int64, 2)
	}
	P[0][0] = (M - 1) % MOD
	P[0][1] = (M - 1) % MOD
	P[1][0] = 1

	R := matrixPow(P, N-1)

	F := make([][]int64, 2)
	F[0] = []int64{M % MOD}
	F[1] = []int64{0}

	R = matrixMul(R, F)

	ans := R[0][0] + R[1][0]
	ans %= MOD
	return ans
}

func matrixMul(a, b [][]int64) [][]int64 {
	m := len(a)
	n := len(b[0])
	c := make([][]int64, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int64, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < len(a[i]); k++ {
				c[i][j] += (a[i][k] * b[k][j]) % MOD
				c[i][j] %= MOD
			}
		}
	}
	return c
}

func matrixPow(mat [][]int64, n uint64) [][]int64 {
	m := len(mat)

	res := make([][]int64, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int64, m)
		res[i][i] = 1
	}

	mat1 := make([][]int64, m)
	for i := 0; i < m; i++ {
		mat1[i] = make([]int64, m)
		copy(mat1[i], mat[i])
	}

	for n != 0 {
		if n&1 == 1 {
			res = matrixMul(res, mat1)
		}
		mat1 = matrixMul(mat1, mat1)
		n >>= 1
	}
	return res
}
