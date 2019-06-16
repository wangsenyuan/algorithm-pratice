package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}

	t := make([]byte, len(s), len(s))
	j := 0
	r := 0
	p := [...]int{2 * (numRows - 1), 0}
	x := 0
	for i := 0; i < len(s); i++ {
		t[i] = s[j]
		if j+p[x] >= len(s) {
			r++
			j = r
			p[0] = p[0] - 2
			p[1] = p[1] + 2
			x = 0
			if p[0] == 0 {
				x = 1
			}
			continue
		}

		j += p[x]
		if p[1-x] > 0 {
			x = 1 - x
		}
	}

	return string(t)
}

func main() {
	fmt.Println(convert("PAHNAPLSIIGYIR", 3))
}
