package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
		t--
	}
}

func solve(n int, A []int) int64 {
	x, y := int64(A[0]), int64(A[0])

	for i := 1; i < n; i++ {
		a := int64(A[i])
		if a >= 0 {
			x = max(x+a, x-a, x*a)
			y = min(y+a, y-a, y*a)
		} else {
			z := max(x+a, x-a, y*a)
			y = min(y+a, y-a, x*a)
			x = z
		}
	}

	return y
}

func min(a, b, c int64) int64 {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func max(a, b, c int64) int64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
