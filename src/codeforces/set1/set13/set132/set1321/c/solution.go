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

func solve(s string) int {
	n := len(s)
	prev := make([]int, n)
	next := make([]int, n)
	pos := make([][]int, 26)
	for i := 0; i < n; i++ {
		prev[i] = i - 1
		next[i] = i + 1
		pos[int(s[i]-'a')] = append(pos[int(s[i]-'a')], i)
	}

	var res int
	for x := 25; x >= 0; x-- {
		for i := 0; i < len(pos[x]); i++ {
			j := i
			i++
			for i < len(pos[x]) && prev[pos[x][i]] == pos[x][i-1] {
				i++
			}
			i--
			u, v := prev[pos[x][j]], next[pos[x][i]]
			if u >= 0 && int(s[u]-'a') == x-1 || v < n && int(s[v]-'a') == x-1 {
				res += i - j + 1
				if u >= 0 {
					next[u] = v
				}
				if v < n {
					prev[v] = u
				}
			}
		}
	}

	return res
}
