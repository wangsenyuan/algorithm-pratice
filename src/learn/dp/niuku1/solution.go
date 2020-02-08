// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	buf := scanner.Bytes()

	fmt.Println(solve(buf))

}

func solve(buf []byte) string {
	if len(buf) == 0 {
		return "."
	}

	var i, j int

	n := len(buf)
	begin := true
	for i < n {
		if isLetter(buf[i]) {
			buf[j] = buf[i]

			if begin {
				if buf[j] >= 'a' && buf[j] <= 'z' {
					buf[j] = byte(buf[j] - 'a' + 'A')
				}
				begin = false
			}

			j++
		} else {
			if !begin {
				buf[j] = ' '
				j++
			}
			begin = true
		}

		i++
	}

	if buf[j-1] == ' ' {
		buf[j-1] = '.'
		return string(buf[:j])
	}

	return string(buf) + "."
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}
