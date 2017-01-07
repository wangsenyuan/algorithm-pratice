package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var cnt []int
	var s string

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		if i == 0 {
			cnt = countLetters(s)
		} else {
			cnt = merge(cnt, countLetters(s))
		}
	}

	var res string
	for i := 0; i < 26; i++ {
		x := cnt[i]
		for j := 0; j < x; j++ {
			res += string('a' + i)
		}
	}
	if len(res) > 0 {
		fmt.Println(res)
	} else {
		fmt.Println("no such string")
	}
}

func merge(a, b []int) []int {
	c := make([]int, 26)
	for i := 0; i < 26; i++ {
		c[i] = min(a[i], b[i])
	}
	return c
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countLetters(s string) []int {
	cnt := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cnt[s[i] - 'a']++
	}

	return cnt
}
