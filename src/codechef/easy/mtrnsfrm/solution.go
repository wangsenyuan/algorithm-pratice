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
		n, m := readTwoNums(scanner)
		A, B := make([][]int, n), make([][]int, n)
		for j := 0; j < n; j++ {
			A[j] = readNNums(scanner, m)
		}
		for j := 0; j < n; j++ {
			B[j] = readNNums(scanner, m)
		}
		fmt.Println(solve(n, m, A, B))
	}
}

func solve(n, m int, A, B [][]int) int64 {
	AB := make([][]int, n)
	for i := 0; i < n; i++ {
		AB[i] = make([]int, m)
		for j := 0; j < m; j++ {
			AB[i][j] = A[i][j] - B[i][j]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if AB[i][j]-AB[0][j] != AB[i][0]-AB[0][0] {
				return -1
			}
		}
	}

	ps := make([]int, 0, n+m)
	for i := 0; i < n; i++ {
		ps = append(ps, AB[0][0]-AB[i][0])
	}

	for i := 0; i < m; i++ {
		ps = append(ps, AB[0][i])
	}

	sort.Ints(ps)
	mid := (n + m) / 2
	var ans int64

	for _, x := range ps {
		ans += abs(int64(ps[mid] - x))
	}

	return ans
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
