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

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadBytes('\n')
	for i := range len(s) {
		if s[i] == '\r' || s[i] == '\n' {
			return string(s[:i])
		}
	}
	return string(s)
}

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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	res := make([]int, n)
	for i := range n {
		a := readString(reader)
		b := readString(reader)
		c := readString(reader)
		res[i] = solve(a, b, c)
	}
	return res
}

const inf = 1 << 60

func solve(a string, b string, c string) int {
	n := len(a)
	m := len(b)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i+j == 0 {
				continue
			}
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+check(c[i+j-1] != a[i-1]))
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+check(c[i+j-1] != b[j-1]))
			}
		}
	}
	return dp[n][m]
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
