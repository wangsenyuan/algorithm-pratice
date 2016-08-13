package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		var s string
		fmt.Scanf("%s", &s)
		if check(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}

func check(s string) bool {
	if len(s) <= 1 {
		return true
	}
	c := make([]int, 26)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		c[s[i]-'a']++
		c[s[j]-'a']--
	}

	for _, x := range c {
		if x != 0 {
			return false
		}
	}

	return true
}
