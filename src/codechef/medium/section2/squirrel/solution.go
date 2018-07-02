package main

import (
	"bufio"
	"fmt"
	"math"
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
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		line := readNNums(scanner, 3)
		m, n, k := line[0], line[1], line[2]
		T := readNNums(scanner, m)
		P := readNNums(scanner, m)
		fmt.Println(solve(n, m, k, T, P))
		tc--
	}
}

func solve(n int, m int, k int, T []int, P []int) int {
	n = min(n, m)
	nuts := make([]int, m)

	check := func(tt int) bool {
		for i := 0; i < m; i++ {
			if tt >= T[i] {
				nuts[i] = (tt-T[i])/P[i] + 1
			} else {
				nuts[i] = 0
			}
		}
		// log m
		nthLargest(nuts, n)
		var total int
		for i := 0; i < n; i++ {
			total += nuts[i]
		}
		return total >= k
	}

	left, right := 0, math.MaxInt32

	for left < right {
		mid := left + (right-left)>>1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func nthLargest(xs []int, K int) {
	if K == 0 {
		return
	}

	var loop func(start, end int)
	loop = func(start, end int) {
		if start >= end {
			return
		}
		pivot := xs[end]
		i := start - 1

		for j := start; j <= end-1; j++ {
			if xs[j] >= pivot {
				i++
				xs[i], xs[j] = xs[j], xs[i]
			}
		}
		xs[i+1], xs[end] = xs[end], xs[i+1]

		// all left before (including) pos is greater or equal to pivot
		pos := i + 1

		if K == pos+1 {
			return
		}
		if K < pos+1 {
			loop(start, pos-1)
		} else {
			loop(pos+1, end)
		}
	}

	loop(0, len(xs)-1)
}
