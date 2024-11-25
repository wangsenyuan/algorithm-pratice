package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(n, s, t)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

var all = []string{"abc", "acb", "bac", "bca", "cab", "cba"}

func solve(n int, s string, t string) string {
	if n == 1 {
		for _, x := range all {
			if !strings.Contains(x, s) && !strings.Contains(x, t) {
				return x
			}
		}
		return ""
	}

	avoid := make(map[string]int)

	if s[0] != s[1] {
		avoid[s]++
	}
	if t[0] != t[1] {
		avoid[t]++
	}

	if len(avoid) == 0 {
		return strings.Repeat("abc", n)
	}

	if len(avoid) == 1 {
		k := keys(avoid)[0]
		for _, x := range all {
			y := x + x
			if !strings.Contains(y, k) {
				return strings.Repeat(x, n)
			}
		}
		return ""
	}
	buf := make([]byte, 3*n)

	if reverse(s) == t {
		// ab ba
		// 那必须把c放在中间
		for i := 0; i < n; i++ {
			buf[i] = s[0]
		}
		var w byte = 'a'
		if s[0] == w || s[1] == w {
			w++
		}
		if s[0] == w || s[1] == w {
			w++
		}

		for i := n; i < 2*n; i++ {
			buf[i] = w
		}
		for i := 2 * n; i < 3*n; i++ {
			buf[i] = s[1]
		}

		return string(buf)
	}

	// len(avoid) = 2, ab, bc（他们肯定有一个是共同的字符）
	if s[1] == t[0] {
		// ab bc
		for i := 0; i < n; i++ {
			buf[i] = t[1]
		}
		for i := n; i < 2*n; i++ {
			buf[i] = t[0]
		}
		for i := 2 * n; i < 3*n; i++ {
			buf[i] = s[0]
		}
	} else if t[1] == s[0] {
		// ab ca
		for i := 0; i < n; i++ {
			buf[i] = s[1]
		}
		for i := n; i < 2*n; i++ {
			buf[i] = s[0]
		}
		for i := 2 * n; i < 3*n; i++ {
			buf[i] = t[0]
		}
	} else if s[0] == t[0] {
		// ab  ac
		for i := 0; i < n; i++ {
			buf[i] = s[1]
		}
		for i := n; i < 2*n; i++ {
			buf[i] = t[1]
		}
		for i := 2 * n; i < 3*n; i++ {
			buf[i] = t[0]
		}
	} else {
		// s[1] = t[1], ac, bc
		for i := 0; i < n; i++ {
			buf[i] = s[1]
		}
		for i := n; i < 2*n; i++ {
			buf[i] = s[0]
		}
		for i := 2 * n; i < 3*n; i++ {
			buf[i] = t[0]
		}
	}

	return string(buf)
}

func keys(m map[string]int) []string {
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
