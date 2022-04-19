package p2243

import (
	"bytes"
	"fmt"
)

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	var buf bytes.Buffer
	var sum int

	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
		if (i+1)%k == 0 || i+1 == len(s) {
			buf.WriteString(fmt.Sprintf("%d", sum))
			sum = 0
		}
	}
	return digitSum(buf.String(), k)
}
