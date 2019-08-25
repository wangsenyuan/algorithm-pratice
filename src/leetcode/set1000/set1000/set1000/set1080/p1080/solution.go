package p1080

import (
	"bytes"
)

func smallestSubsequence(text string) string {
	n := len(text)

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		j := int(text[i] - 'a')
		cnt[j]++
	}

	buf := make([]int, 26)
	pos := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
	}
	var p int
	for i := 0; i < n; i++ {
		c := text[i]
		j := int(c - 'a')
		cnt[j]--

		if pos[j] < 0 {
			for p > 0 && cnt[buf[p-1]] > 0 && buf[p-1] > j {
				pos[buf[p-1]] = -1
				p--
			}
			buf[p] = j
			pos[j] = p
			p++
		}
	}

	var res bytes.Buffer

	for i := 0; i < p; i++ {
		a := byte('a' + buf[i])
		res.WriteByte(a)
	}

	return res.String()
}
