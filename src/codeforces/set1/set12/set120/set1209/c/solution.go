package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		readString(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(s string) string {
	n := len(s)
	ans := make([]byte, n)
	next := make([]byte, n+1)
	next[n] = '9'

	for i := n - 1; i >= 0; i-- {
		next[i] = min(next[i+1], s[i])
	}

	buf := make([]byte, n)

	check := func(x byte) string {
		for i := 0; i < n; i++ {
			if s[i] > x || s[i] == x && next[i] < s[i] {
				ans[i] = '2'
			} else {
				ans[i] = '1'
			}
		}
		var j int
		for i := 0; i < n; i++ {
			if ans[i] == '1' {
				buf[j] = s[i]
				j++
			}
		}
		for i := 0; i < n; i++ {
			if ans[i] == '2' {
				buf[j] = s[i]
				j++
			}
		}
		for i := 1; i < n; i++ {
			if buf[i-1] > buf[i] {
				return "-"
			}
		}
		return string(buf)
	}

	for i := byte('0'); i <= '9'; i++ {
		tmp := check(i)
		if tmp != "-" {
			return string(ans)
		}
	}
	return "-"
}
