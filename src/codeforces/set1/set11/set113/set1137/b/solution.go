package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readString(reader)
	b := readString(reader)

	res := solve(a, b)

	fmt.Println(res)
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

func solve(s string, x string) string {
	pi := kmp(x)

	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}

	var buf bytes.Buffer
	var j int
	for cnt[int(x[j]-'0')&1] > 0 {
		buf.WriteByte(x[j])
		cnt[int(x[j]-'0')&1]--
		j++
		if j == len(x) {
			j = pi[j-1]
		}
	}

	for cnt[0] > 0 {
		buf.WriteByte('0')
		cnt[0]--
	}
	for cnt[1] > 0 {
		buf.WriteByte('1')
		cnt[1]--
	}
	return buf.String()
}

func kmp(s string) []int {
	n := len(s)
	next := make([]int, n)
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func solve1(s string, x string) string {
	if len(s) < len(x) {
		return s
	}
	cnt := make([]int, 2)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
	}
	n := len(x)
	pref := make([][]int, n+1)

	for i := 0; i < n; i++ {
		pref[i+1] = make([]int, 2)
		copy(pref[i+1], pref[i])
		pref[i+1][int(x[i]-'0')]++
	}

	z := zFunction(x)

	var buf bytes.Buffer

	for cnt[0] >= pref[n][0] && cnt[1] >= pref[n][1] {
		buf.WriteString(x)
		cnt[0] -= pref[n][0]
		cnt[1] -= pref[n][1]

		for i := 1; i < n; i++ {
			if z[i] == n-i {
				// 从z[i]开始的可以到达结尾
				expect := []int{pref[n][0] - pref[n-i][0], pref[n][1] - pref[n-i][1]}
				if cnt[0] < expect[0] || cnt[1] < expect[1] {
					continue
				}
				buf.WriteString(x[n-i:])
				cnt[0] -= expect[0]
				cnt[1] -= expect[1]
				// keep at the same position
				i--
			}
		}
	}

	for cnt[0] > 0 {
		buf.WriteByte('0')
		cnt[0]--
	}
	for cnt[1] > 0 {
		buf.WriteByte('1')
		cnt[1]--
	}

	return buf.String()
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
