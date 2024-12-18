package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int) string {

	m := 2

	// 13333377 是最多的一种

	for {
		sum := m * (m - 1) / 2
		if 2*sum >= n {
			break
		}
		m++
	}
	var arr []int
	for n > 0 {
		sum := m * (m - 1) / 2
		for n >= sum {
			arr = append(arr, m)
			n -= sum
		}
		m--
	}

	var buf []byte

	buf = append(buf, '1')

	arr = append(arr, 0)

	reverse(arr)

	for i := 0; i+1 < len(arr); i++ {
		cnt := arr[i+1] - arr[i]
		for cnt > 0 {
			buf = append(buf, '3')
			cnt--
		}
		buf = append(buf, '7')
	}

	return string(buf)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
