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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m int
	pos := readInt(scanner.Bytes(), 0, &n)
	pos = readInt(scanner.Bytes(), pos+1, &m)
	var D int64
	readInt64(scanner.Bytes(), pos+1, &D)
	buy, sell := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		buy[i] = make([]int, n)
		sell[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		tmp := readNNums(scanner, 2*m)
		for j := 0; j < m; j++ {
			sell[j][i] = tmp[2*j]
			buy[j][i] = tmp[2*j+1]
		}
	}
	res := solve(n, m, buy, sell, D)
	if res > 1e18 {
		fmt.Println("Quintillionnaire")
	} else {
		fmt.Printf("%.7f\n", res)
	}
}

func solve(n, m int, buy, sell [][]int, D int64) float64 {
	// dp[t][i][c] means the max money at the end of time t, at kiosk i, having currency c,
	dp := make([][][]float64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, 2)
			// 1 means dollar, while 0 means peper corns
		}
	}

	for i := 0; i < n; i++ {
		dp[0][i][0] = float64(D)
		dp[0][i][1] = 0
	}

	for t := 1; t <= m; t++ {
		for i := 0; i < n; i++ {
			// case one, no exchange at time t
			dp[t][i][0] = dp[t-1][i][0]
			dp[t][i][1] = dp[t-1][i][1]

			if i+1 < n {
				dp[t][i][0] = max(dp[t][i][0], dp[t-1][i+1][0])
				dp[t][i][1] = max(dp[t][i][1], dp[t-1][i+1][1])
			}
			if i-1 >= 0 {
				dp[t][i][0] = max(dp[t][i][0], dp[t-1][i-1][0])
				dp[t][i][1] = max(dp[t][i][1], dp[t-1][i-1][1])
			}

			dp[t][i][0] = max(dp[t][i][0], dp[t-1][i][1]*float64(buy[t-1][i]))
			dp[t][i][1] = max(dp[t][i][1], dp[t-1][i][0]/float64(sell[t-1][i]))
		}
	}

	var ans float64
	for i := 0; i < n; i++ {
		if ans < dp[m][i][0] {
			ans = dp[m][i][0]
		}
	}
	return ans
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
