package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Scanf("%s", &str)
		res := solve(str)
		fmt.Println(res)
	}
}

func solve(s string) int64 {
	n := int64(len(s))

	total := n * (n + 1) / 2

	j := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '4' {
			m := int64(i - j)
			tmp := m * (m + 1) / 2
			total -= tmp
			j = i + 1
		}
	}
	return total
}
