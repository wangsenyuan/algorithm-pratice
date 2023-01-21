package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	cnt := make([]int, 3)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			cnt[0]++
		} else if s[i] == ')' {
			cnt[1]++
		} else {
			cnt[2]++
		}
	}

	// open = close = n / 2
	if cnt[0] > n/2 || cnt[1] > n/2 {
		return false
	}

	if cnt[2] == 0 {
		return cnt[0] == cnt[1]
	}

	// cnt[0] + x = n / 2
	x := n/2 - cnt[0]

	buf := []byte(s)
	a, b := -1, -1
	for i := 0; i < n; i++ {
		if buf[i] == '?' {
			if x > 0 {
				buf[i] = '('
				x--
				a = i
			} else {
				if b < 0 {
					b = i
				}
				buf[i] = ')'
			}
		}
	}
	var open int
	for i := 0; i < n; i++ {
		if buf[i] == '(' {
			open++
		} else {
			open--
		}
		if open < 0 {
			return false
		}
	}

	if open != 0 {
		return false
	}
	// a, b
	if a < 0 || b < 0 {
		// can't change
		return true
	}

	// a < b
	buf[a], buf[b] = buf[b], buf[a]

	for i := 0; i < n; i++ {
		if buf[i] == '(' {
			open++
		} else {
			open--
		}
		if open < 0 {
			return true
		}
	}

	return open != 0
}
