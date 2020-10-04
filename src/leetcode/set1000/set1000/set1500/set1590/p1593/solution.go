package p1593

import "bytes"

func reorderSpaces(text string) string {
	var cnt int
	ss := make([]string, 0, 1+len(text)/10)
	var buf bytes.Buffer
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			cnt++

			if buf.Len() > 0 {
				ss = append(ss, buf.String())
			}

			buf.Reset()
			continue
		}
		buf.WriteByte(text[i])
	}

	if buf.Len() > 0 {
		ss = append(ss, buf.String())
	}

	var res bytes.Buffer

	n := len(ss)
	if n == 1 {
		res.WriteString(ss[0])
		for i := 0; i < cnt; i++ {
			res.WriteByte(' ')
		}
		return res.String()
	}

	x := cnt / (n - 1)
	r := cnt % (n - 1)

	for i := 0; i < n-1; i++ {
		res.WriteString(ss[i])
		for j := 0; j < x; j++ {
			res.WriteByte(' ')
		}
	}
	res.WriteString(ss[n-1])
	for j := 0; j < r; j++ {
		res.WriteByte(' ')
	}
	return res.String()
}
