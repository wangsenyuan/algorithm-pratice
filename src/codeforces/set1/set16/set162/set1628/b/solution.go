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
		n := readNum(reader)
		s := make([]string, n)
		for i := 0; i < n; i++ {
			s[i] = readString(reader)
		}
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

func solve(s []string) bool {
	// s[i] + s[j] + s[k] is palindrome
	for _, w := range s {
		if len(w) == 1 {
			return true
		}
		if reverse(w) == w {
			return true
		}
	}
	// len(w) >= 2 and reverse(w) != w
	// 最多有两个连起来
	// abc ba
	if solve1(s) {
		return true
	}
	if solve2(s) {
		return true
	}
	return solve3(s)
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func solve1(s []string) bool {
	X := 26 * 26
	dp := make([]bool, X)

	for _, w := range s {
		if len(w) == 2 {
			x := int(w[1]-'a')*26 + int(w[0]-'a')
			if dp[x] {
				return true
			}
			x = int(w[0]-'a')*26 + int(w[1]-'a')
			dp[x] = true
		} else {
			x := int(w[2]-'a')*26 + int(w[1]-'a')
			if dp[x] {
				return true
			}
			// just continue
		}
	}
	return false
}

func solve2(s []string) bool {
	X := 26 * 26
	dp := make([]bool, X)

	for i := len(s) - 1; i >= 0; i-- {
		w := s[i]
		if len(w) == 2 {
			x := int(w[0]-'a')*26 + int(w[1]-'a')
			if dp[x] {
				return true
			}
			x = int(w[1]-'a')*26 + int(w[0]-'a')
			dp[x] = true
		} else {
			x := int(w[0]-'a')*26 + int(w[1]-'a')
			if dp[x] {
				return true
			}
		}
	}
	return false
}

func solve3(s []string) bool {
	mem := make(map[string]bool)

	for _, w := range s {
		r := reverse(w)
		if mem[r] {
			return true
		}
		mem[w] = true
	}

	return false
}
