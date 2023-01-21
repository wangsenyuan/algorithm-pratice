package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var res int
	var cnt int
	for {
		s, err := reader.ReadString('\n')
		if len(s) == 1 || err == io.EOF {
			break
		}
		if s[0] == '+' {
			cnt++
		} else if s[0] == '-' {
			cnt--
		} else {
			var i int
			for i < len(s) && s[i] != ':' {
				i++
			}
			i++
			// s[i] == ':'
			n := len(s)
			for n > i && (s[n-1] == '\n' || s[n-1] == ' ') {
				n--
			}
			l := n - i
			res += l * cnt
		}
	}

	fmt.Println(res)
}
