package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		pairs := make([][]int, n)
		for i := 0; i < n; i++ {
			pairs[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, pairs))
		tc--
	}
}

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

func solve(n int, pairs [][]int) int {
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = math.MaxInt32
		dp[i][1] = math.MinInt32
	}

	var loop func(node int) bool
	var ans int
	loop = func(node int) bool {
		pair := pairs[node-1]
		if pair[0] == -1 {
			dp[node][0] = pair[1]
			dp[node][1] = pair[1]
			return true
		}
		if !loop(pair[0]) {
			return false
		}
		if !loop(pair[1]) {
			return false
		}
		a, b := dp[pair[0]][0], dp[pair[0]][1]
		c, d := dp[pair[1]][0], dp[pair[1]][1]
		if b > c && d > a {
			// overlap
			return false
		}
		dp[node][0] = min(a, c)
		dp[node][1] = max(b, d)
		if a > d {
			ans++
		}
		return true
	}
	can := loop(1)
	if !can {
		ans = -1
	}
	return ans
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
