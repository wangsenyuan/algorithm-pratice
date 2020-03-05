package main

import (
	"bufio"
	"fmt"
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
	// scanner := bufio.NewScanner(os.Stdin)

	// tc := readNum(scanner)

	// for tc > 0 {
	// 	tc--
	// 	scanner.Scan()
	// 	scanner.Scan()
	// 	s := scanner.Text()
	// 	fmt.Println(solve(s))
	// }
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var n int
		fmt.Scanf("%d", &n)
		var s string
		fmt.Scanf("%s", &s)

		fmt.Println(solve(s))
	}
}

const MOD = 1000000007
const MAX_N = 10000007

var dp []int64

func init() {
	dp = make([]int64, 2*MAX_N+1)
}

func pow(a, b int) int {
	A := int64(a)
	res := int64(1)

	for b > 0 {
		if b&1 == 1 {
			res = (res * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}

	return int(res)
}

func solve(s string) int {
	n := len(s)

	for i := 0; i < 2*n+1; i++ {
		dp[i] = 0
	}

	var offset = n

	var ans int64

	for i := 0; i < n; i++ {
		dp[offset]++
		if s[i] == '(' {
			offset--
		} else if s[i] == ')' {
			offset++
			ans += dp[offset-1] * int64(n-i)
			ans %= MOD
			dp[offset+1] += dp[offset-1]
			dp[offset-1] = 0
		}
	}

	N := int64(n)
	m := int((N * (N + 1) / 2) % MOD)

	ans = (ans % MOD) * int64(pow(m, MOD-2))
	ans %= MOD

	return int(ans)
}

func solve1(s string) int {
	n := len(s)
	stack := make([]int, n)
	var p int

	dp := make([]int64, n+1)

	var res int64

	for i := n - 1; i >= 0; i-- {
		if s[i] == '*' {
			dp[i] = dp[i+1]
		} else if s[i] == '(' {
			if p > 0 {
				// match one )
				p--
			}
			if p > 0 {
				// the first unmatched one )
				dp[i] = dp[stack[p-1]]
			}
		} else {
			dp[i] = int64(n - i)
			if p > 1 {
				dp[i] += dp[stack[p-2]]
				if dp[i] >= MOD {
					dp[i] -= MOD
				}
			}
			stack[p] = i
			p++
		}

		res += dp[i]
		res %= MOD
	}
	N := int64(n)
	m := int((N * (N + 1) / 2) % MOD)
	res = (res * int64(pow(m, MOD-2))) % MOD

	return int(res)
}
