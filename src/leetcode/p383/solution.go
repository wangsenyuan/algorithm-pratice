package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := len(ransomNote)
	n := len(magazine)

	if m > n {
		return false
	}

	a := make([]int, 26)

	for i := 0; i < m; i++ {
		a[ransomNote[i]-'a']++
	}

	for j := 0; j < n; j++ {
		a[magazine[j]-'a']--
	}

	for i := 0; i < 26; i++ {
		if a[i] > 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("bjaajgea", "affhiiicabhbdchbidghccijjbfjfhjeddgggbajhidhjchiedhdibgeaecffbbbefiabjdhggihccec"))
}
