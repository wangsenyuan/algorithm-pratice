package main

import "fmt"

func main() {
	//input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	input := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	//input := "dir\n        file.txt";
	fmt.Println(lengthLongestPath(input))
}

func lengthLongestPath(input string) int {
	stack := make([]folder, len(input))
	p := 0
	from := 0
	isFile := false
	res := 0
	level := 0
	i := 0
	for i < len(input) {
		c := input[i]
		if c == '\n' {
			for p > 0 && stack[p-1].level >= level {
				p--
			}
			if isFile {
				l := fileLength(stack, p, i-from)
				if l > res {
					res = l
				}
			} else {
				path := input[from:i]
				if p > 0 {
					last := stack[p-1]
					path = last.path + "/" + path
				}
				stack[p] = folder{path, level}
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
		for p > 0 && stack[p-1].level >= level {
			p--
		}
		l := fileLength(stack, p, i-from)
		if l > res {
			res = l
		}
	}

	return res
}

func fileLength(stack []folder, p int, l int) int {
	if p == 0 {
		return l
	}

	last := stack[p-1]
	return len(last.path) + l + 1
}

type folder struct {
	path  string
	level int
}
