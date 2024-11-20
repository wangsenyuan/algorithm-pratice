package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(s string) []string {
	n := len(s)

	dp := make([]bool, n)
	fp := make([]bool, n)

	pos := make(map[string]int)

	for i := n - 2; i > 4; i-- {
		if i+2 == n || fp[i+2] || i+4 <= n && dp[i+2] && s[i:i+2] != s[i+2:i+4] {
			dp[i] = true
			pos[s[i:i+2]] = i
		}
		if i+3 == n || i+3 <= n && dp[i+3] || i+6 <= n && fp[i+3] && s[i:i+3] != s[i+3:i+6] {
			fp[i] = true
			pos[s[i:i+3]] = i
		}
	}

	var res []string
	for s := range pos {
		res = append(res, s)
	}

	sort.Strings(res)
	return res
}
