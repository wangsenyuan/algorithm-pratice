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
	for i := 0; i < t; i++ {
		n, m := readTwoNums(scanner)
		fmt.Println(solve(n, m))
	}
}

func solve(n, m int) int {
	if n > m {
		return solve(m, n)
	}

	if n == 1 {
		return m
	}
	if n == 2 {
		return solve2(m)
	}
	return solve3(m)
}

func solve2(m int) int {
	x := m / 6
	y := m % 6
	return 4*x + min(y, 2)*2
}

func solve3(m int) int {
	x := m / 6
	y := m % 2
	if m > 12 && y == 2 {
		return 4*x + 3
	}
	return 4*x + min(y, 2)*2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
