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

	for t > 0 {
		n, m := readTwoNums(scanner)
		fmt.Println(solve(n, m))
		t--
	}
}

func solve(n int, m int) int {
	max := (n + 3) * n / 2

	mins := make([]int, n+1)

	var calMin func(n int) int

	calMin = func(n int) int {

		if n == 1 {
			return 2
		}

		if n == 2 {
			return 5
		}

		if mins[n] > 0 {
			return mins[n]
		}

		mins[n] = calMin(n/2) + calMin(n-1-n/2) + n + 1
		return mins[n]
	}

	min := calMin(n)

	if m >= min && m <= max {
		return 0
	}
	if m < min {
		return -1
	}

	return m - max
}
