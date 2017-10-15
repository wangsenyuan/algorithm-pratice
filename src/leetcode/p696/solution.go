package main

import "fmt"

func countBinarySubstrings(s string) int {
	var res int

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' && s[i+1] == '1' {
			for k := 0; i-k >= 0 && s[i-k] == '0' && i+1+k < len(s) && s[i+1+k] == '1'; k++ {
				res += 1
			}
		} else if s[i] == '1' && s[i+1] == '0' {
			for k := 0; i-k >= 0 && s[i-k] == '1' && i+1+k < len(s) && s[i+1+k] == '0'; k++ {
				res += 1
			}
		}
	}

	return res
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}
