package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
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

func solve(s string) int {
	// n is for ng and g is for max(n, g) - ng
	cnt := make([]int, 5)

	const vowels = "AEIOU"

	for _, c := range []byte(s) {
		if c == 'N' {
			cnt[2]++
		} else if c == 'G' {
			cnt[3]++
		} else if c == 'Y' {
			cnt[4]++
		} else if strings.IndexByte(vowels, c) >= 0 {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}
	n := len(s)

	a := cnt[0]
	y := cnt[4]
	b := n - a - y
	ng := min(cnt[2], cnt[3])

	if (a+y)*2 <= ng {
		return (a + y) * 5
	}
	if (a+y)*2 <= b-ng {
		return (a+y)*3 + ng
	}
	if (b+y)/2 <= a {
		res := (b + y) / 2 * 3
		if (b+y)%2 == 1 && ng > 0 {
			res++
		}
		return res
	}
	// 如果没有 NG，那么答案一定是 3 的倍数
	// 如果只有一个 NG，那么当 n 是 3k+2 时，一定会多出一个字母，例如 AYNGG
	// 其余情况可以用 NG 和 Y 灵活调整，答案是 n
	return n - max(n%3-ng, 0)
}
