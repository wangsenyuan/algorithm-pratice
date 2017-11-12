package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("Snug & raw was I ere I saw war & guns."))
}

func reverseVowels(s string) string {
	ans := make([]byte, len(s))

	i, j := 0, len(s)-1
	for i <= j {
		if !isVowel(s[i]) {
			ans[i] = s[i]
			i++
			continue
		}

		if !isVowel(s[j]) {
			ans[j] = s[j]
			j--
			continue
		}

		ans[i], ans[j] = s[j], s[i]
		i++
		j--
	}
	return string(ans)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'o' || b == 'e' || b == 'i' || b == 'u' ||
		b == 'A' || b == 'O' || b == 'E' || b == 'I' || b == 'U'
}
