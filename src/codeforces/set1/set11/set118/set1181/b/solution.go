package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readString(reader)
	n := readString(reader)
	res := solve(n)
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

func solve(num string) string {
	n := len(num)

	if n == 2 {
		return add(num[:1], num[1:])
	}
	h := n / 2
	var ans string
	if num[h] != '0' {
		ans = add(num[:h], num[h:])
	}
	for i := h + 1; i < n; i++ {
		if num[i] != '0' {
			// must be positive
			tmp := add(num[:i], num[i:])
			if len(ans) == 0 {
				ans = tmp
			} else {
				ans = compare(ans, tmp)
			}
			break
		}
	}

	for i := h - 1; i > 0; i-- {
		if num[i] != '0' {
			tmp := add(num[:i], num[i:])
			if len(ans) == 0 {
				ans = tmp
			} else {
				ans = compare(ans, tmp)
			}
			break
		}
	}

	return ans
}

func compare(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	if len(a) == len(b) && a < b {
		return a
	}
	return b
}

func add(a string, b string) string {
	sa := reverse(a)
	sb := reverse(b)

	var carray int
	var res []byte

	for i := 0; i < len(sa) || i < len(sb); i++ {
		var x, y int
		if i < len(sa) {
			x = int(sa[i] - '0')
		}
		if i < len(sb) {
			y = int(sb[i] - '0')
		}
		z := x + y + carray
		carray = z / 10
		z %= 10
		res = append(res, byte(z)+'0')
	}
	if carray > 0 {
		res = append(res, byte(carray)+'0')
	}
	return reverse(string(res))
}

func reverse(s string) string {
	n := len(s)
	bs := []byte(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
