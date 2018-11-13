package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1e9 + 7
const MAX_N = 2010

var C [][]int64

func init() {
	C = make([][]int64, MAX_N)
	for i := 0; i < MAX_N; i++ {
		C[i] = make([]int64, MAX_N)
	}
	C[0][0] = 1
	for i := 1; i < MAX_N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
			if C[i][j] >= MOD {
				C[i][j] -= MOD
			}
		}
	}
}

func solve(n int, A []int) int {
	var res int64
	for i := 1; i <= n; i += 2 {
		res = modAdd(res, C[n][i])
	}

	sort.Ints(A)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n && A[i] == A[j]; j++ {
			x := i
			y := n - 1 - j
			res = modAdd(res, C[x+y][min(x, y)])
		}
	}

	return int(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
