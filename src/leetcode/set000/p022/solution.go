package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, 10)
	process(n, 0, make([]rune, 0, 2*n), &res)
	return res
}

func process(open int, close int, path []rune, res *[]string) {
	if len(path) == cap(path) {
		*res = append(*res, string(path))
		return
	}

	if open > 0 {
		process(open-1, close+1, append(path, '('), res)
	}

	if close > 0 {
		process(open, close-1, append(path, ')'), res)
	}
}

func main() {
	res := generateParenthesis(4)
	for _, r := range res {
		fmt.Printf("%v\n", r)
	}
}
