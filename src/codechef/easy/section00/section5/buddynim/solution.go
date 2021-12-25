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
		n, m := readTwoNums(scanner)
		A := readNNums(scanner, n)
		B := readNNums(scanner, m)
		res := solve(A, B)
		if res {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}
	}
}

func solve(A []int, B []int) bool {
	var n int
	for j := 0; j < len(A); j++ {
		if A[j] == 0 {
			continue
		}
		A[n] = A[j]
		n++
	}
	sort.Ints(A[:n])
	var m int
	for i := 0; i < len(B); i++ {
		if B[i] == 0 {
			continue
		}
		B[m] = B[i]
		m++
	}
	sort.Ints(B[:m])
	if m != n {
		return true
	}
	for i := 0; i < m; i++ {
		if A[i] != B[i] {
			return true
		}
	}
	return false
}
