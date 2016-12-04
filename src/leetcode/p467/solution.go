package main

import "fmt"

func main() {
	fmt.Println(findSubstringInWraproundString("zabd"))
}

func findSubstringInWraproundString(p string) int {
	letters := make([]int, 26)
	curLen := 0
	res := 0
	for i := 0; i < len(p); i++ {
		if i > 0 && p[i] != (p[i-1]-'a'+1)%26+'a' {
			curLen = 0
		}

		curLen++
		if curLen > letters[p[i]-'a'] {
			res += curLen - letters[p[i]-'a']
			letters[p[i]-'a'] = curLen
		}
	}

	return res
}
