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

func main() {
	// f, err := os.Open("./A-large-practice.in")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		A := readNNums(scanner, 6)
		B := readNNums(scanner, 6)
		x, y := solve(A, B)

		fmt.Printf("Case #%d: %.7f %.7f\n", i, x, y)
	}
}

func solve(A []int, B []int) (float64, float64) {
	a := []complex128{toComplex(A[0], A[1]), toComplex(A[2], A[3]), toComplex(A[4], A[5])}
	b := []complex128{toComplex(B[0], B[1]), toComplex(B[2], B[3]), toComplex(B[4], B[5])}

	ea := a[1] - a[0]
	eb := b[1] - b[0]
	s := eb / ea
	t := b[0] - s*a[0]
	// ans - s * ans == b0 - s * a0
	ans := t / (1.0 - s)
	return real(ans), imag(ans)
}

func toComplex(a, b int) complex128 {
	return complex(float64(a), float64(b))
}
