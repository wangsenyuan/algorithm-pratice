package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("aba", "cdc"))
	fmt.Println(findLUSlength("aba", "cdca"))
	fmt.Println(findLUSlength("aba", "acbca"))

}

func findLUSlength(a string, b string) int {

	isSubSeq := func(x, y string) bool {
		i, j := 0, 0
		for i < len(x) && j < len(y) {
			if x[i] == y[j] {
				i++
			}
			j++
		}
		return i == len(x)
	}

	var res = -1
	if !isSubSeq(a, b) {
		res = len(a)
	}

	if !isSubSeq(b, a) && len(b) > res {
		res = len(b)
	}

	return res
}
