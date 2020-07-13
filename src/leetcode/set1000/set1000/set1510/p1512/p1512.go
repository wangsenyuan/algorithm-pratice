package p1512

import (
	"bytes"
	"fmt"
	"strings"
)

var mem map[string]string

func init() {
	mem = make(map[string]string)
	arr := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	for i := 0; i < len(arr); i++ {
		mem[arr[i]] = fmt.Sprintf("%02d", i+1)
	}
}

func reformatDate(date string) string {
	ss := strings.Split(date, " ")

	var buf bytes.Buffer

	buf.WriteString(ss[2])
	buf.WriteByte('-')

	buf.WriteString(mem[ss[1]])

	buf.WriteByte('-')

	var i int
	var day int
	for i < len(ss[0]) && ss[0][i] >= '0' && ss[0][i] <= '9' {
		day = day*10 + int(ss[0][i]-'0')
		i++
	}

	buf.WriteString(fmt.Sprintf("%02d", day))

	return buf.String()
}
