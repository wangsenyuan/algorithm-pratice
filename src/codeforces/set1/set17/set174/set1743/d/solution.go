package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)[:n]
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(s string) string {
	pos1 := -1
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			pos1 = i
			break
		}
	}
	if pos1 < 0 {
		return "0"
	}
	pos2 := pos1 + 1
	for pos2 < n && s[pos2] == '1' {
		pos2++
	}
	if pos2 == n {
		return s[pos1:]
	}

	ans := s[pos1:]
	cur := pos1
	var not_needed int
	for {
		if cur == n || s[cur] == '1' && cur > pos2 {
			break
		}
		tmp := or(s[pos1:], s[pos1:n-not_needed])
		if ans < tmp {
			ans = tmp
		}
		cur++
		not_needed++
	}

	return ans
}

func or(a, b string) string {
	sz := max(len(a), len(b))
	buf := make([]byte, sz)

	for i := sz - 1; i >= 0; i-- {
		buf[i] = '0'
		ln := sz - i
		if ln <= len(a) && a[len(a)-ln] == '1' {
			buf[i] = '1'
		}
		if ln <= len(b) && b[len(b)-ln] == '1' {
			buf[i] = '1'
		}
	}

	return string(buf)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
