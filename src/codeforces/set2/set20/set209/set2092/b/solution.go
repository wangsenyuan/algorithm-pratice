package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	buf.WriteTo(os.Stdout)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) bool {
	readString(reader)
	a := readString(reader)
	b := readString(reader)
	return solve(a, b)
}

func solve(a string, b string) bool {
	n := len(a)

	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		if a[i] == '1' {
			cnt[i&1]--
		}
		if b[i] == '0' {
			cnt[(i&1)^1]++
		}
	}
	return cnt[0] >= 0 && cnt[1] >= 0
}
