package main

import "fmt"

func main() {
	var tc int

	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var s string

		fmt.Scanf("%s", &s)
		if solve(s) {
			fmt.Println("WIN")
		} else {
			fmt.Println("LOSE")
		}
	}
}

func solve(s string) bool {
	var cnt int

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		}
	}

	return cnt%2 == 1
}
