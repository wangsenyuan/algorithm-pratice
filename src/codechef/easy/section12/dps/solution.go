package main

import (
	"fmt"
)

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var s string
		fmt.Scanf("%s", &s)
		res := solve(s)
		if res {
			fmt.Println("DPS")
		} else {
			fmt.Println("!DPS")
		}
	}
}

func solve(s string) bool {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	var x int

	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			x++
		}
	}

	n := len(s)

	if n%2 == 0 {
		return x == 2
	}

	return x == 1 || x == 3
}
