package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
	if k == 0 {
		return s
	}
	res := []byte(s)
	for i := 0; i < len(s); {
		j := i + 2*k
		if j > len(s) {
			j = len(s)
		}

		a := i + k - 1

		if a > j-1 {
			a = j - 1
		}

		for i < a {
			res[i], res[a] = res[a], res[i]
			i++
			a--
		}

		i = j
	}

	return string(res)
}
