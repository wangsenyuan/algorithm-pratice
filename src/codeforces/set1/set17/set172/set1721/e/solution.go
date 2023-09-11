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
	n := readNum(reader)
	Q := make([]string, n)
	for i := 0; i < n; i++ {
		Q[i] = readString(reader)
	}
	res := solve(s, Q)
	for _, cur := range res {
		var buf bytes.Buffer
		for i := 0; i < len(cur); i++ {
			buf.WriteString(fmt.Sprintf("%d ", cur[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(s string, Q []string) [][]int {
	//pi = max k, s[1...k] = s[i - k + 1, i]
	// 是不是就是 kmp中的lps?
	p := kmp(s)
	n := len(s)

	A := make([][]int, n+11)
	for i := 0; i < n; i++ {
		A[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			if i > 0 && j != int(s[i]-'a') {
				A[i][j] = A[p[i-1]][j]
			} else {
				A[i][j] = i
				// 有可能 i = 0
				if j == int(s[i]-'a') {
					A[i][j]++
				}
			}
		}
	}
	for i := n; i < len(A); i++ {
		A[i] = make([]int, 26)
		p = append(p, 0)
	}

	ans := make([][]int, len(Q))

	for it, t := range Q {
		m := len(t)
		ans[it] = make([]int, m)
		for i := n; i < n+m; i++ {
			for j := 0; j < 26; j++ {
				if j != int(t[i-n]-'a') {
					A[i][j] = A[p[i-1]][j]
				} else {
					A[i][j] = i + 1
				}
			}
			tmp := A[p[i-1]][int(t[i-n]-'a')]
			p[i] = tmp
			ans[it][i-n] = tmp
		}
	}

	return ans
}

func kmp(s string) []int {
	n := len(s)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		res[i] = j
	}
	return res
}
