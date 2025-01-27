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

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	var ans int

	for i := 0; i < n; i += 2 {
		if i == n-1 {
			if s[i] != '1' {
				ans++
			}
		} else {
			if s[i] == s[i+1] {
				ans++
			} else if s[i] == '0' {
				ans += 2
			}
		}
	}

	if ans*2 > n {
		ans = n - ans
	}
	return ans
}
