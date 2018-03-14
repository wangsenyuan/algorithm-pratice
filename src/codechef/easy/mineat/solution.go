package main

import (
	"bufio"
	"fmt"
	"os"
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
		n, h := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, h, A))
		t--
	}
}

func solve(n int, h int, A []int) int {

	check := func(k int) bool {
		var took int
		var i int
		for i < n && took < h {
			tmp := A[i] / k
			if tmp*k < A[i] {
				tmp++
			}
			took += tmp
			i++
		}

		return i == n && took <= h
	}

	max := A[0]

	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	left, right := 1, max

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
