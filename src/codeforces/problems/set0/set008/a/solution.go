package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')

	ans := solve(normalize(s), normalize(a), normalize(b))

	fmt.Println(ans)
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
}

func solve(s string, a string, b string) string {
	res1 := check(s, a, b)
	res2 := check(reverse(s), a, b)

	if res1 {
		if res2 {
			return "both"
		}
		return "forward"
	}

	if res2 {
		return "backward"
	}
	return "fantasy"
}

func check(s string, a string, b string) bool {
	x := getLps(a)
	y := getLps(b)

	i := find(s, a, x)
	if i >= 0 {
		j := i + len(a)

		if find(s[j:], b, y) >= 0 {
			return true
		}
	}
	return false
}

func find(s string, pat string, lps []int) int {
	for i, j := 0, 0; i < len(s); {
		if s[i] == pat[j] {
			i++
			j++
		}
		if j == len(pat) {
			return i - len(pat)
		}
		if i < len(s) && s[i] != pat[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

func getLps(pat string) []int {
	m := len(pat)
	lps := make([]int, m)
	for l, i := 0, 1; i < m; {
		if pat[l] == pat[i] {
			l++
			lps[i] = l
			i++
		} else {
			if l > 0 {
				l = lps[l-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func reverse(s string) string {
	buf := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}
