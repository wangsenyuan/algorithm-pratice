package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for n > 0 {
		var s string
		fmt.Scanf("%s", &s)

		if isPalindrome(s) {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}

		n--
	}
}
func isPalindrome(str string) bool {
	for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
