package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s := readString(reader)
	return solve(s)
}

func solve(s string) int {
	// 遇到一个 :增加一层, 遇到一个 .减少一层

	freq := make(map[string]int)

	var i int

	var res int
	var dfs func()

	dfs = func() {
		if i >= len(s) {
			return
		}
		// i是开始
		j := i
		for i < len(s) && s[i] != '.' && s[i] != ':' {
			i++
		}
		name := s[j:i]
		res += freq[name]
		freq[name]++
		i++
		if s[i-1] == ':' {
			// 它有子节点
			for {
				dfs()
				i++
				if i >= len(s) || s[i-1] != ',' {
					break
				}
			}
		}
		freq[name]--
	}

	dfs()

	return res
}
