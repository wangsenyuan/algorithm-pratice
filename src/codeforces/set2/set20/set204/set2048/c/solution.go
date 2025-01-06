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
		s := readString(reader)
		res := solve(s)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(s string) []int {
	n := len(s)

	first_zero := -1

	for i := range n {
		if s[i] == '0' {
			first_zero = i
			break
		}
	}

	if first_zero < 0 {
		// all one
		return []int{1, n, n, n}
	}

	// 第二个字符的长度是 n - first_zero, m
	m := n - first_zero
	tmp := make([]byte, n)
	copy(tmp, s[:first_zero])

	res := []int{1, n, first_zero + 1, n}
	ans := make([]byte, n)
	copy(ans, s[:first_zero])
	for i := first_zero; i < n; i++ {
		ans[i] = '0'
	}

	for i := 0; i+m <= n; i++ {
		if s[i] == '0' {
			continue
		}
		for j := 0; j < m; j++ {
			if s[i+j] == s[first_zero+j] {
				tmp[first_zero+j] = '0'
			} else {
				tmp[first_zero+j] = '1'
			}
		}
		if len(ans) == 0 {
			copy(ans, tmp)
			res[2] = i + 1
			res[3] = i + m
		} else {
			j := first_zero
			for j < n && ans[j] == tmp[j] {
				j++
			}
			if j < n && ans[j] < tmp[j] {
				copy(ans, tmp)
				res[2] = i + 1
				res[3] = i + m
			}
		}
	}

	return res
}
