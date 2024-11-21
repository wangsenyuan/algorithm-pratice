package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
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

const bad = ":("

func solve(s string) string {
	n := len(s)

	if n%2 == 1 {
		return bad
	}
	buf := []byte(s)

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

	if cnt[0] > n/2 || cnt[1] > n/2 {
		// 最后都没法平衡
		return bad
	}
	// cnt[0] + x = cnt[1] + y => x - y = cnt[1] - cnt[0]
	// x + y = cnt[2]
	if (cnt[1]-cnt[0]+cnt[2])%2 == 1 {
		return bad
	}
	// x变成（， y个变成 )
	x := (cnt[1] - cnt[0] + cnt[2]) / 2
	// y := cnt[2] - x
	var level int
	for i := 0; i < n; i++ {
		if buf[i] == '?' {
			if x > 0 {
				buf[i] = '('
				x--
			} else {
				buf[i] = ')'
			}
		}
		if buf[i] == '(' {
			level++
		} else {
			level--
		}
		if level < 0 || level == 0 && i+1 < n {
			return bad
		}
	}

	return string(buf)
}
