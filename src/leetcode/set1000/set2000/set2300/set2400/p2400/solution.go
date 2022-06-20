package p2400

import "fmt"

func greatestLetter(s string) string {
	cnt := make([]int, 52)

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			cnt[int(s[i]-'a')+26]++
		} else {
			cnt[int(s[i]-'A')]++
		}
	}

	for i := 25; i >= 0; i-- {
		if cnt[i] > 0 && cnt[i+26] > 0 {
			return fmt.Sprintf("%c", byte(i+'A'))
		}
	}
	return ""
}
