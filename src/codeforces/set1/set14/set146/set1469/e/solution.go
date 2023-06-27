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
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			buf.WriteString(res)
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const inf = 1 << 30

func solve(s string, k int) string {
	// 如果在0...n-k中间有0/1
	// 如果t[0] = 0， 则那么以0开头的子串和t已经match了，这些可以去掉了
	n := len(s)
	m := min(k, ceilLog(n-k+2))
	used := make([]bool, 1<<m)
	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = inf
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			next[i] = i
		} else if i+1 < n {
			next[i] = next[i+1]
		}
	}
	mask := (1 << m) - 1
	d := k - m
	for j := 0; j < n-k+1; j++ {
		if next[j]-j < d {
			// 在前d个字符中包含0
			continue
		}
		var cur int
		for i := j + d; i < j+k; i++ {
			cur = cur*2 + int(s[i]-'0')
		}
		used[mask^cur] = true
	}

	for i := 0; i <= mask; i++ {
		if !used[i] {
			buf := make([]byte, k)
			for j := k - 1; j >= 0; j-- {
				buf[j] = byte('0' + (i % 2))
				i /= 2
			}
			return string(buf)
		}
	}

	return ""
}

func ceilLog(n int) int {
	var h int
	for (1 << h) < n {
		h++
	}
	// 1 << h >= n
	return h
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
