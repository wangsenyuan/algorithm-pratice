package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}

func shortestPalindrome(s string) string {
	n := len(s)
	N := n<<1 | 1
	S := make([]byte, N)

	for i := 0; i < n; i++ {
		S[i<<1] = '|'
		S[i<<1|1] = s[i]
	}
	S[N-1] = '|'
	L := make([]int, N)
	C, R, P := 0, -1, 0
	for i := 0; i < N; i++ {
		var rad int

		if R >= i {
			rad = min(L[2*C-i], R-i)
		}
		for i-rad >= 0 && i+rad < N && S[i-rad] == S[i+rad] {
			rad++
		}
		L[i] = rad

		if rad+i > R {
			C = i
			R = rad + i - 1
		}
		if rad == i+1 {
			P = R
		}
	}

	P /= 2

	res := make([]byte, 2*n-P)
	// var buf bytes.Buffer
	var j int
	for i := len(s) - 1; i >= P; i-- {
		res[j] = s[i]
		j++
	}

	copy(res[j:], s)

	return string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
