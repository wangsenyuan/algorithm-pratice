package p1271

import (
	"bytes"
)

const S = "0123456789ABCDEF"

func toHexspeak(num string) string {
	n := len(num)

	var N int64

	for i := 0; i < n; i++ {
		N = N*10 + int64(num[i]-'0')
	}

	var res bytes.Buffer

	for N > 0 {
		r := N % 16
		res.WriteByte(S[r])
		N /= 16
	}
	bs := res.Bytes()
	for i := 0; i < len(bs); i++ {
		if bs[i] == '1' {
			bs[i] = 'I'
		} else if bs[i] == '0' {
			bs[i] = 'O'
		} else if bs[i] < 'A' || bs[i] > 'F' {
			return "ERROR"
		}
	}

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
