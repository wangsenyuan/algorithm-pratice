package main

import (
	"bufio"
	"bytes"
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

func solve(s string, t string) string {
	n := len(s)
	m := len(t)
	dp := make([]int, n+1)
	dp[n] = m
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	for i, j := m-1, n-1; i >= 0 && j >= 0; i-- {
		for j >= 0 && t[i] != s[j] {
			j--
		}
		if j < 0 {
			break
		}
		dp[j] = i
		j--
	}
	for j := n - 1; j >= 0; j-- {
		if dp[j] < 0 {
			dp[j] = dp[j+1]
		}
	}
	if dp[0] == 0 {
		return t
	}
	// dp[0] < 0 不代表没有答案， 只是表示后缀不能匹配
	ans := []int{-1, -1}

	if dp[0] > 0 {
		// 把dp[0]前面的删除
		ans[0] = 0
		ans[1] = dp[0] - 1
	}
	for i, j := 0, 0; i < m; i++ {
		for j < n && t[i] != s[j] {
			j++
		}
		if j == n {
			// not found
			break
		}
		r := max(i+1, dp[j+1]-1)
		if ans[0] < 0 || r-i < ans[1]-ans[0]+1 {
			ans[0] = i + 1
			ans[1] = r
		}
		j++
	}

	if ans[0] < 0 || ans[1]-ans[0]+1 == m {
		return "-"
	}
	var buf bytes.Buffer
	if ans[0] > 0 {
		buf.WriteString(t[:ans[0]])
	}
	if ans[1]+1 < m {
		buf.WriteString(t[ans[1]+1:])
	}

	return buf.String()
}
