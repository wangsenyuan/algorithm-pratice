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

	n, m := readTwoNums(scanner)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = readNNums(scanner, m)
	}
	fmt.Println(solve(n, m, A))
}

const MOD = 1e9 + 7
const MAX_M = 1e5 + 1

func solve(n, m int, A [][]int) int64 {
	cnt := make([][]int64, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int64, MAX_M)
	}
	var mx int
	for i := 0; i < n; i++ {
		for _, a := range A[i] {
			cnt[i][a]++
			if a > mx {
				mx = a
			}
		}
	}
	var ans int64
	dp := make([]int64, mx+1)
	curCnt := make([]int64, n)
	for i := mx; i > 0; i-- {
		for l := 0; l < n; l++ {
			curCnt[l] = 0
			for j := i; j <= mx; j += i {
				curCnt[l] += cnt[l][j]
			}
		}
		dp[i] = 1
		for j := 0; j < n; j++ {
			dp[i] = (dp[i] * (curCnt[j] + 1)) % MOD
		}
		// subtract only choose one row
		for j := 0; j < n; j++ {
			dp[i] = (dp[i] + MOD - curCnt[j]) % MOD
		}
		// substruct choose no row
		dp[i] = (dp[i] + MOD - 1) % MOD

		for j := 2 * i; j <= mx; j += i {
			dp[i] = (dp[i] + MOD - dp[j]) % MOD
		}
		ans = (ans + dp[i]*int64(i)%MOD) % MOD
	}
	return ans
}
