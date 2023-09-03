package p2840

import "sort"

func checkStrings(s1 string, s2 string) bool {
	a := getSortedAt(s1, 0)
	b := getSortedAt(s1, 1)
	c := getSortedAt(s2, 0)
	d := getSortedAt(s2, 1)

	return a == c && b == d
}

func getSortedAt(s string, pos int) string {
	var buf []byte
	for i := 0; i < len(s); i++ {
		if i&1 == pos {
			buf = append(buf, s[i])
		}
	}
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	return string(buf)
}
