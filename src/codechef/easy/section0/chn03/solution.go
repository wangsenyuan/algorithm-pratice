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

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readNum(scanner *bufio.Scanner) (a int) {
	tmp := readNNums(scanner, 1)
	a = tmp[0]
	return
}

func readThree(scanner *bufio.Scanner) (a int, b int, c int) {
	tmp := readNNums(scanner, 3)
	a, b, c = tmp[0], tmp[1], tmp[2]
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, k, time := readThree(scanner)
		tm := make([][]int, n)
		for i := 0; i < n; i++ {
			tm[i] = readNNums(scanner, 3)
		}
		pm := make([][]int, n)
		for i := 0; i < n; i++ {
			pm[i] = readNNums(scanner, 3)
		}
		fmt.Println(solve(n, k, time, tm, pm))
		t--
	}
}

const oo = 1 << 30

func solve(n, k, time int, tm [][]int, pm [][]int) int {
	// dp[i][swap][used][time] means the max pleasure got before contest i, using 'swap' swaps
	// took 'used' contest, and at time 'time'
	dp := make([][][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([][]int, n+1)
			for u := 0; u <= n; u++ {
				dp[i][j][u] = make([]int, time+1)
				for v := 0; v <= time; v++ {
					dp[i][j][u][v] = -1
				}
			}
		}
	}

	var fn func(i, s, u, t int) int
	fn = func(i, s, u, t int) int {
		if s > k {
			return -oo
		}
		if s+u > n {
			// if swaps + used exceed total n, no valid
			return -oo
		}
		if t > time {
			return -oo
		}
		if i == n {
			return 0
		}

		if dp[i][s][u][t] < 0 {
			// skip contest i
			ret := fn(i+1, s, u, t)

			for j := 0; j < 3; j++ {
				// take j (easy, medium, hard) problem of contest i
				ret = max(ret, pm[i][j]+fn(i+1, s, u+1, t+tm[i][j]))
				for jj := j + 1; jj < 3; jj++ {
					// swap problem jj with another contest problem
					ret = max(ret, pm[i][j]+pm[i][jj]+fn(i+1, s+1, u+1, t+tm[i][j]+tm[i][jj]))
				}
			}
			// take all three problems, means 2 swaps needed
			ret = max(ret, pm[i][0]+pm[i][1]+pm[i][2]+fn(i+1, s+2, u+1, t+tm[i][0]+tm[i][1]+tm[i][2]))
			dp[i][s][u][t] = ret
		}

		return dp[i][s][u][t]
	}
	return fn(0, 0, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
