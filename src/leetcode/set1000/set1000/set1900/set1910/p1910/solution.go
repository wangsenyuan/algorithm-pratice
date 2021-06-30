package p1910

import "bytes"

func removeOccurrences(s string, part string) string {
	n := len(s)
	lps := findLps(part)
	stack := make([]int, n)
	L := make([]int, n)
	V := make([]bool, n)
	L[0] = -1
	var p int
	for i, j := 0, 0; i < n; {
		if s[i] == part[j] {
			L[i] = j
			stack[p] = i
			p++

			i++
			j++
		}
		if j == len(part) {
			// find a match
			for k := p - len(part); k < p; k++ {
				V[stack[k]] = true
			}
			p -= len(part)
			if p > 0 {
				j = L[stack[p-1]] + 1
			} else {
				j = 0
			}
			continue
		}
		if i < n && s[i] != part[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if !V[i] {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}

func findLps(s string) []int {
	n := len(s)
	lps := make([]int, n)

	var l int

	for i := 1; i < n; {
		if s[i] == s[l] {
			l++
			lps[i] = l
			i++
			continue
		}
		if l != 0 {
			l = lps[l-1]
		} else {
			i++
		}
	}
	return lps
}
