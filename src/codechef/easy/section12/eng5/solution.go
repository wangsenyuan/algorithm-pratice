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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	A := readNNums(scanner, n)
	fmt.Println(solve(n, m, A))
}

func solve(n int, m int, A []int) int {

	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum[i+1][j] = sum[i][j]
		}

		sum[i+1][A[i]-1]++
	}

	total := 1 << uint(m)
	count := make([]int, total)
	dp := make([]int, total)
	for i := 0; i < total; i++ {
		dp[i] = n
	}
	dp[0] = 0

	for mask := 0; mask < total; mask++ {
		if dp[mask] < n {
			a := count[mask]

			for i := 0; i < m; i++ {
				if mask&(1<<uint(i)) == 0 {
					// put type i items at position [a....a + b]
					b := sum[n][i]
					c := sum[a+b][i] - sum[a][i]
					d := b - c
					next := mask | (1 << uint(i))
					if dp[mask]+d < dp[next] {
						dp[next] = dp[mask] + d
						count[next] = a + b
					}
				}
			}
		}
	}

	return dp[total-1]
}
