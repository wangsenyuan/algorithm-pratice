package main

import (
	"bufio"
	"os"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		scanner.Scan()
		A := scanner.Bytes()
		scanner.Scan()
		B := scanner.Bytes()
		res := solve(len(A), A, B)
		fmt.Println(res)
	}

}

func solve(n int, A, B []byte) int {
	C := make([]byte, 2*n)
	copy(C, B)
	copy(C[n:], B)

	lps := computeLPS(A)

	from := 0
	best := 0
	for i, j := 0, 0; i < len(C); {
		if A[j] == C[i] {
			i++
			j++
			if j > best {
				best = j
				from = i - j
			}
		}

		if j == n {
			break
		} else if i < len(C) && A[j] != C[i] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}

	}

	return from
}

func computeLPS(pat []byte) []int {
	n := len(pat)
	lps := make([]int, n)

	L := 0
	i := 1
	for i < n {
		if pat[i] == pat[L] {
			L++
			lps[i] = L
			i++
		} else {
			if L > 0 {
				L = lps[L-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}
