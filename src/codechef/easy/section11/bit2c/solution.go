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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(s))
	}
}

func solve(s string) int {
	nums := make([]int, 0, 11)
	ops := make([]byte, 0, 11)

	var num int

	for i := 0; i <= len(s); i++ {
		if i == len(s) || !isDigit(s[i]) {
			nums = append(nums, num)
			if i < len(s) {
				ops = append(ops, s[i])
			}
			num = 0
			continue
		}
		x := int(s[i] - '0')
		num = num*10 + x
	}

	n := len(nums)
	dp := make([][]map[int]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]map[int]bool, n)
	}

	var dfs func(i, j int) map[int]bool

	dfs = func(i, j int) map[int]bool {
		if dp[i][j] != nil {
			return dp[i][j]
		}
		dp[i][j] = make(map[int]bool)
		if i == j {
			dp[i][j][nums[i]] = true
			return dp[i][j]
		}

		for k := i; k < j; k++ {
			a := dfs(i, k)
			b := dfs(k+1, j)
			for aa := range a {
				for bb := range b {
					cc := bitOp(aa, bb, ops[k])
					dp[i][j][cc] = true
				}
			}
		}
		return dp[i][j]
	}

	res := dfs(0, n-1)

	var ans int

	for k := range res {
		ans = max(ans, k)
	}

	return ans
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}

func bitOp(a int, b int, c byte) int {
	if c == '&' {
		return a & b
	}
	if c == '|' {
		return a | b
	}
	return a ^ b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
