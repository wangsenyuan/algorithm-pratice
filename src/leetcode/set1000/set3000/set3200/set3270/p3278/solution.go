package p3278

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	ss := strings.Split(date, "-")
	h, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])
	d, _ := strconv.Atoi(ss[2])
	return fmt.Sprintf("%s-%s-%s", toBinary(h), toBinary(m), toBinary(d))
}

func toBinary(num int) string {
	var buf []byte
	for num > 0 {
		buf = append(buf, byte(num%2+'0'))
		num /= 2
	}
	reverse(buf)

	return string(buf)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
