package main

import (
	"bufio"
	"os"
	"fmt"
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

		scanner.Scan()
		s := scanner.Bytes()
		nums := readNNums(scanner, n)
		fmt.Println(solve(n, s, nums))
		t--
	}
}

func solve(n int, s []byte, nums []int) int64 {
	x := make([]byte, n+1)
	x[0] = '$'
	copy(x[1:], s)
	stack1 := make([]int, n+1)
	p1 := 1
	hsh := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if isClosing(x[i]) && stack1[p1-1] > 0 && isPaired(x[stack1[p1-1]], x[i]) {
			hsh[i] = stack1[p1-1]
			p1--
		} else {
			stack1[p1] = i
			p1++
		}
	}

	dp := make([]int64, n+1)
	var best int64
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int64(nums[i-1])
		j := hsh[i]
		if j > 0 {
			dp[i] = max(dp[i], dp[j-1]+sum[i]-sum[j-1])
		}

		best = max(best, dp[i])
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func isClosing(symbol byte) bool {
	return symbol == ')' || symbol == '}' || symbol == ']' || symbol == '>'
}

func isPaired(left, right byte) bool {
	return left == '(' && right == ')' ||
		left == '{' && right == '}' ||
		left == '[' && right == ']' ||
		left == '<' && right == '>'

}
