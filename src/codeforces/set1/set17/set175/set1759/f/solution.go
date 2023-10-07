package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, p := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(p, a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(p int, x []int) int {
	// 如果一个数字未出现在原来的序列中
	// 考虑最坏的一种情况，先将x递增到最后一位是0
	// 然后+1，直到最大的未出现的字符
	// 如果是这样，是不是太简单了？只需要考虑最后一位就可以了
	digits := make(map[int]int)

	for _, d := range x {
		digits[d]++
	}

	if len(digits) == p {
		return 0
	}

	newDigit := 1

	n := len(x)

	for i := n - 2; i >= 0; i-- {
		if x[i] < p-1 {
			newDigit = x[i] + 1
			break
		}
	}

	check := func(l int, r int, useNewDigit bool) bool {
		for i := l; i <= r; i++ {
			if useNewDigit && i == newDigit {
				continue
			}
			if digits[i] == 0 {
				return false
			}
		}
		return true
	}

	l, r := 0, p-1
	last := x[n-1]

	for l < r {
		m := (l + r) >> 1
		var tmp bool

		if last+m >= p {
			tmp = check(last+m+1-p, last-1, true)
		} else {
			tmp = check(0, last-1, false) && check(last+m+1, p-1, false)
		}
		if tmp {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
