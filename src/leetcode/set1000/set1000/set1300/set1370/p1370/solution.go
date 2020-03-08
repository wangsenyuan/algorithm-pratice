package p1370

import (
	"bytes"
)

func sortString(s string) string {
	cnt := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}

	var ps []*Pair

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			ps = append(ps, NewPair(byte('a'+i), cnt[i]))
		}
	}

	var buf bytes.Buffer
	for len(ps) > 0 {
		var j int
		for i := 0; i < len(ps); i++ {
			buf.WriteByte(ps[i].letter)
			ps[i].cnt--
			if ps[i].cnt > 0 {
				ps[j] = ps[i]
				j++
			}
		}

		for i := j - 1; i >= 0; i-- {
			buf.WriteByte(ps[i].letter)
			ps[i].cnt--
		}

		var k int

		for i := 0; i < j; i++ {
			if ps[i].cnt > 0 {
				ps[k] = ps[i]
				k++
			}
		}
		ps = ps[:k]
	}

	return buf.String()
}

type Pair struct {
	letter byte
	cnt    int
}

func NewPair(letter byte, cnt int) *Pair {
	return &Pair{letter, cnt}
}
