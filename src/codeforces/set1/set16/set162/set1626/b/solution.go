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
		x := readString(reader)
		res := solve(x)
		buf.WriteString(res)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(x string) string {
	// to get a larger result
	n := len(x)
	// n <=200000
	// 所以没法所有的生成再比较
	//  xabcy
	//  abc => a + b, b + c,
	// a(b+c) (a+b)c
	if n == 2 {
		tmp := int(x[0]-'0') + int(x[1]-'0')
		return fmt.Sprintf("%d", tmp)
	}

	last := -1

	for i := 0; i+1 < n; i++ {
		tmp := int(x[i]-'0') + int(x[i+1]-'0')
		if tmp >= 10 {
			last = i
		}
	}

	if last >= 0 {
		return merge(x, last)
	}

	for i := 0; i+1 < n; i++ {
		tmp := int(x[i]-'0') + int(x[i+1]-'0')
		// tmp < 10
		if tmp > int(x[i]-'0') {
			return merge(x, i)
		}
	}
	return merge(x, 0)
}

func merge(x string, i int) string {
	tmp := int(x[i]-'0') + int(x[i+1]-'0')
	return fmt.Sprintf("%s%d%s", x[:i], tmp, x[i+2:])
}

func bruteForce(x string) string {
	var res string
	n := len(x)
	for i := 0; i+1 < n; i++ {
		tmp := int(x[i]-'0') + int(x[i+1]-'0')
		s := fmt.Sprintf("%s%d%s", x[:i], tmp, x[i+2:])
		if len(res) < 0 || (len(res) < len(s) || len(res) == len(s) && res < s) {
			res = s
		}
	}

	return res
}
