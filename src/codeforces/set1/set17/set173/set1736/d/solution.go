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
		s := readString(reader)
		ok, b, p := solve(s)
		if !ok {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d", len(b)))
		for i := 0; i < len(b); i++ {
			buf.WriteString(fmt.Sprintf(" %d", b[i]))
		}
		buf.WriteByte('\n')
		for i := 0; i < len(p); i++ {
			buf.WriteString(fmt.Sprintf("%d ", p[i]))
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

func solve(s string) (bool, []int, []int) {
	n := len(s)
	h := n / 2
	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'0')]++
	}

	if cnt[0]%2 == 1 || cnt[1]%2 == 1 {
		return false, nil, nil
	}
	// p中cnt[0]/2, cnt[1]/2 个1， 要个数相同比较容易，但是必须位置也匹配
	var arr []int
	var res []int
	for i := 0; i < h; i++ {
		if s[i*2] != s[i*2+1] {
			x := int(s[i*2] - '0')
			if x == len(arr)&1 {
				arr = append(arr, i*2+1)
			} else {
				arr = append(arr, i*2+2)
			}
		}
		res = append(res, i*2+1)
	}

	return true, arr, res
}
