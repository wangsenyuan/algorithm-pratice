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

	for t > 0 {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		fmt.Printf("%d", res[0])
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
		t--
	}
}

func solve(n int, A []int) []int {
	left := voteBehind(n, A)
	// B is reverse of A
	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[n-1-i] = A[i]
	}
	right := voteBehind(n, B)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		right[i], right[j] = right[j], right[i]
	}

	vote := make([]int, n)

	var tmp int
	for i := 0; i < n; i++ {
		tmp += left[i]
		vote[i] = tmp
	}

	tmp = 0
	for i := n - 1; i >= 0; i-- {
		tmp += right[i]
		vote[i] += tmp
	}
	return vote
}

func voteBehind(n int, A []int) []int {
	sum := make([]int64, n)

	sum[0] = int64(A[0])

	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + int64(A[i])
	}

	left := make([]int, n)

	for i := 0; i < n-1; i++ {
		left[i+1]++
		j := sort.Search(n, func(j int) bool {
			return sum[j]-sum[i] > int64(A[i])
		})
		if j < n-1 {
			// will not vote after j
			left[j+1]--
		}
	}
	return left
}
