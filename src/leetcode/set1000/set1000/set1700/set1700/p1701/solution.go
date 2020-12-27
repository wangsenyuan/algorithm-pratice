package p1701

import "bytes"

func maximumBinaryString(binary string) string {
	n := len(binary)
	var flag = true
	var left, right int
	for i := 0; i < n; i++ {
		flag = flag && binary[i] == '1'
		if flag && binary[i] == '1' {
			left++
		}
		if !flag && binary[i] == '1' {
			right++
		}
	}

	if left+right > n-1 {
		return binary
	}

	k := n - right - 1
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if i == k {
			buf.WriteByte('0')
		} else {
			buf.WriteByte('1')
		}
	}

	return buf.String()
}
