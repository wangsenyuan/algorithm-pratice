package p3039

import (
	"bytes"
	"sort"
)

func lastNonEmptyString(s string) string {
	pos := make([][]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = make([]int, 0, 1)
	}

	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		pos[x] = append(pos[x], i)
	}
	var ln int
	for i := 0; i < 26; i++ {
		ln = max(ln, len(pos[i]))
	}
	var arr []int
	for i := 0; i < 26; i++ {
		if ln == len(pos[i]) {
			arr = append(arr, pos[i][ln-1])
		}
	}

	sort.Ints(arr)
	var buf bytes.Buffer

	for _, i := range arr {
		buf.WriteByte(s[i])
	}

	return buf.String()
}
