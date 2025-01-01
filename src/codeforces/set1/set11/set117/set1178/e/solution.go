package main

import (
	"bufio"
	"fmt"
	"os"
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

type pair struct {
	first  int
	second int
}

func solve(s string) string {

	n := len(s)
	marked := make([]bool, n)
	for i, j := 0, n-1; i <= j; {
		if s[i] == s[j] {
			marked[i] = true
			marked[j] = true
			i++
			j--
			continue
		}
		if i+1 <= j && s[i+1] == s[j] {
			i++
			continue
		}
		if i <= j-1 && s[j-1] == s[i] {
			j--
			continue
		}
		// ac     cb
		i++
		j--
	}
	var res []byte
	for i := 0; i < n; i++ {
		if marked[i] {
			res = append(res, s[i])
		}
	}
	return string(res)
}
