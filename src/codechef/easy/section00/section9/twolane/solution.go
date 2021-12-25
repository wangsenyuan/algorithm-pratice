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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N := line[0]
		K := line[1]
		D := line[2]
		X := readNNums(scanner, N)
		L := readNNums(scanner, N)
		fmt.Println(solve(N, K, D, X, L))
	}
}

func solve(N, K, D int, X []int, L []int) int {

	lane := 3 - L[0]
	var lastSwitch = -(1 << 30)
	for i := 0; i < N; i++ {
		if lane != L[i] {
			continue
		}
		best := max(lastSwitch+D, X[i-1]+1)
		if best >= X[i] {
			return X[i]
		}
		lastSwitch = best
		lane = 3 - L[i]
	}

	return K
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(N, K, D int, X []int, L []int) int {
	A := make([]int, 0, N)
	B := make([]int, 0, N)

	for i := 0; i < N; i++ {
		if L[i] == 1 {
			A = append(A, X[i])
		} else {
			B = append(B, X[i])
		}
	}

	reach := func(dd, a, b int) int {
		if dd >= a || b+1 == a {
			// block by a
			return a
		}
		// dd < a
		if dd > b {
			return dd
		}
		// dd <= b
		return b + 1
	}

	var cur int
	var i, j int

	if A[0] < B[0] {
		if A[0]+1 == B[0] {
			return B[0]
		}
		cur = A[0] + 1
		i++
	} else {
		if B[0]+1 == A[0] {
			return A[0]
		}
		cur = B[0] + 1
		j++
	}

	for i < len(A) && j < len(B) {
		a := A[i]
		b := B[j]
		dd := cur + D

		if a > b {
			cur = reach(dd, a, b)
			if cur == a {
				return a
			}
			j++
		} else {
			// a < b
			cur = reach(dd, b, a)
			if cur == b {
				return b
			}
			i++
		}
	}

	return K
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
