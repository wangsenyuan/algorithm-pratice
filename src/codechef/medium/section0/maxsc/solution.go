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

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		matrix := make([][]int, n)
		for j := 0; j < n; j++ {
			matrix[j] = readNNums(scanner, n)
		}
		fmt.Println(solve(n, matrix))
	}
}

func solve(n int, matrix [][]int) int64 {
	for i := 0; i < n; i++ {
		sort.Ints(matrix[i])
	}
	last := matrix[n-1][n-1]
	ans := int64(last)

	for i := n - 2; i >= 0; i-- {
		row := matrix[i]
		j := sort.SearchInts(row, last)
		if j == 0 {
			// all is greater or equal to last
			return -1
		}

		last = row[j-1]
		ans += int64(last)
	}

	return ans
}
