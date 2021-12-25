package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readPair(scanner *bufio.Scanner) (a int, b int) {
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
		n, k := readPair(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A, k))
	}
}

func solve(n int, A []int, k int) int64 {
	// fmt.Printf("[debug]%d %v %d\n", n, A, k)
	if k == 1 {
		return maxSum(A)
	}

	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(A[i])
	}

	B := make([]int, 2*n)
	copy(B, A)
	copy(B[n:], A)

	tmp := maxSum(B)

	if sum > 0 {
		return tmp + sum*int64(k-2)
	}

	return tmp
}

func maxSum(A []int) int64 {
	ans := int64(math.MinInt64)
	var tmp int64
	for i := 0; i < len(A); i++ {
		tmp += int64(A[i])
		if tmp > ans {
			ans = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return ans
}
