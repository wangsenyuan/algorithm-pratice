package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	f, err := os.Open("./A-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		N := readNum(scanner)
		H := make([]int, N)
		W := make([]int, N)
		X := make([]string, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			bs := scanner.Bytes()

			pos := readInt(bs, 0, &H[j])
			pos = readInt(bs, pos+1, &W[j])
			X[j] = string(bs[pos+1:])
		}
		M := readNum(scanner)

		Q := make([][]int, M)
		for j := 0; j < M; j++ {
			Q[j] = readNNums(scanner, 2)
		}
		res := solve(N, H, W, X, M, Q)
		fmt.Printf("Case #%d:\n", i)
		for _, ans := range res {
			fmt.Println(ans)
		}
	}
}

func solve(N int, H []int, W []int, S []string, M int, Q [][]int) []string {
	//(A, B) is the bottom left corner of the bird,
	//(C,, D) is the top right corner of the bird
	A, B := math.MaxInt32, math.MaxInt32
	C, D := math.MinInt32, math.MinInt32

	for i := 0; i < N; i++ {
		if S[i] == "BIRD" {
			A = min(A, H[i])
			B = min(B, W[i])
			C = max(C, H[i])
			D = max(D, W[i])
		}
	}

	check1 := func(X, Y int) string {
		for i := 0; i < N; i++ {
			if H[i] == X && W[i] == Y {
				return "NOT BIRD"
			}
		}
		return "UNKNOWN"
	}

	check := func(X, Y int) string {
		// pretend it is bird
		if A == math.MaxInt32 {
			// no bird range
			return check1(X, Y)
		}

		if X >= A && X <= C && Y >= B && Y <= D {
			// in the range
			return "BIRD"
		}
		AA := min(A, X)
		BB := min(B, Y)
		CC := max(C, X)
		DD := max(D, Y)
		for i := 0; i < N; i++ {
			if S[i] == "NOT BIRD" {
				h, w := H[i], W[i]
				if h >= AA && h <= CC && w >= BB && w <= DD {
					return "NOT BIRD"
				}
			}
		}
		return "UNKNOWN"
	}

	res := make([]string, M)

	for i := 0; i < M; i++ {
		h, w := Q[i][0], Q[i][1]
		res[i] = check(h, w)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
