package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(decodeString("2[20[bc]31[xy]]xd4[rt]"))
	//fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	var buf bytes.Buffer
	m := 0
	for m < len(s) {
		if s[m] > '0' && s[m] <= '9' {
			break
		}
		buf.WriteByte(s[m])
		m++
	}
	if m == len(s) {
		return buf.String()
	}

	strs := split(s[m:])

	for _, str := range strs {
		k, x := decompose(str)
		if k == 0 {
			buf.WriteString(x)
			continue
		}

		y := decodeString(x)
		for i := 0; i < k; i++ {
			buf.WriteString(y)
		}
	}

	return buf.String()
}

func decompose(s string) (int, string) {
	k := 0
	i := 0
	for ; i < len(s); i++ {
		if s[i] == '[' {
			break
		}
		x := s[i] - '0'
		k = k*10 + int(x)
	}

	if i == len(s) {
		return 0, s
	}

	return k, s[i+1 : len(s)-1]
}

func split(s string) []string {
	var res []string
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			j := pair(s, i)
			res = append(res, s[k:j+1])
			i = j
			k = j + 1
		}

		if s[i] > '0' && i > k && s[k] >= 'a' && s[k] <= 'z' {
			res = append(res, s[k:i])
			k = i
		}
	}

	if k < len(s) {
		res = append(res, s[k:])
	}

	return res
}

func pair(s string, from int) int {
	level := 0
	for i := from; i < len(s); i++ {
		if s[i] == '[' {
			level++
		}
		if s[i] == ']' {
			level--
		}
		if level == 0 {
			return i
		}
	}
	panic("wrong format")
}
