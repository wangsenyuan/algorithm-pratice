package main

import (
	"bufio"
	"fmt"
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

	for i := 0; i < t; i++ {
		n, m := readTwoNums(scanner)
		A := readNNums(scanner, n)
		B := readNNums(scanner, m)
		fmt.Println(solve(n, A, m, B))
	}
}

func solve(n int, A []int, m int, B []int) int64 {
	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(A[i])
	}

	dp := make([]int64, n)

	for k, i := B[m-1], 0; i < n && i+k <= n; i++ {
		dp[i] = sum[i+k] - sum[i]
	}

	que := make([]int, n)
	for k := m - 2; k >= 0; k-- {
		head, tail := 0, 0
		b := B[k]
		a := B[k] - B[k+1] - 1
		for i, j := 0, 1; i < n; i++ {
			if i+b <= n {
				for j <= i+a {
					// j is in the range of k + 1
					for tail > head && dp[que[tail-1]] < dp[j] {
						tail--
					}
					que[tail] = j
					tail++
					j++
				}
				if que[head] == i {
					head++
				}
				dp[i] = sum[i+b] - sum[i] - dp[que[head]]
			} else {
				dp[i] = 0
			}
		}
	}

	var ans int64
	for i := 0; i < n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
