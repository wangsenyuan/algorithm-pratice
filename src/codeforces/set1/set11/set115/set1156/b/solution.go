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
		s := readString(reader)
		res := solve(s)
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

func solve(s string) string {
	cnt := make([]int, 26)
	for _, x := range []byte(s) {
		cnt[int(x-'a')]++
	}

	var arr []int

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			arr = append(arr, i)
		}
	}

	if len(arr) == 1 {
		// already good
		return s
	}
	if len(arr) == 2 {
		if arr[0]+1 == arr[1] {
			return "No answer"
		}
		// also good
		return s
	}
	var buf bytes.Buffer
	add := func(x int, c int) {
		for c > 0 {
			buf.WriteByte(byte(x + 'a'))
			c--
		}
	}

	if len(arr) == 3 {

		if arr[0]+1 < arr[1] {
			add(arr[1], cnt[arr[1]])
			add(arr[0], cnt[arr[0]])
			add(arr[2], cnt[arr[2]])
			return buf.String()
		}

		if arr[1]+1 < arr[2] {
			add(arr[0], cnt[arr[0]])
			add(arr[2], cnt[arr[2]])
			add(arr[1], cnt[arr[1]])
			return buf.String()
		}

		return "No answer"
	}

	// len(arr) >= 4, has solution always
	// cadb
	// 其他的放在两边
	res := make([]int, len(arr)*2)
	head, tail := len(arr), len(arr)
	res[head] = arr[2]
	head++
	res[head] = arr[0]
	head++
	res[head] = arr[3]
	head++
	res[head] = arr[1]
	head++
	for i := 4; i < len(arr); {
		res[head] = arr[i]
		i++
		head++
		if i < len(arr) {
			tail--
			res[tail] = arr[i]
			i++
		}
	}

	for i := tail; i < head; i++ {
		add(res[i], cnt[res[i]])
	}

	return buf.String()
}
