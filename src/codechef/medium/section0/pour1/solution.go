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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		A := readNum(scanner)
		B := readNum(scanner)
		C := readNum(scanner)
		fmt.Println(solve(A, B, C))
		t--
	}
}

func solve(A, B, C int) int {
	if A < C && B < C {
		return -1
	}

	if A == C || B == C {
		return 1
	}

	if A == 0 || B == 0 || A == B {
		return -1
	}
	res1 := pour(A, B, C)
	res2 := pour(B, A, C)

	if res1 < 0 {
		return res2
	}
	if res2 < 0 {
		return res1
	}
	return min(res1, res2)
}

func pour(A, B, C int) int {
	// from A to B
	var steps int
	var a, b int
	for {
		if a == 0 {
			a = A
			steps++
		}
		for b < B {
			x := min(B, a+b)
			a = a + b - x
			b = x
			steps++
			if b == C || a == C {
				return steps
			}
			if a == 0 {
				steps++
				a = A
			}
		}
		if a == A && b == B {
			break
		}
		if b == B {
			steps++
			b = 0
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
