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
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, k, A))
	}
}

func solve(n, k int, A []int) int {
	// x[i] = 1 << i
	x := make([]int, k)

	for _, a := range A {
		for i := 0; i < k; i++ {
			if a == 1<<uint(i) {
				x[i] = 1
				break
			}
		}
	}

	var s int
	for i := 0; i < k; i++ {
		s += x[i]
	}
	return k - s
}
