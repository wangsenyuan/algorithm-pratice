package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	tc := readNum(scanner)
	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N, K, x := line[0], line[1], line[2]
		A := readNNums(scanner, K)
		fmt.Println(solve(N, K, x, A))
	}
}

func solve(N, K, x int, A []int) int64 {
	sort.Ints(A)

	first := int64(A[0])
	last := first + int64(x)
	num := last - int64(N) + 1
	total := (num + last) * int64(N) >> 1

	for i := 0; i < K && int64(A[i]) < num; i, num = i+1, num+1 {
		total -= num - int64(A[i])
	}

	return total
}
