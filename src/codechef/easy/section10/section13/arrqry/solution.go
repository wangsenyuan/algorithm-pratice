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
		N, M := readTwoNums(scanner)
		A := readNNums(scanner, N)
		B := readNNums(scanner, M)
		Q := readNum(scanner)

		queries := make([][]int, Q)

		for i := 0; i < Q; i++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var t int
			pos := readInt(bs, 0, &t) + 1
			if t == 1 || t == 2 {
				queries[i] = make([]int, 4)
				queries[i][0] = t
				for j := 1; j < 4; j++ {
					pos = readInt(bs, pos, &queries[i][j]) + 1
				}
			} else {
				queries[i] = []int{t}
			}
		}

		res := solve(N, M, Q, A, B, queries)

		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

const MOD = 998244353

func solve(N, M, Q int, A []int, B []int, queries [][]int) []int64 {
	var sumA, sumB, ans int64

	for i := 0; i < N; i++ {
		sumA += int64(A[i])
		sumA %= MOD
	}

	for j := 0; j < M; j++ {
		sumB += int64(B[j])
		sumB %= MOD
		ans += int64(B[j]) * sumA
		ans %= MOD
	}

	res := make([]int64, 0, Q)

	for i := 0; i < Q; i++ {
		qry := queries[i]
		if qry[0] == 1 {
			l, r, x := qry[1], qry[2], qry[3]
			y := int64(r-l+1) * int64(x)
			y %= MOD
			sumA += y
			sumA %= MOD
			y = (y * sumB) % MOD
			ans += y
			ans %= MOD
		} else if qry[0] == 2 {
			l, r, x := qry[1], qry[2], qry[3]
			y := int64(r-l+1) * int64(x)
			y %= MOD
			sumB += y
			sumB %= MOD
			y = (y * sumA) % MOD
			ans += y
			ans %= MOD
		} else {
			res = append(res, ans)
		}
	}

	return res
}
