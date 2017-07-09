package main

import "fmt"

const mod = 1000000007

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	first := 1
	second := 1
	if s[0] == '*' {
		second = 9
	}

	for i := 1; i < len(s); i++ {
		tmp := second
		if s[i] == '*' {
			second *= 9
			if s[i-1] == '1' {
				second = (second + 9*first) % mod
			} else if s[i-1] == '2' {
				second = (second + 6*first) % mod
			} else if s[i-1] == '*' {
				second = (second + 15*first) % mod
			}
		} else {
			if s[i] == '0' {
				second = 0
			}
			if s[i-1] == '1' {
				second = (second + first) % mod
			} else if s[i-1] == '2' && s[i] < '7' {
				second = (second + first) % mod
			} else if s[i-1] == '*' && s[i] < '7' {
				second = (second + 2*first) % mod
			} else if s[i-1] == '*' {
				second = (second + first) % mod
			}
		}
		first = tmp
	}

	return second
}

func main() {
	fmt.Println(numDecodings("1*"))
	fmt.Println(numDecodings("*"))
	fmt.Println(numDecodings("2*"))
	fmt.Println(numDecodings("**"))
	fmt.Println(numDecodings("0"))
}
