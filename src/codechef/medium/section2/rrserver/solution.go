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

	n, m := readTwoNums(scanner)
	pairs := make([][]int, m)

	for i := 0; i < m; i++ {
		pairs[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, m, pairs))
}

func solve(n int, m int, pairs [][]int) int {
	N := 1 << uint(n)

	table := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		table[i] = make(map[int]bool)
	}

	for _, pair := range pairs {
		a, b := pair[0]-1, pair[1]-1
		table[a][b] = true
		table[b][a] = true
	}

	dp1 := make([]int, N)
	dp2 := make([]int, N)
	for i := 1; i < N; i++ {
		dp1[i] = math.MaxInt32
	}

	for state := 0; state < N; state++ {
		var cnt int
		for i := 0; i < n; i++ {
			if state&(1<<uint(i)) > 0 {
				for j := range table[i] {
					if state&(1<<uint(j)) == 0 {
						cnt++
					}
				}
			}
		}
		dp2[state] = cnt
	}

	for state := 0; state < N; state++ {
		for i := 0; i < n; i++ {
			if state&(1<<uint(i)) == 0 {
				// try to add i
				next := state | (1 << uint(i))
				dp1[next] = min(dp1[next], dp1[state]+dp2[state])
			}
		}
	}

	return dp1[N-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
