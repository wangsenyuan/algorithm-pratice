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
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(s string) string {
	n := len(s)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		x := s[i]
		if x == '?' {
			arr[i] = 1
		} else {
			arr[i] = int(x - '0')
		}
	}

	for i := 0; i < n; i++ {
		if s[i] == '?' {
			var start int
			if i == 0 {
				start++
			}
			for num := start; num <= 9; num++ {
				arr[i] = num
				if check(toNum(arr)) {
					return toString(arr)
				}
			}

			arr[i] = 1
		}
	}

	return ""
}

func check(num int64) bool {
	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for x := int64(3); x <= num/x; x += 2 {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func toString(arr []int) string {
	buf := make([]byte, len(arr))

	for i := 0; i < len(arr); i++ {
		buf[i] = byte(arr[i] + '0')
	}
	return string(buf)
}

func toNum(arr []int) int64 {
	var res int64
	for i := 0; i < len(arr); i++ {
		res = res*10 + int64(arr[i])
	}
	return res
}
