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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		t--
		N, K := readTwoNums(scanner)
		fillNNums(scanner, N, A)
		for i := 0; i < K; i++ {
			fillNNums(scanner, 4, queries[i])
		}
		res := solve(N, K, A, queries)
		for _, ans := range res {
			if ans {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

const MOD = 1e9 + 7

const MAX_N = 1e5 + 5

var A []int
var S []int64
var S2 []int64
var X []int
var queries [][]int
var res []bool

func init() {
	A = make([]int, MAX_N)
	S = make([]int64, MAX_N)
	S2 = make([]int64, MAX_N)
	X = make([]int, MAX_N)
	queries = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		queries[i] = make([]int, 4)
	}
	res = make([]bool, MAX_N)
}

func solve(n int, K int, A []int, queries [][]int) []bool {

	for i := 0; i < n; i++ {
		S[i+1] = modAdd(S[i], int64(A[i]))
		S2[i+1] = modAdd(S2[i], int64(A[i])*int64(A[i]))
		X[i+1] = X[i] ^ A[i]
	}

	window := func(a, b, x int) (int, int) {
		i, j := 0, 0

		for k := a; k <= b; k++ {
			if A[k] < x {
				i++
			}
			if A[k] == x {
				j++
			}
		}
		return i, i + j - 1
	}

	simliar := func(a, b, c, d, x, y int) bool {
		a, b = window(a, b, x)
		c, d = window(c, d, y)
		return !(b < c || d < a)
	}

	// res := make([]bool, K)

	for i := 0; i < K; i++ {
		query := queries[i]
		a, b, c, d := query[0], query[1], query[2], query[3]
		t1 := modSub(S[b], S[a-1])
		t2 := modSub(S2[b], S2[a-1])
		T1 := modSub(S[d], S[c-1])
		T2 := modSub(S2[d], S2[c-1])
		ab := X[b] ^ X[a-1]
		cd := X[d] ^ X[c-1]

		res[i] = false
		if t1 == T1 && t2 == T2 && ab == cd {
			res[i] = true
		} else if t1 != T1 {
			// can calculate x, y
			if (T2-t2)%(T1-t1) == 0 {
				tt := (T2 - t2) / (T1 - t1)
				// x - y = T1 - t1
				// x + y = tt
				x := (tt + T1 - t1) / 2
				y := (tt - T1 + t1) / 2
				if x*2 == tt+T1-t1 && y*2 == tt-T1+t1 {
					if int64(ab^cd) == x^y && simliar(a-1, b-1, c-1, d-1, int(y), int(x)) {
						res[i] = true
					}
				}
			}
		}
	}

	return res[:K]
}

func modAdd(a, b int64) int64 {
	c := a + b
	if c >= MOD {
		c -= MOD
	}
	return c
}

func modSub(a, b int64) int64 {
	c := a - b
	if c < 0 {
		c += MOD
	}
	return c
}
