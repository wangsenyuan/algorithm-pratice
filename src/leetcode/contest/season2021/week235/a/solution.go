package a

import "bytes"

func truncateSentence(s string, k int) string {
	var buf bytes.Buffer

	var kk int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			kk++
			if kk == k {
				break
			}
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
