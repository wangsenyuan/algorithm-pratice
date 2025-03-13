package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	if len(res) > 0 {
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) []int {
	// 假设删除了i
	// 那么 s[:i] = t[:i], 且 s[i+1:] = t[i:]
	n := len(s)
	// dp[i] = true 表示 s[i:] = t[i-1:]?
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n - 1; i > 0; i-- {
		if s[i] == t[i-1] {
			dp[i] = true
		}
		if !dp[i] {
			break
		}
	}
	var res []int

	for i := range n {
		if dp[i+1] {
			res = append(res, i+1)
		}
		if i == n-1 || s[i] != t[i] {
			break
		}
	}

	return res
}
