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
		readNum(reader)
		a := readString(reader)
		b := readString(reader)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
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

func solve(a string, b string) []int {
	var res []int
	var flip int
	n := len(a)
	var idx int
	for i := n - 1; i >= 0; i-- {
		if flip^check(a[idx] == b[i]) == 1 {
			res = append(res, 1)
		}
		res = append(res, i+1)
		if flip == 1 {
			idx -= i
		} else {
			idx += i
		}
		flip ^= 1
	}

	return res
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
func solve1(a string, b string) []int {
	// 先建a全部变成0
	// n := len(a)

	x := process(a)
	y := process(b)
	reverse(y)
	x = append(x, y...)

	for i := 0; i < len(x); i++ {
		x[i]++
	}

	return x
}

func process(s string) []int {
	var res []int
	n := len(s)

	for i := 0; i < n; {
		if s[i] == '0' {
			i++
			continue
		}
		j := i
		for i < n && s[i] == '1' {
			i++
		}
		if j > 0 {
			res = append(res, j-1)
		}
		res = append(res, i-1)
	}
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
