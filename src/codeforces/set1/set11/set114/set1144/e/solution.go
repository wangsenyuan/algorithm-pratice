package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)
	fmt.Println(res)
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

func solve(s string, t string) string {
	n := len(s)

	a := toNumArray(s)
	b := toNumArray(t)
	for i := n - 1; i >= 0; i-- {
		a[i] += b[i]
		if i > 0 {
			a[i-1] += a[i] / 26
			a[i] %= 26
		}
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		rem := a[i] % 2
		a[i] /= 2
		buf[i] = byte(a[i] + 'a')
		if rem > 0 {
			a[i+1] += rem * 26
		}
	}

	return string(buf)
}

func toNumArray(s string) []int {
	n := len(s)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(s[i] - 'a')
	}
	return arr
}
