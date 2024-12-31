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

func getDepth(s string) int {
	var res int
	var level int
	for i := range len(s) {
		if s[i] == '(' {
			level++
		} else {
			level--
		}
		res = max(res, level)
	}
	return res
}

func solve(s string) string {
	x := getDepth(s)
	h := (x + 1) / 2

	n := len(s)
	stack := make([]int, n)
	var top int

	res := make([]byte, n)
	for i := range n {
		res[i] = '0'
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			j := stack[top-1]
			if top > h {
				res[j] = '1'
				res[i] = '1'
			}

			top--
		}
	}

	return string(res)
}
