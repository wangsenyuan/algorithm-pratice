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
		x := readString(reader)
		for i := 0; i < len(x); i++ {
			if x[i] == ' ' {
				x = x[i+1:]
				break
			}
		}
		s := readString(reader)
		res := solve(s, x)
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

func solve(s string, x string) int {
	cur := x[0]
	if cur == 'g' {
		// no waiting
		return 0
	}
	n := len(s)
	green_pos := -1
	ans := -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'g' {
			green_pos = i
		} else if green_pos > 0 && cur == s[i] {
			ans = max(ans, green_pos-i)
		}
	}

	// green_pos is now the first position
	for i := n - 1; i >= 0 && s[i] != 'g'; i-- {
		if s[i] == cur {
			ans = max(ans, n+green_pos-i)
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
