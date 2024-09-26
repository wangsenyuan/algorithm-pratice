package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([]int, m)

	for i, j := 0, 0; i < m; i++ {
		for j < n && t[i] != s[j] {
			j++
		}
		dp[i] = j
		j++
	}

	// 删掉后缀
	res := n - dp[m-1] - 1

	for i, j := m-1, n-1; i >= 0; i-- {
		for j >= 0 && t[i] != s[j] {
			j--
		}
		if i == 0 {
			res = max(res, j)
		} else if dp[i-1] < j {
			res = max(res, j-dp[i-1]-1)
		}
		j--
	}

	return res
}
