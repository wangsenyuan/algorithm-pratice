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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		n, k := readTwoNums(scanner)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(scanner, k)
		}
		res := solve(n, k, A)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

func solve(n int, k int, A [][]int) int {
	// can[i][j] => A[i] > A[j] => A[i] can placed on top of A[j]
	can := make([][]bool, n)
	for i := 0; i < n; i++ {
		can[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				can[i][j] = false
			} else {
				can[i][j] = true
				for u := 0; u < k; u++ {
					if A[i][u] <= A[j][u] {
						can[i][j] = false
						break
					}
				}
			}
		}
	}

	vis := make([]int, n)
	pair := make([]int, n)
	for i := 0; i < n; i++ {
		pair[i] = -1
	}

	var dfs func(u int, flag int) bool
	dfs = func(u int, flag int) bool {
		if vis[u] == flag {
			return false
		}
		vis[u] = flag
		for v := 0; v < n; v++ {
			if can[u][v] {
				if pair[v] == -1 || dfs(pair[v], flag) {
					pair[v] = u
					return true
				}
			}
		}
		return false
	}
	var ans int
	for i := 0; i < n; i++ {
		if !dfs(i, i+1) {
			ans++
		}
	}
	return ans
}
