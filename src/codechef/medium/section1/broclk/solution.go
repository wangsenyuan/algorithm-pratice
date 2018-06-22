package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readUInt64(bytes []byte, from int, val *uint64) int {
	i := from
	tmp := uint64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		scanner.Scan()
		var l, d int
		var t uint64

		pos := readInt(scanner.Bytes(), 0, &l)
		pos = readInt(scanner.Bytes(), pos+1, &d)
		readUInt64(scanner.Bytes(), pos+1, &t)

		fmt.Println(solve(l, d, t))
		tc--
	}
}

const MOD = 1000000007

func solve(l, d int, t uint64) int64 {
	L := int64(l)
	D := int64(d)
	// T := int64(t)
	cos_x := (D * pow(L, MOD-2)) % MOD

	mat := make(Matrix, 2)
	mat[0] = make([]int64, 2)
	mat[1] = make([]int64, 2)

	mat[0][0] = (2 * cos_x) % MOD
	mat[0][1] = MOD - 1
	mat[1][0] = 1
	mat = powMat(mat, t-1)

	res := (mat[0][0] * cos_x) % MOD
	res += mat[0][1]
	if res >= MOD {
		res -= MOD
	}

	res = (res * L) % MOD

	return res
}

func pow(a, n int64) int64 {
	res := int64(1)

	for n > 0 {
		if n&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		n >>= 1
	}
	return res
}

type Matrix [][]int64

func multiple(a, b Matrix) Matrix {
	m := len(a)

	c := make(Matrix, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int64, m)
	}

	for i := 0; i < m; i++ {
		for k := 0; k < m; k++ {
			for j := 0; j < m; j++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}
	return c
}

func uintMat() Matrix {
	c := make(Matrix, 2)
	for i := 0; i < 2; i++ {
		c[i] = make([]int64, 2)
		c[i][i] = 1
	}
	return c
}

func powMat(mat Matrix, n uint64) Matrix {
	if n == 0 {
		return uintMat()
	}
	res := make(Matrix, 2)
	tmp := make(Matrix, 2)

	for i := 0; i < 2; i++ {
		res[i] = make([]int64, 2)
		tmp[i] = make([]int64, 2)
		copy(res[i], mat[i])
		copy(tmp[i], mat[i])
	}
	n--
	for n > 0 {
		if n&1 == 1 {
			res = multiple(res, tmp)
		}
		tmp = multiple(tmp, tmp)
		n >>= 1
	}
	return res
}
