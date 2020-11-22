package p5606

import "bytes"

func getSmallestString(n int, k int) string {
	//take a as more as possible, then b, then c and so on
	// when we can get k-1 from n-1 string, then we could get a
	var buf bytes.Buffer

	for i := 1; i <= n; i++ {
		left := 26 * (n - i)
		for x := 1; x <= 26; x++ {
			if k-x <= left {
				buf.WriteByte(byte(x - 1 + 'a'))
				k -= x
				break
			}
		}
	}

	return buf.String()
}
