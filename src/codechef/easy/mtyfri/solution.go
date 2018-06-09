package main

import (
	"sort"
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

	for t > 0 {
		n, K := readTwoNums(scanner)
		A := readNNums(scanner, n)
		res := solve(n, K, A)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}

func solve(n int, K int, A []int) bool {
	m := (n + 1) / 2

	B := make([]int, m)
	C := make([]int, n-m)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			B[i/2] = A[i]
		} else {
			C[i/2] = A[i]
		}
	}

	sort.Ints(B)

	// reverse B
	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		B[i], B[j] = B[j], B[i]
	}

	sort.Ints(C)

	// len(C) <= len(B)
	for i, k := 0, 0; i < n-m && k < K; i++ {
		if C[i] < B[i] {
			// swap them
			C[i], B[i] = B[i], C[i]
			k++
		}
	}

	sb := sum(B)
	sc := sum(C)
	return sb < sc
}

func sum(A []int) int64 {
	var res int64
	for i := 0; i < len(A); i++ {
		res += int64(A[i])
	}
	return res
}
