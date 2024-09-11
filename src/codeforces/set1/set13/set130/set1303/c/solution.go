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
		res := solve(s)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%s\n", res))
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
func solve(s string) string {
	vis := make([]bool, 26)
	vis[int(s[0]-'a')] = true
	res := s[:1]
	var pos int

	for i := 1; i < len(s); i++ {
		if vis[int(s[i]-'a')] {
			if pos > 0 && res[pos-1] == s[i] {
				pos--
			} else if pos+1 < len(res) && res[pos+1] == s[i] {
				pos++
			} else {
				return ""
			}
		} else {
			if pos == 0 {
				res = fmt.Sprintf("%c%s", s[i], res)
			} else if pos == len(res)-1 {
				res = fmt.Sprintf("%s%c", res, s[i])
				pos++
			} else {
				return ""
			}
		}
		vis[int(s[i]-'a')] = true
	}

	var buf bytes.Buffer

	buf.WriteString(res)

	for i := 0; i < 26; i++ {
		if !vis[i] {
			buf.WriteByte(byte(i + 'a'))
		}
	}

	return buf.String()
}

func abs(num int) int {
	return max(num, -num)
}
