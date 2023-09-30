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
		n := readNum(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		ok, res := solve(a, b)
		if !ok {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
			}
		}
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

func solve(a string, b string) (bool, [][]int) {
	n := len(a)
	if n == 1 {
		if a == b {
			return a == "0", nil
		}
		// b won't change
		return b == "0", [][]int{{1, 1}}
	}
	var res [][]int

	add := func(l int, r int) {
		res = append(res, []int{l + 1, r + 1})
	}

	var flip int
	prev := -1
	for i := 0; i < n; i++ {
		x := int(a[i]-'0') ^ flip
		y := int(b[i] - '0')
		if x == 1 {
			// a flip
			if prev == y {
				// 如果和前面的相同，前面的flip后，就没法带着i继续了
				return false, nil
			}
			// prev != y
			add(i, n-1)
			flip ^= 1
			prev = y
		}
		if prev != -1 && y != prev {
			return false, nil
		}
		prev = y
	}

	if prev == 0 {
		return true, res
	}
	// prev = 1
	// change to 00010
	add(n-2, n-2)
	// change to 000001
	add(n-2, n-1)
	add(n-1, n-1)
	return true, res
}
