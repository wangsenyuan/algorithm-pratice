package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		s := readString(reader)
		a := readString(reader)
		res := solve(s, a)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(s string, t string) bool {
	if check1(s, t) {
		return true
	}

	n := len(s)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		next[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			next[i][j] = n + 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		copy(next[i], next[i+1])
		x := int(s[i] - 'a')
		next[i][x] = i
	}

	check := func(t1, t2 string) bool {
		dp := make([][]int, len(t1)+1)
		for i := 0; i <= len(t1); i++ {
			dp[i] = make([]int, len(t2)+1)
			for j := 0; j <= len(t2); j++ {
				dp[i][j] = len(s) + 10
			}
		}
		dp[0][0] = -1

		for i := 0; i <= len(t1); i++ {
			for j := 0; j <= len(t2); j++ {
				if i == 0 && j == 0 {
					continue
				}
				if i > 0 && dp[i-1][j]+1 < n {
					x := int(t1[i-1] - 'a')
					y := next[dp[i-1][j]+1][x]
					dp[i][j] = min(dp[i][j], y)
				}
				if j > 0 && dp[i][j-1]+1 < n {
					x := int(t2[j-1] - 'a')
					y := next[dp[i][j-1]+1][x]
					dp[i][j] = min(dp[i][j], y)
				}
			}
		}

		return dp[len(t1)][len(t2)] < len(s)
	}

	for i := 1; i < len(t); i++ {
		if check(t[:i], t[i:]) {
			return true
		}
	}
	return false
}

func check1(s, t string) bool {
	for i, j := 0, 0; i < len(t); i++ {
		for j < len(s) && s[j] != t[i] {
			j++
		}
		if j == len(s) {
			return false
		}
		j++
	}
	return true
}
