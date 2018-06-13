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
		n, m := readTwoNums(scanner)

		prices := make([][]int, n)
		for i := 0; i < n; i++ {
			prices[i] = readNNums(scanner, m)
		}

		discounts := make([][]int, n)
		for i := 0; i < n; i++ {
			discounts[i] = readNNums(scanner, m)
		}
		fmt.Println(solve(n, m, prices, discounts))

		t--
	}
}

func solve(n int, m int, prices [][]int, discounts [][]int) int64 {
	// f[i][j] = buy items [1..i] ending at store j
	f := make([][]int64, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			f[i][j] = -1
		}
	}
	g := make([]int64, n)
	for i := 0; i < n; i++ {
		g[i] = math.MaxInt64
	}
	// f[i][j] = min (f[i-1][j] + prices[i][j] - discount[i][j], f[i-1][k] + prices[i][j])

	for j := 0; j < m; j++ {
		if int64(prices[0][j]) < g[0] {
			g[0] = int64(prices[0][j])
		}
		f[0][j] = int64(prices[0][j])
	}

	for i := 1; i < n; i++ {
		g[i] = math.MaxInt64
		for j := 0; j < m; j++ {
			a := f[i-1][j]
			if prices[i][j] >= discounts[i-1][j] {
				a += int64(prices[i][j]) - int64(discounts[i-1][j])
			}
			b := g[i-1] + int64(prices[i][j])
			f[i][j] = min(a, b)
			g[i] = min(g[i], f[i][j])
		}
	}

	return g[n-1]
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
