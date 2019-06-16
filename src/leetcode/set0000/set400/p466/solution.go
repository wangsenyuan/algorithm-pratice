package main

import "fmt"

func main() {
	s1 := "acb"
	n1 := 4
	s2 := "ab"
	n2 := 2
	fmt.Println(getMaxRepetitions(s1, n1, s2, n2))
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	pass2s, idx2s := make([]int, len(s2)+1), make([]int, len(s2)+1)

	pass2, idx2 := 0, 0

	for pass1 := 1; pass1 <= n1; pass1++ {
		for idx1 := 0; idx1 < len(s1); idx1++ {
			if s1[idx1] == s2[idx2] {
				idx2++
				if idx2 == len(s2) {
					idx2 = 0
					pass2++
				}
			}
		}
		pass2s[pass1] = pass2
		idx2s[pass1] = idx2

		for prePass1 := 0; prePass1 < pass1; prePass1++ {
			if idx2s[prePass1] == idx2 {
				repeatCount := (n1 - prePass1) / (pass1 - prePass1)
				remainPass1Count := (n1 - prePass1) % (pass1 - prePass1)
				prefixPass2Num := pass2s[prePass1]
				repetitivePass2Num := repeatCount * (pass2s[pass1] - pass2s[prePass1])
				suffixPass2Num := pass2s[prePass1+remainPass1Count] - pass2s[prePass1]
				overallPass2Num := prefixPass2Num + repetitivePass2Num + suffixPass2Num
				return overallPass2Num / n2
			}
		}
	}

	// no repeative part found
	return pass2s[n1] / n2
}
