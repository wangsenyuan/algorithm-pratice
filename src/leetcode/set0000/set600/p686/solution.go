package main

import (
	"strings"
	"fmt"
)

func repeatedStringMatch(A string, B string) int {
	if strings.Contains(A, B) {
		return 1
	}

	if strings.Contains(B, A) {
		a := strings.Index(B, A)

		ans := 0

		if a > 0 {
			if !strings.HasSuffix(A, B[:a]) {
				// the prefix of B can't be gain from A suffix
				return -1
			}
			ans = 1
			B = B[a:]
		}

		for len(B) > 0 && strings.HasPrefix(B, A) {
			B = B[len(A):]
			ans++
		}

		if len(B) > 0 {
			if !strings.HasPrefix(A, B) {
				return -1
			}
			ans++
		}

		return ans
	}

	if strings.Contains(A+A, B) {
		return 2
	}

	return -1
}

func main() {
	//fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abababaaba", "aabaaba"))
}
