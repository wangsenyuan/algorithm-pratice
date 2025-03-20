package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	k := readNum(reader)
	s := readString(reader)
	res := solve(k, s)
	fmt.Println(res)
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

func solve(k int, s string) string {
	n := len(s)
	buf := []byte(s)
	var marks []int
	var flag int
	for l, r := 0, n-1; l <= r; l, r = l+1, r-1 {
		if s[l] != '?' && s[r] != '?' && s[l] != s[r] {
			return "IMPOSSIBLE"
		}

		if s[l] == '?' && s[r] == '?' {
			marks = append(marks, l)
		} else if s[l] == '?' {
			buf[l] = s[r]
		} else if s[r] == '?' {
			buf[r] = s[l]
		}
		if buf[l] != '?' {
			flag |= 1 << (buf[l] - 'a')
		}
	}
	j := len(marks)
	// 从尾部开始
	for i := k - 1; i >= 0 && j > 0; i-- {
		l := marks[j-1]
		r := n - 1 - l
		if flag&(1<<i) == 0 {
			buf[l] = 'a' + byte(i)
			buf[r] = 'a' + byte(i)
			flag |= 1 << i
			j--
		}
	}

	if flag != (1<<k)-1 {
		return "IMPOSSIBLE"
	}

	for j > 0 {
		l := marks[j-1]
		r := n - 1 - l
		buf[l] = 'a'
		buf[r] = 'a'
		j--
	}

	return string(buf)
}
