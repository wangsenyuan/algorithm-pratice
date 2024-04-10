package p1701

import "bytes"

func maximumBinaryString(binary string) string {
	n := len(binary)
	buf := []byte(binary)
	// 10 => 01 意味着可以把0移动到左边
	// 00 => 10意味着可以把头部的0变成1
	var p int
	for p < n && buf[p] == '1' {
		p++
	}
	var cnt int
	for i := p; i < n; i++ {
		cnt += int(buf[i] - '0')
	}

	if cnt+p >= n-1 {
		return binary
	}

	// 先把cnt个移动到右边
	k := n - cnt - 1
	for i := p; i < n; i++ {
		if i == k {
			buf[i] = '0'
		} else {
			buf[i] = '1'
		}
	}
	return string(buf)
}

func maximumBinaryString1(binary string) string {
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
