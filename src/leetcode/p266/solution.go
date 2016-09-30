package main

func canPermutePalindrome(s string) bool {
	cnts := make(map[rune]int)
	for _, b := range s {
		cnts[b]++
	}
	odd := false
	for _, v := range cnts {
		if v%2 == 1 {
			if !odd {
				odd = true
			} else {
				return false
			}
		}
	}

	if len(s)%2 == 1 {
		return odd == true
	}
	return odd == false
}
