package main

import (
	"bufio"
	"fmt"
	"os"
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
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		var N, B int64
		scanner.Scan()
		pos := readInt64(scanner.Bytes(), 0, &N)
		readInt64(scanner.Bytes(), pos+1, &B)
		fmt.Println(solve(N, B))
	}
}

func solve(N, B int64) int64 {
	if N%B == 0 {
		return N - N/B + 1
	}
	return N - N/B
}

func solve1(N, B int64) int64 {

	check := func(A int64) bool {
		cnt := A

		for A > 1 && cnt < N {
			x := A / B
			cnt += x
			A = x + (A - x*B)
		}
		return cnt >= N
	}

	left, right := int64(1), N

	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
