package p459

// import "fmt"

// func main() {
// 	fmt.Println(repeatedSubstringPattern("abab"))
// 	fmt.Println(repeatedSubstringPattern("abb"))
// 	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))

// }

func repeatedSubstringPattern(str string) bool {
	lps := computeLPS(str)
	// fmt.Println(lps)
	n := len(lps)
	last := lps[n-1]
	if last > 0 && n%(n-last) == 0 {
		return true
	}
	return false
}

func computeLPS(str string) []int {
	lps := make([]int, len(str))

	for i := 1; i < len(str); i++ {
		j := lps[i-1]
		for j > 0 && str[i] != str[j] {
			j = lps[j-1]
		}
		lps[i] = j
		if str[i] == str[j] {
			lps[i]++
		}
	}
	return lps
}

func repeatedSubstringPattern1(str string) bool {
	for i := 1; i <= len(str)/2; i++ {
		sub := str[:i]
		j := i
		for ; j+i <= len(str); j += i {
			if str[j:j+i] != sub {
				break
			}
		}
		if j == len(str) {
			return true
		}
	}

	return false
}
