package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	res := solve(s)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(s string) bool {
	// @ and /
	var i int
	for i < len(s) && s[i] != '@' {
		i++
	}
	if i == len(s) || !check(s[:i]) {
		return false
	}
	// s[i] = '@'
	i++
	j := i
	for j < len(s) && s[j] != '/' {
		if s[j] == '.' {
			if !check(s[i:j]) {
				return false
			}
			i = j + 1
		}
		j++
	}
	if i >= j || !check(s[i:j]) {
		return false
	}
	j++
	if j == len(s) || j < len(s) && !check(s[j:]) {
		return false
	}
	return true
}

func check(name string) bool {
	//sequence of Latin letters (lowercase or uppercase), digits or underscores characters _
	if len(name) > 16 || len(name) == 0 {
		return false
	}
	for _, x := range name {
		if unicode.IsUpper(x) || unicode.IsLower(x) || unicode.IsDigit(x) || x == '_' {
			continue
		}
		return false
	}

	return true
}
