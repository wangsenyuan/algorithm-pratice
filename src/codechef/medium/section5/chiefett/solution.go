package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n, k := readTwoNums(scanner)
		fillNNums(scanner, n, Prices)
		fillNNums(scanner, k, Discount)
		res := solve(n, k, Prices, Discount)
		fmt.Printf("%.5f\n", res)
	}
}

const MAX_N = 1010

var dp [][]float64

var Prices []int
var Discount []int

func init() {
	dp = make([][]float64, MAX_N)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]float64, MAX_N)
	}
	Prices = make([]int, MAX_N)
	Discount = make([]int, MAX_N)
}

func solve(n, k int, prices []int, discount []int) float64 {
	sort.Ints(prices)
	sort.Ints(discount)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i && j <= k; j++ {
			p := float64(j) / float64(i)
			dp[i][j] = (float64(prices[i-1])*float64(discount[j-1])/100.0+dp[i-1][j-1])*p + dp[i-1][j]*(1.0-p)
		}
	}

	return dp[n][k]
}
