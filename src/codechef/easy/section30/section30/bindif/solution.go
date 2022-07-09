package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	a, b := solve(len(s), s)

	fmt.Println(a)
	fmt.Println(b)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, s string) (string, string) {
	var a bytes.Buffer
	var b bytes.Buffer

	s = reverseString(s)

	for i := 0; i < n; {
		if s[i] == '0' {
			a.WriteByte('0')
			b.WriteByte('0')
			i++
			continue
		}
		j := i
		for i < n && s[i] == '1' {
			i++
		}
		if i == j+1 {
			// only one
			a.WriteByte('1')
			b.WriteByte('0')
			continue
		}
		l := i - j
		a.WriteByte('0')
		b.WriteByte('1')
		for k := 0; k < l-1; k++ {
			a.WriteByte('0')
			b.WriteByte('0')
		}
		a.WriteByte('1')
		b.WriteByte('0')
	}

	sa := a.Bytes()
	reverse(sa)
	sb := b.Bytes()
	reverse(sb)
	var p int
	for p < len(sb) && sb[p] == '0' {
		p++
	}
	if p == len(sb) {
		p--
	}
	return string(sa), string(sb[p:])
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseString(x string) string {
	buf := []byte(x)
	reverse(buf)
	return string(buf)
}
