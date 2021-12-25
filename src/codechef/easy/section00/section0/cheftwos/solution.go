package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		a := scanner.Bytes()
		scanner.Scan()
		b := scanner.Bytes()
		fmt.Println(solve(len(a), a, b))
	}
}

const MOD = 1000000007

func solve(n int, a, b []byte) int {
	if n == 1 {
		return 2
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 3)
		dp[i][0] = make([]int, 3)
		dp[i][1] = make([]int, 3)
		dp[i][2] = make([]int, 3)
	}

	dp[0][a[0]-'1'][b[0]-'1'] = 1
	dp[0][b[0]-'1'][a[0]-'1'] = 1
	for i := 1; i < n-1; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				var p, q int
				// not swap
				if getNext(j, a[i], &p) && getNext(k, b[i], &q) {
					dp[i][p][q] = (dp[i][p][q] + dp[i-1][j][k]) % MOD
				}
				// swap
				if getNext(j, b[i], &p) && getNext(k, a[i], &q) {
					dp[i][p][q] = (dp[i][p][q] + dp[i-1][j][k]) % MOD
				}
			}
		}
	}

	// multiple 2 is for the last digit
	return (2 * dp[n-2][0][0]) % MOD
}

func getNext(twos int, digit byte, nextTwos *int) bool {
	if digit == '1' {
		*nextTwos = 0
		// 21 is invalid
		return twos != 1
	}
	*nextTwos = twos + 1
	// 222 is invalid
	return twos != 2
}
