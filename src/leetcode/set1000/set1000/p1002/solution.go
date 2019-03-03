package p1002

import (
	"fmt"
)

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	cnt := countChars(A[0])

	for i := 1; i < len(A); i++ {
		tmp := countChars(A[i])
		for j := 0; j < 26; j++ {
			cnt[j] = min(cnt[j], tmp[j])
		}
	}

	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			res = append(res, fmt.Sprintf("%c", byte(i+'a')))
		}
	}
	return res
}

func countChars(s string) []int {
	res := make([]int, 26)
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		res[x]++
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
