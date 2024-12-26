package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	r := readString(reader)
	res := solve(r)
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

func solve(s string) (res string) {
	n := len(s)
	check := func(exp string) bool {
		i := strings.Index(exp, "=")
		if i < 0 {
			return false
		}
		ok1, left := calc(exp[:i])
		ok2, right := calc(exp[i+1:])
		return ok1 && ok2 && left == right
	}

	if check(s) {
		return "Correct"
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		// invalid
		buf[i] = '*'
	}
	for i := 0; i < n; i++ {
		if isDigit(s[i]) {
			// move it to some place
			for j := 0; j < n; j++ {
				if j == i {
					continue
				}
				if j < i {
					copy(buf[:j], s[:j])
					buf[j] = s[i]
					copy(buf[j+1:i+1], s[j:i])
					copy(buf[i+1:], s[i+1:])
				} else {
					// i < j
					copy(buf[:i], s[:i])
					copy(buf[i:j], s[i+1:j+1])
					buf[j] = s[i]
					copy(buf[j+1:], s[j+1:])
				}
				if check(string(buf)) {
					return string(buf)
				}
			}
		}
	}

	return "Impossible"
}

func calc(s string) (ok bool, res int) {
	if len(s) == 0 {
		return false, 0
	}
	sign := 1
	for i := 0; i < len(s); {
		if !isDigit(s[i]) {
			return false, 0
		}
		var num int
		j := i
		for i < len(s) && isDigit(s[i]) {
			if num == 0 && i > j {
				return false, 0
			}
			num = num*10 + int(s[i]-'0')
			i++
		}
		if i-j > 10 {
			// at most 10 digits
			return false, 0
		}
		res += sign * num
		if i == len(s) {
			sign = 0
			break
		}
		if s[i] == '+' {
			sign = 1
		} else {
			sign = -1
		}
		i++
	}

	return sign == 0, res
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
