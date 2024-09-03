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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = readString(reader)
	}
	res := solve(nums, m)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(res)
	}
}

var digits []string

func init() {
	digits = []string{
		"1110111", "0010010", "1011101", "1011011", "0111010", "1101011", "1101111", "1010010", "1111111", "1111011",
	}

}

func solve(nums []string, k int) string {
	n := len(nums)

	dp := make([][]byte, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]byte, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = ' '
		}
	}
	dp[n][0] = '1'

	check := func(a, b string) (bool, int) {
		var diff int

		for i := 0; i < 7; i++ {
			if a[i] == '1' {
				if b[i] == '0' {
					diff++
				}
			} else {
				if b[i] == '1' {
					return false, -1
				}
			}
		}

		return true, diff
	}

	res := make([]byte, n)

	var dfs func(i int, j int) bool

	dfs = func(i int, j int) bool {
		if i == n {
			return j == 0
		}
		if dp[i][j] != ' ' {
			return dp[i][j] == '1'
		}
		dp[i][j] = '0'
		cur := nums[i]
		for u := 9; u >= 0; u-- {
			ok, diff := check(digits[u], cur)

			if ok && j >= diff {
				res[i] = byte(u + '0')
				if dfs(i+1, j-diff) {
					dp[i][j] = '1'
					return true
				}
			}
		}
		return false
	}

	if !dfs(0, k) {
		return ""
	}

	return string(res)
}
