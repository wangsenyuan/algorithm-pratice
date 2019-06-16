package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
}

func simplifyPath(path string) string {
	var stack []string

	i := 0
	j := 1
	for ; j < len(path); j++ {
		if path[j] != '/' {
			continue
		}

		if i == j-1 {
			i = j
			continue
		}

		stack = update(stack, path[i+1:j])
		i = j
	}

	if i < j-1 {
		stack = update(stack, path[i+1:j])
	}

	return "/" + strings.Join(stack, "/")
}

func update(stack []string, s string) []string {
	if s == ".." {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	} else if s != "." {
		stack = append(stack, s)
	}
	return stack
}
