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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		X := make([]int, n)
		Y := make([]int, n)
		for i := 0; i < n; i++ {
			X[i], Y[i] = readTwoNums(scanner)
		}
		fmt.Println(solve(n, X, Y))
		t--
	}
}

func solve(n int, X []int, Y []int) int {
	x := pivot(n, X)
	y := pivot(n, Y)

	return x * y
}

func pivot(n int, X []int) int {
	if n%2 == 1 {
		return 1
	}

	sort.Ints(X)
	half := n / 2
	A, B := X[half-1], X[half]
	return B - A + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
