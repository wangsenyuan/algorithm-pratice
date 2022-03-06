package p2194

import "bytes"

func cellsInRange(s string) []string {
	c1 := s[0]
	r1 := s[1]
	c2 := s[3]
	r2 := s[4]
	var res []string
	for c := c1; c <= c2; c++ {
		for r := r1; r <= r2; r++ {
			var buf bytes.Buffer
			buf.WriteByte(c)
			buf.WriteByte(r)
			res = append(res, buf.String())
		}
	}
	return res
}
