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
	n, k := readTwoNums(scanner)
	arr := readNNums(scanner, n)
	fmt.Println(solve(n, k, arr))
}

const MOD = 1000000007

func solve(n, K int, A []int) int64 {
	// sort.Ints(A)
	cnt := make([]int64, 8000)
	var ps []int

	for i := 0; i < n; i++ {
		cnt[A[i]]++
		if cnt[A[i]] == 1 {
			ps = append(ps, A[i])
		}
	}
	// sort.Ints(ps)
	m := len(ps)

	K = min(m, K)

	dp := make([][]int64, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int64, K+1)
	}

	dp[0][0] = 1

	for i := 1; i <= m; i++ {
		dp[i][0] = 1
		x := int64(cnt[ps[i-1]])
		// append ps[i] into the set
		for k := 1; k <= K; k++ {
			dp[i][k] = dp[i-1][k] + dp[i-1][k-1]*x
			dp[i][k] %= MOD
		}
	}

	var res int64
	for j := 0; j <= K; j++ {
		res += dp[m][j]
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
