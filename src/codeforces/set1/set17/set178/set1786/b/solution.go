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
		n, w, h := readThreeNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(w, h, a, b)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

const inf = 1 << 30

func solve(w int, h int, a []int, b []int) bool {
	var l, r int = -inf, inf
	ok := true
	for i := 0; i < len(a); i++ {
		// 如果考虑就是往左移动距离x
		// a[i] - x - w <= b[i] - h
		// a[i] - y + w >= b[i] + h
		x := a[i] - w - (b[i] - h)
		// x >=
		y := a[i] + w - (b[i] + h)
		// y <=
		// y >= x
		l = max(l, x)
		r = min(r, y)
		if l > r {
			ok = false
			break
		}
	}
	if ok {
		return true
	}
	l, r = -inf, inf
	for i := 0; i < len(a); i++ {
		// 如果考虑就是往右移动距离x
		// a[i] + x - w <= b[i] - h
		// a[i] + y + w >= b[i] + h
		x := b[i] - h - (a[i] - w)
		// x >=
		y := b[i] + h - (a[i] + w)

		l = max(l, y)
		r = min(r, x)
		if l > r {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
