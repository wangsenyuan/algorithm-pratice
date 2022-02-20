package p2182

import "bytes"

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, x := range s {
		cnt[int(x-'a')]++
	}
	repeatLimit++
	var buf bytes.Buffer
	for i := 25; i >= 0; i-- {
		k := i - 1
		var p int
		for cnt[i] > 0 {
			p++
			if p%repeatLimit == 0 {
				for k >= 0 && cnt[k] == 0 {
					k--
				}
				if k < 0 {
					break
				}
				cnt[k]--
				buf.WriteByte(byte(k + 'a'))
			} else {
				buf.WriteByte(byte(i + 'a'))
				cnt[i]--
			}
		}
	}

	return buf.String()
}
