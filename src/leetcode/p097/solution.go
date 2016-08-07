package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) (result bool) {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	checked := make([][]bool, len(s1)+1)

	for i := range checked {
		checked[i] = make([]bool, len(s2)+1)
	}

	var check func(a, b, c int) bool

	check = func(a, b, c int) bool {
		//fmt.Printf("check %d %d %d\n", a, b, c)
		if c == len(s3) {
			panic("solution found")
		}

		if checked[a][b] {
			return false
		}

		if a < len(s1) && s1[a] == s3[c] {
			check(a+1, b, c+1)
		}

		if b < len(s2) && s2[b] == s3[c] {
			check(a, b+1, c+1)
		}
		checked[a][b] = true
		return false
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("aa")
			result = true
		}
	}()

	check(0, 0, 0)

	return
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[len(s1)][len(s2)] = true

	for a := len(s1); a >= 0; a-- {
		for b := len(s2); b >= 0; b-- {
			c := a + b
			if c == len(s3) {
				continue
			}

			if a < len(s1) && s1[a] == s3[c] {
				dp[a][b] = dp[a][b] || dp[a+1][b]
			}

			if b < len(s2) && s2[b] == s3[c] {
				dp[a][b] = dp[a][b] || dp[a][b+1]
			}
		}
	}

	return dp[0][0]
}

func main() {
	a := "aabcc"
	b := "dbbca"
	c := "aadbbcbcac"

	fmt.Println(isInterleave2(a, b, c))
}
