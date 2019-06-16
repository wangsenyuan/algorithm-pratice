package main

import "fmt"

func main() {
	//input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	//input := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	input := "dir\n        file.txt"
	fmt.Println(lengthLongestPath(input))
}

func lengthLongestPath(input string) int {
	stack := make([]int, len(input))
	p := 0
	from := 0
	isFile := false
	res := 0
	level := 0
	i := 0
	for i < len(input) {
		c := input[i]
		if c == '\n' {
			for p > level {
				p--
			}
			l := i - from
			if p > 0 {
				l += 1 + stack[p-1]
			}
			if isFile {
				if l > res {
					res = l
				}
			} else {
				stack[p] = l
				p++
			}
			level = 0
			from = i + 1
			isFile = false
		} else if c == '\t' {
			level++
			from = i + 1
		} else if c == '.' {
			isFile = true
		}
		i++
	}

	if i > from && isFile {
		for p > level {
			p--
		}
		l := i - from
		if p > 0 {
			l += 1 + stack[p-1]
		}
		if l > res {
			res = l
		}
	}

	return res
}
