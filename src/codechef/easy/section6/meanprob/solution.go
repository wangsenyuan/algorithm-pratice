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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		q := readNum(scanner)
		X := make([]int, q)

		for i := 0; i < q; i++ {
			X[i] = readNum(scanner)
		}

		Y := solve(n, A, q, X)

		for i := 0; i < q; i++ {
			fmt.Println(Y[i])
		}
	}
}

func solve(n int, A []int, q int, X []int) []int {
	B := make([]int, 1000*n)
	copy(B, A)

	var sum int
	var i int
	for i < n {
		sum += B[i]
		i++
	}

	for i < len(B) {
		B[i] = sum / n
		sum -= B[i-n]
		sum += B[i]
		i++
	}

	Y := make([]int, q)

	for i := 0; i < q; i++ {
		x := X[i] - 1
		if x < len(B) {
			Y[i] = B[x]
		} else {
			Y[i] = B[len(B)-1]
		}
	}

	return Y
}
